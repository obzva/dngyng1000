package server

import (
	"net/http"
	"strings"
)

func Run() error {
	http.HandleFunc("/posts/", postHandler)

	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))

	return http.ListenAndServe(":8080", nil)
}

func postHandler(w http.ResponseWriter, r *http.Request) {
	trimmedPath := strings.TrimPrefix(r.URL.Path, "/posts/")
	switch trimmedPath {
	case "":
		renderTemplate(w, "posts", postMap.s)
	default:
		post, err := postMap.Get(trimmedPath)
		if err != nil {
			http.NotFound(w, r)
			return
		}
		renderTemplate(w, "post", post)
	}
}
