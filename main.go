package main

import (
	"context"
	"sync"

	"github.com/SphericalKat/katbox/internal/config"
	"github.com/SphericalKat/katbox/internal/db"
	"github.com/SphericalKat/katbox/internal/lifecycle"
	log "github.com/sirupsen/logrus"
)

func main() {
	// load config
	config.Load()

	// create a waitgroup for all tasks
	wg := sync.WaitGroup{}

	// create context for background tasks
	ctx, cancelFunc := context.WithCancel(context.Background())

	// connect to database
	wg.Add(1)
	go db.Connect(ctx, &wg)

	// add signal handler to gracefully shut down tasks
	wg.Add(1)
	go lifecycle.ShutdownListener(&wg, &cancelFunc)

	// wait for all tasks to finish
	wg.Wait()

	log.Info("Graceful shutdown complete.")

}
