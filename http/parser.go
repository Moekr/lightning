package http

import (
	"html/template"
	"strings"
)

func ParseTags(tags []string) template.HTML {
	if len(tags) == 0 {
		return "无可用标签"
	}
	ss := make([]string, len(tags))
	for idx, tag := range tags {
		ss[idx] = "<a href=\"/search.html?q=tag:" + tag + "\" target=\"_blank\">" + tag + "</a>"
	}
	return template.HTML(strings.Join(ss, " / "))
}
