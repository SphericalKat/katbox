package api

import (
	"fmt"
	"path/filepath"
	"strings"
	"time"

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

	ext := filepath.Ext(fileHeader.Filename)
	filename := strings.TrimSuffix(fileHeader.Filename, ext)

	file, err := fileHeader.Open()
	if err != nil {
		logrus.Error(err)
		return err
	}

	storageKey := fmt.Sprintf("%s/%s_%s%s", user.Email, filename, uuid.NewString(), ext)

	_, err = storage.UploadMinio(c.Context(), storageKey, contentType, file)
	if err != nil {
		logrus.Error(err)
		return err
	}

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
	//email := c.Cookies("authToken")
	//if email == "" {
	//	return c.Redirect("/auth/login")
	//}

	fileIdStr := c.Params("fileId")
	fileId, err := uuid.Parse(fileIdStr)
	if err != nil {
		return fiber.ErrUnprocessableEntity
	}

	dbFile, err := db.Client.File.
		Query().
		Where(file.IDEQ(fileId)).
		Only(c.Context())
	if err != nil {
		logrus.Error(err)
		return err
	}
	//if dbUser == nil {
	//	return fiber.ErrNotFound
	//}
	//
	//if len(dbUser.Edges.Files) == 0 {
	//	return fiber.ErrNotFound
	//}

	url, err := storage.MC.Presign(c.Context(), "GET", config.Conf.S3BucketName, dbFile.StorageKey, 1*time.Hour, nil)
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
