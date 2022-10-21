package api

import (
	"github.com/SphericalKat/katbox/ent/user"
	"github.com/SphericalKat/katbox/internal/db"
	"github.com/gofiber/fiber/v2"
)

func index(c *fiber.Ctx) error {
	email := c.Cookies("authToken")
	if email == "" {
		return c.Redirect("/auth/login")
	}

	exists, err := db.Client.User.Query().Where(user.EmailEQ(email)).Exist(c.Context())
	if err != nil {
		return c.Redirect("/auth/login")
	}

	if !exists {
		return c.Redirect("/auth/login")
	}

	return c.Render("index", fiber.Map{
		"Title": "Katbox",
	})
}

func MountIndex(app *fiber.App) {
	app.Get("/", index)
}
