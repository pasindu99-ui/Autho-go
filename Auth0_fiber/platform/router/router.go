package router

import (
	"AUTH0_FIBER/platform/authenticator"
	"AUTH0_FIBER/web/app/callback"
	logout "AUTH0_FIBER/web/app/logOut"
	"AUTH0_FIBER/web/app/login"
	"AUTH0_FIBER/web/app/user"
	"encoding/gob"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/encryptcookie"
)

// New registers the routes and returns the router.
func New(app *fiber.App, auth *authenticator.Authenticator) *fiber.App {
	// To store custom types in our cookies,
	// we must first register them using gob.Register
	gob.Register(map[string]interface{}{})

	store := encryptcookie.New(encryptcookie.Config{
		Key: "secret",
	})
	fmt.Println("store: ", store)

	// Use the session middleware.
	app.Use(store)

	app.Static("/public", "web/static")

	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.Render("home", nil)
	})

	app.Get("/login", login.Handler(auth))
	app.Get("/callback", callback.Handler(auth))
	app.Get("/user", user.Handler)
	app.Get("/logout", logout.Handler)
	// app.Get("/", home.Handler)

	return app
}
