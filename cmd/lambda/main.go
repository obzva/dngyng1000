package main

import (
	"log/slog"
	"os"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/awslabs/aws-lambda-go-api-proxy/httpadapter"
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

	lambda.Start(httpadapter.New(srv).ProxyWithContext)
}
