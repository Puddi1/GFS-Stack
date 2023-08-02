package engine

import (
	"github.com/Puddi1/GFS-Stack/handlers"
	"github.com/gofiber/fiber/v2"
)

// handleError is used to redirect users to a default error page
func handleError(errorCode *fiber.Error, errorComment string) *fiber.Error {
	errInfHandler.sw.Wait()
	errInfHandler.sw.Add(1)
	errInfHandler.errorComment = errorComment
	return errorCode
}

// Routes filters
// onlyAdmin allows only admin, redirect to error page
func onlyAdmin(fn fiber.Handler) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Gatekeeping logic

		// Check if use is an Admin
		shallNotPass := true
		if shallNotPass {
			handleError(fiber.ErrForbidden, "This is an only-Admin page")
		}

		// Passed the gatekeeper
		return fn(c)
	}
}

// onlyUser allows only users, redirect to signup
func onlyUser(fn fiber.Handler) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Gatekeeping logic

		// Check
		isUser := true
		if !isUser {
			rf := handlers.NewRedirectFlash(c, fiber.Map{}, "/signup")
			return handlers.HandleRedirectWithFlash(rf, handlers.WithNotifyAlert(handlers.AlertWarning, "Not A User", "Sign up to use the app"))
		}

		// Pased the gatekeeper
		return fn(c)
	}
}
