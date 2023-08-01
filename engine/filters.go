package engine

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func handleError(errorCode *fiber.Error, errorComment string) *fiber.Error {
	errInfHandler.sw.Wait()
	errInfHandler.sw.Add(1)
	errInfHandler.errorComment = errorComment
	return errorCode
}

// Routes filters
func onlyAdmin(fn fiber.Handler) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Gatekeeping logic
		fmt.Print("Gatekeeping checks")
		if true {
			handleError(fiber.ErrForbidden, "This is an only-Admin page")
		}

		// Passed the gatekeeper
		return fn(c)
	}
}
