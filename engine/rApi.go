package engine

import (
	"fmt"
	"time"

	"github.com/Puddi1/GFS-Stack/env"
	"github.com/Puddi1/GFS-Stack/handlers"
	"github.com/Puddi1/GFS-Stack/utils"
	"github.com/gofiber/fiber/v2"
)

// All requests to the API
func apiRequest(r fiber.Router) {
	// Signup
	r.Post("/signup/email", func(c *fiber.Ctx) error {
		// Actions
		// Parse the body request
		b := new(handlers.UpdateUserBody)
		if err := c.BodyParser(b); err != nil {
			return err
		}

		// Request check
		_, isEmail := utils.ValidMailAddress(b.Email)
		if !isEmail || len([]rune(b.Password)) < 6 {
			rf := handlers.NewRedirectFlash(c, fiber.Map{
				"Email": b.Email,
			}, "/signup")

			if !isEmail && len([]rune(b.Password)) < 6 {
				return handlers.HandleRedirectWithFlash(rf, handlers.WithNotifyAlert(handlers.AlertError, "Error Signing up", "Email is not valid and password must be at least 6 characters long"))
			} else if !isEmail {
				return handlers.HandleRedirectWithFlash(rf, handlers.WithNotifyAlert(handlers.AlertError, "Error Signing up", "Not a valid email"))
			} else if len([]rune(b.Password)) < 6 {
				return handlers.HandleRedirectWithFlash(rf, handlers.WithNotifyAlert(handlers.AlertError, "Error Signing up", "Password must be at least 6 characters long"))
			}
		}

		// Handle request
		err := handlers.HandleSignUpUserWithEmail(b.Email, b.Password)
		if err != nil {
			fmt.Printf("Error during user signup: %e", err)
		}

		time.Sleep(5 * time.Second) // demonstrative of htmx rotating indicator
		// Response redirect to dashboard
		return nil
	})
	r.Get("/signup/OAuth/:provider", func(c *fiber.Ctx) error {
		// Actions
		redirectUrl := "" + env.ENVs["APP_URL"] + "/dashboard"
		location, status := handlers.HandleLoginWithThirdPartyOAuth(handlers.OAuth[c.Params("provider")], redirectUrl)
		// Render
		return c.Redirect(location, status)
	})

	// Signin
	r.Post("/signin/email", func(c *fiber.Ctx) error {
		// Actions
		b := new(handlers.UpdateUserBody)
		if err := c.BodyParser(b); err != nil {
			return err
		}
		res, err := handlers.HandleLoginUserWithEmail(b.Email, b.Password)
		_ = res
		if err != nil {
			fmt.Printf("during user signin: %e", err)
		}

		time.Sleep(5 * time.Second) // demonstrative of htmx rotating indicator
		// Response
		return nil
	})
	r.Get("/signin/OAuth/:provider", func(c *fiber.Ctx) error {
		// Actions
		redirectUrl := "" + env.ENVs["APP_URL"] + "/dashboard"
		location, status := handlers.HandleLoginWithThirdPartyOAuth(handlers.OAuth[c.Params("provider")], redirectUrl)
		// Render
		return c.Redirect(location, status)
	})
}
