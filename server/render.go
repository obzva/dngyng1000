package server

import (
	"html/template"
	"net/http"
)

const (
	tmplExt = ".tmpl"
	tmplDir = "templates"
)

var templates = template.Must(template.ParseGlob(tmplDir + "/*" + tmplExt))

func renderTemplate(w http.ResponseWriter, tmplName string, data any) {
	err := templates.ExecuteTemplate(w, tmplName+tmplExt, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
