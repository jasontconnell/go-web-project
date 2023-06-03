package handlers

import (
	"html/template"
)

func toHtml(str string) template.HTML {
	return template.HTML(str)
}
