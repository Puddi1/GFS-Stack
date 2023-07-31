package engine

import (
	"fmt"
	"time"

	"github.com/Puddi1/GFS-Stack/env"
	"github.com/Puddi1/GFS-Stack/handlers"
	"github.com/gofiber/fiber/v2"
)

// All requests to the API
func apiRequest(r fiber.Router) {
	// Signup
	r.Post("/signup/email", func(c *fiber.Ctx) error {
		// Actions
		b := new(handlers.UpdateUserBody)
		if err := c.BodyParser(b); err != nil {
			return err
		}
		err := handlers.HandleSignUpUserWithEmail(b.Email, b.Password)
		if err != nil {
			fmt.Printf("during user signup: %e", err)
		}

		time.Sleep(5 * time.Second) // demonstrative of htmx rotating indicator
		// Response
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
