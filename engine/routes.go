package engine

import (
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
		// Actions
		// ...
		// Render
		return c.Render("index", fiber.Map{
			"pageTitle": "GFS - Stack",

			"stringFromBackend": "Ready to ship!",
		}, "layouts/main")
	})

	// signin
	r.Get("/signin", func(c *fiber.Ctx) error {
		// Actions
		// ...
		// Render
		return c.Render("signin/index", fiber.Map{
			"pageTitle": "GFS - Signin",
		}, "layouts/main")
	})

	// signup
	r.Get("/signup", func(c *fiber.Ctx) error {
		// Actions
		// ...
		// Render
		return c.Render("signup/index", fiber.Map{
			"pageTitle": "GFS - Signup",
		}, "layouts/main")
	})

	// admin (ex forbidden request)
	r.Get("/admin", onlyAdmin(func(c *fiber.Ctx) error { return nil }))
}
