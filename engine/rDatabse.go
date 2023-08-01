package engine

import "github.com/gofiber/fiber/v2"

// All direct requests to the database backend side, needed?
func databaseRequest(r fiber.Router) {
	r.Get("/", onlyAdmin(func(c *fiber.Ctx) error {
		// handle db
		return nil
	}))
}
