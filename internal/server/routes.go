package server

import (
	"log/slog"
	"net/http"

	"github.com/obzva/dngyng1000/internal/template"
	"github.com/obzva/dngyng1000/ui"
)

func routes(logger *slog.Logger, tmpl *template.Template) *http.ServeMux {
	mux := http.NewServeMux()

	fs := http.FileServerFS(ui.Files)
	mux.Handle("GET /static/", fs)

	mux.HandleFunc("GET /{$}", rootHandler(logger, tmpl))

	return mux
}
