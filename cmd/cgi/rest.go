package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"strings"

	"xie.sh.cn/panabit-ddns-go-manager/v2/pkg/html"
)

func callRestfulDispatcher(act string) {
	var res string
	switch strings.ToLower(act) {
	case strings.ToLower("UpdateCertificates"):
		res = renderRest(updateCertificates())
	case strings.ToLower("StartInstance"):
		res = renderRest(startInstance())
	case strings.ToLower("StopInstance"):
		res = renderRest(stopInstance())
	case strings.ToLower("DescribeInstanceStatus"):
		res = renderRest(describeInstanceStatus())
	default:
	}
	fmt.Println(res)
}

func renderRest(code int, data any) string {
	j, _ := json.Marshal(
		struct {
			Code int
			Data any
		}{
			Code: code,
			Data: data,
		},
	)
	d := struct {
		ContentType string
		Data        template.HTML
	}{
		ContentType: "application/json; charset=UTF-8",
		Data:        template.HTML(j),
	}
	s, _ := html.Render(DefaultTemplatePath, d, HttpTemplate, RestTemplate)
	return s
}
