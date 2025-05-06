package server

import (
	"log/slog"
	"net/http"

	"github.com/obzva/dngyng1000/internal/post"
	"github.com/obzva/dngyng1000/internal/template"
)

func New(logger *slog.Logger, tc template.Cache, postMap post.Map) http.Handler {
	mux := routes(logger, tc, postMap)

	return mux
}
