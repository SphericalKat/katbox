package api

import (
	"github.com/SphericalKat/katbox/ent/user"
	"github.com/SphericalKat/katbox/internal/db"
	"github.com/SphericalKat/katbox/internal/utils"
	"github.com/gofiber/fiber/v2"
)

func renderLogin(c *fiber.Ctx) error {
	return c.Render("login", fiber.Map{
		"Error": nil,
	})
}

type LoginReq struct {
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

func doLogin(c *fiber.Ctx) error {
	l := &LoginReq{}
	if err := c.BodyParser(l); err != nil {
		return err
	}

	user, err := db.Client.User.Query().Where(
		user.EmailEQ(l.Email),
	).Only(c.Context())
	if err != nil {
		return c.Render("login", fiber.Map{
			"Error": err,
		})
	}
	if user == nil {
		return c.Render("login", fiber.Map{
			"Error": "Incorrect email or password.",
		})
	}

	match, err := utils.ComparePassword(l.Password, user.Password)
	if err != nil {
		return c.Render("login", fiber.Map{
			"Error": err,
		})
	}
	if !match {
		return c.Render("login", fiber.Map{
			"Error": "Incorrect email or password.",
		})
	}

	c.Cookie(&fiber.Cookie{
		Name:  "authToken",
		Value: l.Email,
	})

	return c.Redirect("/")
}

func MountAuth(app *fiber.App) {
	api := app.Group("/auth")
	api.Get("/login", renderLogin)
	api.Post("/login", doLogin)
}
