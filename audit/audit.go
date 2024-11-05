package audit

import (
	"io"
	"net/http"

	"github.com/tidwall/gjson"
)

type Audit struct {
	Module  string            // 模块
	Action  string            // 操作
	Extract map[string]string // 额外信息
}

func ExtractFromRequest(request *http.Request, mm map[string]string) map[string]string {
	switch request.Method {
	case http.MethodGet:
		query := request.URL.Query()
		for k, v := range mm {
			mm[k] = query.Get(v)
		}
	default:
		bs, _ := io.ReadAll(request.Body)
		parsed := gjson.Parse(string(bs))
		for k, v := range mm {
			data := parsed.Get(v).String()
			if data != "" {
				mm[k] = data
			}
		}
	}
	return mm
}

func NewAudit(module, action string, extract map[string]string) *Audit {
	return &Audit{
		Module:  module,
		Action:  action,
		Extract: extract,
	}
}
