package home

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

// Handler for our home page.
func Handler(ctx *fiber.Ctx) {
	ctx.Status(http.StatusOK).Render("home", nil)
}
