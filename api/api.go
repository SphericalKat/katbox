package api

import (
	"context"
	"fmt"
	"net/http"
	"sync"

	"github.com/SphericalKat/katbox/ent/user"
	"github.com/SphericalKat/katbox/internal/config"
	"github.com/SphericalKat/katbox/internal/db"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/encryptcookie"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
	log "github.com/sirupsen/logrus"
)

func StartListening(ctx context.Context, wg *sync.WaitGroup, engine fiber.Views, static http.FileSystem) {
	app := fiber.New(fiber.Config{
		StreamRequestBody:     true,
		ServerHeader:          "Katbox",
		AppName:               "Katbox",
		Views:                 engine,
		DisableStartupMessage: true,
	})

	// encrypted cookies
	app.Use(encryptcookie.New(encryptcookie.Config{
		Key: config.Conf.SecretKey,
	}))

	// static file server
	app.Use("/static", filesystem.New(filesystem.Config{
		Root:   static,
		Browse: config.Conf.Env == "dev",
	}))

	app.Get("/", func(c *fiber.Ctx) error {
		email := c.Cookies("authToken")
		if email == "" {
			return c.Redirect("/auth/login")
		}

		exists, err := db.Client.User.Query().Where(user.EmailEQ(email)).Exist(c.Context())
		if err != nil {
			return err
		}

		if !exists {
			return c.Redirect("/auth/login")
		}

		return c.Render("index", fiber.Map{
			"Title": "Katbox",
		})
	})

	// mount routes
	MountUpload(app)
	MountAuth(app)

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
