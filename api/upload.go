package api

import (
	"io"
	"os"

	"github.com/SphericalKat/katbox/ent/user"
	"github.com/SphericalKat/katbox/internal/db"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

func uploadFile(c *fiber.Ctx) error {
	email := c.Cookies("authToken")
	if email == "" {
		return c.Redirect("/auth/login")
	}

	user, err := db.Client.User.Query().Where(user.EmailEQ(email)).First(c.Context())
	if err != nil {
		logrus.Error(err)
		return err
	}
	if user == nil {
		return fiber.ErrNotFound
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

	dbFile, err := db.Client.File.
		Create().
		SetStorageKey(fileHeader.Filename).
		SetUser(user).
		Save(c.Context())
	if err != nil {
		logrus.Error(err)
		return err
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"fileId": dbFile.ID,
	})
}

func MountUpload(app *fiber.App) {
	api := app.Group("/api")
	api.Post("/upload", uploadFile)
}
