package api

import "github.com/gofiber/fiber/v2"

var uploadFile fiber.Handler = func(c *fiber.Ctx) error {
	return c.SendString("uploaded")
}

func MountUpload(app *fiber.App) {
	api := app.Group("/api")
	api.Get("/upload", uploadFile)
}
