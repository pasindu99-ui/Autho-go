package router

import (
	"AUTH0_FIBER/platform/authenticator"
	"AUTH0_FIBER/web/app/callback"
	"AUTH0_FIBER/web/app/home"
	logout "AUTH0_FIBER/web/app/logOut"
	"AUTH0_FIBER/web/app/login"
	"AUTH0_FIBER/web/app/user"
	"encoding/gob"

	"github.com/gofiber/fiber/v2"
	//"github.com/gofiber/fiber/v2/middleware/encryptcookie"
)

// New registers the routes and returns the router.
func New(app *fiber.App, auth *authenticator.Authenticator) *fiber.App {
	// To store custom types in our cookies,
	// we must first register them using gob.Register
	gob.Register(map[string]interface{}{})

	// store := encryptcookie.New(encryptcookie.Config{
	// 	auth-session: "secret",
	// })
	// fmt.Println("store: ", store)

	// // Use the session middleware.
	// app.Use(store)

	app.Static("/public", "/router")

	app.Get("/", home.Handler)
	app.Get("/login", login.Handler(auth))
	app.Get("/callback", callback.Handler(auth))
	app.Get("/user", user.Handler)
	app.Get("/logout", logout.Handler)

	return app
}
