package server

import (
	"fmt"
	"html/template"
	"net/http"
	"strings"
)

var (
	templates = template.Must(template.ParseFiles("templates/post.html"))
	postMap   = MustNewPostMap("posts")
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
		// postListHandler(w, r)
		fmt.Fprint(w, "not yet")
		return
	default:
		post, err := postMap.Get(trimmedPath)
		if err != nil {
			http.NotFound(w, r)
			return
		}
		renderTemplate(w, "post", post)
	}
}

func renderTemplate(w http.ResponseWriter, tmplName string, p *Post) {
	err := templates.ExecuteTemplate(w, tmplName+".html", p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
