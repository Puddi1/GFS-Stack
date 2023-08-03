package engine

import (
	"log"

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

// redirectUser checks if user is authenticated and, if not, redirects him
func redirectUser(fn fiber.Handler) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Gatekeeping logic
		isUser := handlers.CheckJWT(c)

		if isUser {
			rf := handlers.NewRedirectFlash(c, fiber.Map{}, "/dashboard")
			return handlers.HandleRedirectWithFlash(rf, handlers.WithNotifyAlert(handlers.AlertSuccess, "User Already Signed In", "Your session has been restores"))
		}

		return fn(c)
	}
}

// onlyUser allows only users, redirect to signup
func onlyUserFill(fn func(*fiber.Ctx, handlers.User) error) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Gatekeeping logic
		jwt, isUser := handlers.CheckFillJWT(c)
		if isUser != nil {

			rf := handlers.NewRedirectFlash(c, fiber.Map{}, "/signup")
			return handlers.HandleRedirectWithFlash(rf, handlers.WithNotifyAlert(handlers.AlertWarning, "Not A User", "Sign up to use the app"))
		}

		// Create use struct to update UI
		User := handlers.User{
			JWT: jwt,
		}

		res, err := handlers.GetUser(User.JWT.Access_token)
		if err != nil {

			rf := handlers.NewRedirectFlash(c, fiber.Map{}, "/signup")
			return handlers.HandleRedirectWithFlash(rf, handlers.WithNotifyAlert(handlers.AlertWarning, "Not A User", "Sign up to use the app"))
		}

		// Add data to user
		log.Println(*res)
		resB, _ := handlers.HandleResponseBodyToString(res)
		log.Println(resB)
		_ = res

		// Pased the gatekeeper
		return fn(c, User)
	}
}
