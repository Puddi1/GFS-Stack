package engine

import (
	"fmt"
	"net/http"
	"time"

	"github.com/Puddi1/GFS-Stack/env"
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

// All GET request that have to return a Hypertext response
func htmlRequest(r fiber.Router) {
	// index
	r.Get("/", func(c *fiber.Ctx) error {
		return c.Render("index", fiber.Map{
			"pageTitle": "GFS - Stack",

			"stringFromBackend": "Ready to ship!",
		}, "layouts/main")
	})
	// signin
	r.Get("/signin", func(c *fiber.Ctx) error {
		return c.Render("signin/index", fiber.Map{
			"pageTitle": "GFS - Signin",
		}, "layouts/main")
	})
	// signup
	r.Get("/signup", func(c *fiber.Ctx) error {
		return c.Render("signup/index", fiber.Map{
			"pageTitle": "GFS - Signup",
		}, "layouts/main")
	})
}

// All requests to the API
func apiRequest(r fiber.Router) {
	r.Get("/signup/OAuth/:provider", func(c *fiber.Ctx) error {
		switch c.Params("provider") {
		case "google":
			redirectUrl := "" + env.ENVs["APP_URL"] + "/dashboard"
			location, status := handlers.HandleLoginWithThirdPartyOAuth(handlers.GOOGLE, redirectUrl)
			return c.Redirect(location, status)
		default:
			return nil
		}
	})
	r.Get("/signin/OAuth/:provider", func(c *fiber.Ctx) error {
		switch c.Params("provider") {
		case "google":
			redirectUrl := "" + env.ENVs["APP_URL"] + "/dashboard"
			location, status := handlers.HandleLoginWithThirdPartyOAuth(handlers.GOOGLE, redirectUrl)
			return c.Redirect(location, status)
		default:
			return nil
		}
	})
	r.Post("/signup/email", func(c *fiber.Ctx) error {
		b := new(handlers.UpdateUserBody)
		if err := c.BodyParser(b); err != nil {
			return err
		}
		err := handlers.HandleSignUpUserWithEmail(b.Email, b.Password)
		if err != nil {
			fmt.Printf("during user signup: %e", err)
		}

		time.Sleep(5 * time.Second)
		return nil
	})
	r.Post("/signin/email", func(c *fiber.Ctx) error {
		b := new(handlers.UpdateUserBody)
		if err := c.BodyParser(b); err != nil {
			return err
		}
		res, err := handlers.HandleLoginUserWithEmail(b.Email, b.Password)
		_ = res
		if err != nil {
			fmt.Printf("during user signin: %e", err)
		}

		time.Sleep(5 * time.Second)
		return nil
	})
}

// All direct requests to the database backend side, needed?
func databaseRequest(r fiber.Router) {
	r.Get("/", onlyAdmin(func(c *fiber.Ctx) error {
		// handle db
		return nil
	}))
}

// All requests to stripe backend side
func stripeRequest(r fiber.Router) {
	r.Post("/user/create", func(c *fiber.Ctx) error {
		// Create user
		return nil
	})
}

// All requests sent by stripe webhook
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

// Routes filters
func onlyAdmin(fn fiber.Handler) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Gatekeeping logic
		fmt.Print("Gatekeeping checks")
		a := false
		if a {
			return c.SendStatus(http.StatusUnauthorized)
		}

		// Passed the gatekeeper
		return fn(c)
	}
}
