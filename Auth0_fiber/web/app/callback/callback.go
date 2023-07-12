package callback

import (
	"AUTH0_FIBER/platform/authenticator"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

// Handler for our callback.
func Handler(auth *authenticator.Authenticator) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		session, err := session.New().Get(ctx)
		if ctx.Query("state") != session.Get("state") {
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

		session.Set("access_token", token.AccessToken)
		session.Set("profile", profile)
		if err := session.Save(); err != nil {
			return ctx.Status(http.StatusInternalServerError).SendString(err.Error())
		}

		// Redirect to logged in page.
		return ctx.Redirect("/user", http.StatusTemporaryRedirect)
	}
}
