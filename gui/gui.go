package gui

import (
	"html/template"
)

var testTemplate *template.Template

func init() {
	testTemplate = template.Must(template.ParseFiles("../html/test.html"))
}
