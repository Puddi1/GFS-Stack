package engine

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sujit-baniya/flash"
)

// All GET request that have to return a Hypertext response
func htmlRequest(r fiber.Router) {
	// index
	r.Get("/", func(c *fiber.Ctx) error {
		// Get values if there is any flash redirection
		vData := flash.Get(c)
		// Actions

		vData["stringFromBackend"] = "Ready to ship!"

		// Render
		vData["pageTitle"] = "GFS - Stack"
		return c.Render("index", vData, "layouts/main")
	})

	// signin
	r.Get("/signin", func(c *fiber.Ctx) error {
		// Get values if there is any flash redirection
		vData := flash.Get(c)
		// Actions

		// ...

		// Render
		vData["pageTitle"] = "GFS - Signin"
		return c.Render("signin/index", vData, "layouts/main")
	})

	// signup
	r.Get("/signup", func(c *fiber.Ctx) error {
		// Get values if there is any flash redirection
		vData := flash.Get(c)
		// Actions

		// ...

		// Render
		vData["pageTitle"] = "GFS - Signup"
		return c.Render("signup/index", vData, "layouts/main")
	})

	// dashboard
	r.Get("/dashboard", onlyUser(func(c *fiber.Ctx) error {
		// Get values if there is any flash redirection
		vData := flash.Get(c)
		// Actions

		vData["Email"] = "Verified Email"
		vData["Count"] = "0"

		// Render
		vData["pageTitle"] = "GFS - Dashboard"
		return c.Render("dashboard/index", vData, "layouts/main")
	}))

	// admin (ex filtered request)
	r.Get("/admin", onlyAdmin(func(c *fiber.Ctx) error { return nil }))
}
