package server

import (
	"log/slog"
	"net/http"

	"github.com/obzva/dngyng1000/internal/post"
	"github.com/obzva/dngyng1000/internal/template"
	"github.com/obzva/dngyng1000/internal/ui"
)

func routes(logger *slog.Logger, tc template.Cache, postMap post.Map) *http.ServeMux {
	mux := http.NewServeMux()

	fs := http.FileServerFS(ui.Files)
	mux.Handle("GET /static/", fs)

	mux.HandleFunc("GET /{$}", rootHandler(logger, tc))
	mux.HandleFunc("GET /posts", postsHandler(logger, tc, postMap))
	mux.HandleFunc("GET /posts/{slug}", postHandler(logger, tc, postMap))

	return mux
}
