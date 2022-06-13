package main

import (
	"context"
	"io/fs"

	"net/http"
	"sync"

	"embed"

	"github.com/SphericalKat/katbox/api"
	"github.com/SphericalKat/katbox/internal/aws"
	"github.com/SphericalKat/katbox/internal/config"
	"github.com/SphericalKat/katbox/internal/db"
	"github.com/SphericalKat/katbox/internal/lifecycle"
	"github.com/gofiber/template/html"
	log "github.com/sirupsen/logrus"
)

//go:embed template/*
var template embed.FS

//go:embed static/*
var static embed.FS

func main() {
	// load config
	config.Load()

	// connect to s3
	aws.Connect()

	// create template engine
	tmplFs, err := fs.Sub(template, "template")
	if err != nil {
		log.Fatalf("error loading template: %v\n", err)
	}
	engine := html.NewFileSystem(http.FS(tmplFs), ".html")

	// create static file server
	staticFs, err := fs.Sub(static, "static")
	if err != nil {
		log.Fatalf("error loading static assets: %v\n", err)
	}
	staticHttp := http.FS(staticFs)

	// create a waitgroup for all tasks
	wg := sync.WaitGroup{}

	// create context for background tasks
	ctx, cancelFunc := context.WithCancel(context.Background())

	// connect to database
	wg.Add(1)
	go db.Connect(ctx, &wg)

	// start http server
	wg.Add(1)
	go api.StartListening(ctx, &wg, engine, staticHttp)

	// add signal handler to gracefully shut down tasks
	wg.Add(1)
	go lifecycle.ShutdownListener(&wg, &cancelFunc)

	// wait for all tasks to finish
	wg.Wait()

	log.Info("Graceful shutdown complete.")
}
