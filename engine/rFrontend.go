package engine

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/sujit-baniya/flash"
)

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
		log.Println("1")
		m := flash.Get(c)
		log.Println("2")
		m["pageTitle"] = "GFS - Signup"
		log.Println("3")
		log.Println(m)
		// ...
		// Render
		return c.Render("signup/index", m, "layouts/main")
	})

	// admin (ex forbidden request)
	r.Get("/admin", onlyAdmin(func(c *fiber.Ctx) error { return nil }))
}
