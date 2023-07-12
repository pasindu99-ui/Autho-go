package home

import (
	"github.com/gofiber/fiber/v2"
)

// Handler for our home page.
func Handler(ctx *fiber.Ctx) error {
	return ctx.Render("web/template/home.html", nil)
}
