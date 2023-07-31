package engine

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

// Routes filters
func onlyAdmin(fn fiber.Handler) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Gatekeeping logic
		fmt.Print("Gatekeeping checks")
		if true {
			return fiber.ErrForbidden
		}

		// Passed the gatekeeper
		return fn(c)
	}
}
