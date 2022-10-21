package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

func renderLogin(c *fiber.Ctx) error {
	return c.Render("login", nil)
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

	logrus.Info(l)

	return nil
}

func MountAuth(app *fiber.App) {
	api := app.Group("/auth")
	api.Get("/login", renderLogin)
	api.Post("/login", doLogin)
}
