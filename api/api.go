package api

import (
	"context"
	"fmt"
	"sync"

	"github.com/SphericalKat/katbox/internal/config"
	"github.com/gofiber/fiber/v2"
	log "github.com/sirupsen/logrus"
)

func StartListening(ctx context.Context, wg *sync.WaitGroup, engine fiber.Views) {
	app := fiber.New(fiber.Config{
		StreamRequestBody:     true,
		ServerHeader:          "Katbox",
		AppName:               "Katbox",
		Views:                 engine,
		DisableStartupMessage: true,
	})

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("index", fiber.Map{
			"Title": "Katbox",
		})
	})

	// mount routes
	MountUpload(app)

	go func(app *fiber.App) {
		log.Printf("Starting http server at: http://localhost:%s", config.Conf.Port)
		if err := app.Listen(fmt.Sprintf(":%s", config.Conf.Port)); err != nil {
			log.Fatalf("Unable to start http server: %s", err)
		}
	}(app)

	// listen for context cancellation
	<-ctx.Done()

	// shut down http server
	log.Info("Gracefully shutting down http server...")
	if err := app.Shutdown(); err != nil {
		log.Warn("Server shutdown Failed: ", err)
	}
	wg.Done()
}
