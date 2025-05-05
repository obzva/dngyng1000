package main

import (
	"log"
	"log/slog"
	"net/http"
	"os"

	"github.com/obzva/dngyng1000/internal/server"
	"github.com/obzva/dngyng1000/internal/template"
)

func main() {
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	tmpl, err := template.New()
	if err != nil {
		logger.Error("error occurred while initializing templates")
		os.Exit(1)
	}

	srv := server.New(logger, tmpl)

	s := http.Server{
		Handler: srv,
		Addr:    ":3000",
	}

	if err := s.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
