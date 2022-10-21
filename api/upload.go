package api

import (
	"fmt"
	"time"

	"github.com/SphericalKat/katbox/ent"
	"github.com/SphericalKat/katbox/ent/file"
	"github.com/SphericalKat/katbox/ent/user"
	"github.com/SphericalKat/katbox/internal/config"
	"github.com/SphericalKat/katbox/internal/db"
	"github.com/SphericalKat/katbox/internal/storage"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
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

	contentType := fileHeader.Header.Get("Content-Type")
	if contentType == "" {
		contentType = "application/octet-stream"
	}

	file, err := fileHeader.Open()
	if err != nil {
		logrus.Error(err)
		return err
	}

	storageKey := fmt.Sprintf("%s/%s_%s", user.Email, uuid.NewString(), fileHeader.Filename)

	key, err := storage.UploadMinio(c.Context(), storageKey, contentType, file)
	if err != nil {
		logrus.Error(err)
		return err
	}
	logrus.Info(key)

	dbFile, err := db.Client.File.
		Create().
		SetFileName(fileHeader.Filename).
		SetStorageKey(storageKey).
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

func getFile(c *fiber.Ctx) error {
	email := c.Cookies("authToken")
	if email == "" {
		return c.Redirect("/auth/login")
	}

	fileIdStr := c.Params("fileId")
	fileId, err := uuid.Parse(fileIdStr)
	if err != nil {
		return fiber.ErrUnprocessableEntity
	}

	user, err := db.Client.User.
		Query().
		Where(user.EmailEQ(email)).
		WithFiles(func(fq *ent.FileQuery) {
			fq.Where(file.IDEQ(fileId))
		}).First(c.Context())
	if err != nil {
		logrus.Error(err)
		return err
	}
	if user == nil {
		return fiber.ErrNotFound
	}

	if len(user.Edges.Files) == 0 {
		return fiber.ErrNotFound
	}

	url, err := storage.MC.Presign(c.Context(), "GET", config.Conf.S3BucketName, user.Edges.Files[0].StorageKey, 7*24*time.Hour, nil)
	if err != nil {
		logrus.Error(err)
		return err
	}

	return c.Redirect(url.String())
}

func MountUpload(app *fiber.App) {
	app.Get("/file/:fileId", getFile)

	api := app.Group("/api")
	api.Post("/upload", uploadFile)
}
