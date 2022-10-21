package api

import (
	"io"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

func uploadFile(c *fiber.Ctx) error {
	email := c.Cookies("authToken")
	if email == "" {
		return c.Redirect("/auth/login")
	}

	fileHeader, err := c.FormFile("file")
	if err != nil {
		logrus.Error(err)
		return err
	}

	file, err := fileHeader.Open()
	if err != nil {
		logrus.Error(err)
		return err
	}

	readBuffer, err := io.ReadAll(file)
	if err != nil {
		logrus.Error(err)
		return err
	}

	err = os.WriteFile(fileHeader.Filename, readBuffer, 0660)
	if err != nil {
		logrus.Error(err)
		return err
	}

	return nil
}

func MountUpload(app *fiber.App) {
	api := app.Group("/api")
	api.Post("/upload", uploadFile)
}
