package gui

import (
	"html/template"
)

var Templates map[string]*template.Template

func init() {
	Templates = make(map[string]*template.Template)
	Templates["projects"] = template.Must(template.ParseFiles("../html/projects.html"))
}
