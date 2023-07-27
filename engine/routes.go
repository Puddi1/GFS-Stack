package engine

import (
	"fmt"

	"github.com/Puddi1/GFS-Stack/handlers"
	"github.com/gofiber/fiber/v2"
)

// SetRoutes is the function where you set all routes of the app
func SetRoutes(app *fiber.App) error {
	api := app.Group("/api")
	db := api.Group("/database")
	s := api.Group("/stripe")
	s_we := s.Group("/webhook_events")

	// // HTML Requests // //
	htmlRequest(app)
	// // HTMX Requests // //
	htmxRequest(app)
	// // API Requests // //
	apiRequest(api)
	// // DATABASE Requests // //
	databaseRequest(db)
	// // STRIPE Requests // //
	stripeRequest(s)
	// // STRIPE Webhooks - very WIP // //
	stripeWebhooks(s_we)

	return nil
}

func htmlRequest(r fiber.Router) {
	// index
	r.Get("/", func(c *fiber.Ctx) error {
		return c.Render("index", fiber.Map{
			"pageTitle": "GFS - Stack",

			"stringFromBackend": "Ready to ship!",
		}, "layouts/main")
	})
	// login
	r.Get("/login", func(c *fiber.Ctx) error {
		return c.Render("login/index", fiber.Map{}, "layouts/main")
	})
}

func htmxRequest(r fiber.Router) {
	// login/user
	r.Get("/login/user", func(c *fiber.Ctx) error {
		return c.Render("login/user", fiber.Map{})
	})
}

func apiRequest(r fiber.Router) {
	r.Get("/login/redirect", func(c *fiber.Ctx) error {
		location, status := handlers.HandleLoginWithThirdPartyOAuth(handlers.GOOGLE, "")
		return c.Redirect(location, status)
	})
}

func databaseRequest(r fiber.Router) {
	r.Get("/", func(c *fiber.Ctx) error {
		// handle db
		return nil
	})
}

func stripeRequest(r fiber.Router) {
	r.Post("/user/create", func(c *fiber.Ctx) error {
		// Create user
		return nil
	})
}

func stripeWebhooks(r fiber.Router) {
	// All webhook requests are handled by a signle endpoint for simplicity
	// cases not handled will return a 503 status error and a json custom message

	// add Idempotent webhooks, verify signature
	r.Post("/:source", func(c *fiber.Ctx) error {
		c.Accepts("application/json") // "application/json"

		// rb := new(stripe_gfs.)
		// c.Body()

		// Idempotent webhooks
		// if 0 != 0 {
		// Return request already seen json
		// }

		// Handle if new requests
		switch c.Params("source") {
		// case string(stripe.CheckoutSessionStatusComplete):
		// 	fmt.Printf("yay")
		// 	return nil
		default:
			fmt.Print(c.Params("source"))
			return fiber.NewError(fiber.StatusServiceUnavailable, `{"message": "Stripe webhook not supported"}`)
		}
	})
}
