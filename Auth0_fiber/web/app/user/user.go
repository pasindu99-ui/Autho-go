package user

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

func Handler(ctx *fiber.Ctx) error {
	session, err := session.New().Get(ctx)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	profile := session.Get("profile")

	return ctx.Render("web/template/user.html", profile)
}
