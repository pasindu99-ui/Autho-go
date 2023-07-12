package router

import (
	"AUTH0_FIBER/platform/authenticator"
	"AUTH0_FIBER/web/app/callback"
	logout "AUTH0_FIBER/web/app/logOut"
	"AUTH0_FIBER/web/app/login"
	"AUTH0_FIBER/web/app/user"
	"encoding/gob"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/session/v2"
	"github.com/gofiber/session/v2/provider/cookie"
)

// New registers the routes and returns the router.
func New(auth *authenticator.Authenticator) *fiber.App {
	app := fiber.New()

	// To store custom types in our cookies,
	// we must first register them using gob.Register
	gob.Register(map[string]interface{}{})

	store := cookie.New(cookie.Config{
		Key: []byte("secret"),
	})
	session := session.New(session.Config{
		Provider: store,
	})
	app.Use(sessionware.New(session))

	app.Static("/public", "web/static")

	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.Render("home", nil)
	})
	app.Get("/login", login.Handler(auth))
	app.Get("/callback", callback.Handler(auth))
	app.Get("/user", user.Handler)
	app.Get("/logout", logout.Handler)

	return app
}
