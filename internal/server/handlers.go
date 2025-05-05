package server

import (
	"log/slog"
	"net/http"

	"github.com/obzva/dngyng1000/internal/template"
)

func rootHandler(logger *slog.Logger, tmpl *template.Template) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		tmpl.Render(logger, w, r, "home.tmpl")
	}
}
