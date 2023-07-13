package callback

import (
	"AUTH0_FIBER/platform/authenticator"
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

// Handler for our callback.
func Handler(auth *authenticator.Authenticator) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		sess, err := session.New().Get(ctx)
		keys := sess.Keys()
		fmt.Println("keys", keys)
		fmt.Println("sess", sess)
		if err != nil {
			panic(err)
		}
		// Get value
		state := sess.Get("state")
		fmt.Println("stateweeeeeeeeeeeee", state)
		if ctx.Query("state") != state {
			return ctx.Status(http.StatusBadRequest).SendString("Invalid state parameter.")
		}

		// Exchange an authorization code for a token.
		token, err := auth.Exchange(ctx.Context(), ctx.Query("code"))
		if err != nil {
			return ctx.Status(http.StatusUnauthorized).SendString("Failed to exchange an authorization code for a token.")
		}

		idToken, err := auth.VerifyIDToken(ctx.Context(), token)
		if err != nil {
			return ctx.Status(http.StatusInternalServerError).SendString("Failed to verify ID Token.")
		}

		var profile map[string]interface{}
		if err := idToken.Claims(&profile); err != nil {
			return ctx.Status(http.StatusInternalServerError).SendString(err.Error())
		}

		sess.Set("access_token", token.AccessToken)
		sess.Set("profile", profile)
		if err := sess.Save(); err != nil {
			return ctx.Status(http.StatusInternalServerError).SendString(err.Error())
		}

		// Redirect to logged in page.
		return ctx.Redirect("/user", http.StatusTemporaryRedirect)
	}
}
