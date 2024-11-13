package audit

import (
	"bytes"
	"io"
	"net/http"

	"github.com/tidwall/gjson"
)

// MetaValue 元数据值定义
type MetaValue struct {
	Extract string // 从请求中提取字段
	Const   string // 默认值
}

type Meta struct {
	Key   string    // 元数据键
	Value MetaValue // 元数据值
}

// ExtractedMeta 提取后的元数据
type ExtractedMeta struct {
	Key   string // 元数据键
	Value string // 提取后的值
}

type Audit struct {
	Module         string          // 模块
	Action         string          // 操作
	Metas          []Meta          // 元数据
	ExtractedMetas []ExtractedMeta // 提取后的元数据
}

// extractFromQuery 从查询参数中提取
func extractFromQuery(r *http.Request, metas []Meta) map[string]string {
	result := make(map[string]string)
	query := r.URL.Query()

	for _, meta := range metas {
		if meta.Value.Const != "" {
			result[meta.Key] = meta.Value.Const
		} else {
			value := query.Get(meta.Value.Extract)
			result[meta.Key] = value
		}
	}
	return result
}

// extractFromBody 从请求体中提取
func extractFromBody(r *http.Request, metas []Meta) (map[string]string, error) {
	result := make(map[string]string)

	// 读取请求体
	if r.Body == nil {
		return result, nil
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		return result, err
	}
	// 重新设置请求体，供后续中间件使用
	r.Body = io.NopCloser(bytes.NewBuffer(body))

	// 解析JSON
	parsed := gjson.Parse(string(body))
	for _, meta := range metas {
		if meta.Value.Const != "" {
			result[meta.Key] = meta.Value.Const
		} else {
			value := parsed.Get(meta.Value.Extract)
			if !value.Exists() {
				result[meta.Key] = ""
			}
			result[meta.Key] = value.String()

		}
	}
	return result, nil
}

func ExtractFromRequest(r *http.Request, audit *Audit) (*Audit, error) {
	if audit == nil {
		return audit, nil
	}
	// 如果没有需要提取的元数据，直接返回
	if len(audit.Metas) == 0 {
		return audit, nil
	}

	var err error
	var metaMap map[string]string

	switch r.Method {
	case http.MethodGet:
		metaMap = extractFromQuery(r, audit.Metas)
	default:
		metaMap, err = extractFromBody(r, audit.Metas)
		if err != nil {
			return audit, err
		}
	}

	// 将提取的结果转换为 ExtractedMeta 数组
	for k, v := range metaMap {
		audit.ExtractedMetas = append(audit.ExtractedMetas, ExtractedMeta{
			Key:   k,
			Value: v,
		})
	}
	return audit, nil
}

func NewAudit(module string, action string, metas []Meta) *Audit {
	return &Audit{
		Module: module,
		Action: action,
		Metas:  metas,
	}
}
