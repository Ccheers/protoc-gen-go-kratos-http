package main

import (
	"bytes"
	_ "embed"
	"regexp"
	"strings"
	"text/template"
)

//go:embed httpTemplate.tpl
var httpTemplate string

type serviceDesc struct {
	ServiceType     string // Greeter
	ServiceName     string // helloworld.Greeter
	Metadata        string // api/helloworld/helloworld.proto
	Methods         []*methodDesc
	MethodSets      []*methodDesc
	MiddlewareNames []string
}

type methodDesc struct {
	// method
	Name         string
	OriginalName string // The parsed original name
	Num          int
	Request      string
	Reply        string
	Comment      string
	// http_rule
	Path            string
	Method          string
	HasVars         bool
	HasBody         bool
	Body            string
	ResponseBody    string
	MiddlewareNames []string
}

var middleWareMatch = regexp.MustCompile("@[A-Za-z0-9_]+")

func parseMiddleware(str string) []string {
	strs := middleWareMatch.FindAllString(str, -1)
	for i, str := range strs {
		strs[i] = strings.TrimPrefix(str, "@")
	}
	return strs
}

func (s *serviceDesc) execute() string {
	sets := make(map[string]struct{})
	for _, m := range s.Methods {
		_, ok := sets[m.Name]
		if !ok {
			s.MethodSets = append(s.MethodSets, m)
			sets[m.Name] = struct{}{}
		}
	}
	buf := new(bytes.Buffer)
	tmpl, err := template.New("http").Parse(strings.TrimSpace(httpTemplate))
	if err != nil {
		panic(err)
	}
	if err := tmpl.Execute(buf, s); err != nil {
		panic(err)
	}
	return strings.Trim(buf.String(), "\r\n")
}
