package server

import (
	"log/slog"
	"net/http"

	"github.com/obzva/dngyng1000/internal/template"
)

func New(logger *slog.Logger, tmpl *template.Template) http.Handler {
	mux := routes(logger, tmpl)

	return mux
}
