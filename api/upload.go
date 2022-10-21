package api

import "github.com/gofiber/fiber/v2"

func uploadFile(c *fiber.Ctx) error {
	return nil
}

func MountUpload(app *fiber.App) {
	api := app.Group("/api")
	api.Post("/upload", uploadFile)
}
