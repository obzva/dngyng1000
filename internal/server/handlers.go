package server

import (
	"log/slog"
	"net/http"

	"github.com/obzva/dngyng1000/internal/post"
	"github.com/obzva/dngyng1000/internal/template"
)

func rootHandler(logger *slog.Logger, tc template.Cache) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		td := struct {
			PageTitle string
		}{
			PageTitle: "Home",
		}

		tc.Render(logger, w, r, "home.tmpl", td)
	}
}

func postsHandler(logger *slog.Logger, tc template.Cache, postMap post.Map) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		td := struct {
			PageTitle string
			Posts     []*post.Post
		}{
			PageTitle: "Posts",
			Posts:     postMap.Slices(),
		}

		tc.Render(logger, w, r, "posts.tmpl", td)
	}
}

func postHandler(logger *slog.Logger, tc template.Cache, postMap post.Map) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		slug := r.PathValue("slug")
		p := postMap[slug]

		td := struct {
			PageTitle string
			Post      *post.Post
		}{
			PageTitle: p.Title,
			Post:      p,
		}

		tc.Render(logger, w, r, "post.tmpl", td)
	}
}
