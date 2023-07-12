package logout

import (
	"net/http"
	"net/url"
	"os"

	"github.com/gofiber/fiber/v2"
)

// Handler for our logout.

func Handler(ctx *fiber.Ctx) error {
	logoutUrl, err := url.Parse("https://" + os.Getenv("AUTH0_DOMAIN") + "/v2/logout")
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).SendString(err.Error())
	}

	scheme := "http"
	if ctx.Protocol() == "https" {
		scheme = "https"
	}

	returnTo, err := url.Parse(scheme + "://" + ctx.Hostname())
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).SendString(err.Error())
	}

	parameters := url.Values{}
	parameters.Add("returnTo", returnTo.String())
	parameters.Add("client_id", os.Getenv("AUTH0_CLIENT_ID"))
	logoutUrl.RawQuery = parameters.Encode()

	return ctx.Redirect(logoutUrl.String(), http.StatusTemporaryRedirect)
}
