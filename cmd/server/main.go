package main

import (
	"log"
	"log/slog"
	"net/http"
	"os"

	"github.com/obzva/dngyng1000/internal/post"
	"github.com/obzva/dngyng1000/internal/server"
	"github.com/obzva/dngyng1000/internal/template"
)

func main() {
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	tmplCache, err := template.NewCache()
	if err != nil {
		logger.Error("error occurred while initializing template cache")
		os.Exit(1)
	}

	postMap, err := post.NewMap()
	if err != nil {
		logger.Error("error occurred while initializing post map")
		os.Exit(1)
	}

	srv := server.New(logger, tmplCache, postMap)

	s := http.Server{
		Handler: srv,
		Addr:    ":3000",
	}

	if err := s.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
