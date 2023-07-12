package middleware

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/session/v2"
)

// IsAuthenticated is a middleware that checks if
// the user has already been authenticated previously.
func IsAuthenticated(ctx *fiber.Ctx) error {
	sess := session.New().Get(ctx)
	if sess == nil {
		return ctx.Status(http.StatusInternalServerError).SendString("Session not found.")
	}

	if sess.Get("profile") == nil {
		return ctx.Redirect("/", http.StatusSeeOther)
	}

	return ctx.Next()
}
