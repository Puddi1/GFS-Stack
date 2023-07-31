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
			errInfHandler.sw.Wait()
			errInfHandler.sw.Add(1)
			errInfHandler.errorComment = "This is an only-Admin page"
			return fiber.ErrForbidden
		}

		// Passed the gatekeeper
		return fn(c)
	}
}
