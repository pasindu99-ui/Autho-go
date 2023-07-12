package user

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

// Handler for our logged-in user page.
func Handler(ctx *fiber.Ctx) error {
	sess, err := session.New().Get(ctx)
	if err == nil {
		return fiber.NewError(fiber.StatusUnauthorized, "Session not found")
	}

	profile := sess.Get("profile")
	if profile == nil {
		return fiber.NewError(fiber.StatusUnauthorized, "User profile not found")
	}

	return ctx.Render("user", fiber.Map{
		"profile": profile,
	}, "main")
}
