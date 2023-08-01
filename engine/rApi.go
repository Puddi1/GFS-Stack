package engine

import (
	"fmt"
	"log"
	"time"

	"github.com/Puddi1/GFS-Stack/env"
	"github.com/Puddi1/GFS-Stack/handlers"
	"github.com/Puddi1/GFS-Stack/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/sujit-baniya/flash"
)

// All requests to the API
func apiRequest(r fiber.Router) {
	// Signup
	r.Post("/signup/email", func(c *fiber.Ctx) error {
		// Actions
		// Parse the body request
		log.Println("Request started")
		b := new(handlers.UpdateUserBody)
		if err := c.BodyParser(b); err != nil {
			return err
		}

		log.Println("Body request: ", b)

		// Request check
		_, isEmail := utils.ValidMailAddress(b.Email)
		if !isEmail || len([]rune(b.Password)) < 6 {
			log.Println("If values: ", !isEmail, len([]rune(b.Password)))

			// rf := handlers.NewRedirectFlash(c, fiber.Map{
			// 	"Email":    b.Email,
			// 	"Password": b.Password,
			// }, "/signup")

			// var na handlers.NotifyAlert

			// if !isEmail && len([]rune(b.Password)) < 6 {
			// 	na = handlers.NewNotifyAlert(handlers.AlertError, "Error Signing up", "Email is not valid and password must be at least 6 characters long")
			// } else if !isEmail {
			// 	na = handlers.NewNotifyAlert(handlers.AlertError, "Error Signing up", "Not a valid email")
			// } else if len([]rune(b.Password)) < 6 {
			// 	na = handlers.NewNotifyAlert(handlers.AlertError, "Error Signing up", "Password must be at least 6 characters long")
			// }

			// log.Println("Redirect flash: ", rf)

			// rf.Data["test"] = "testValue"

			// return handlers.HandleRedirectWithFlash(rf, handlers.WithNotifyAlert(na))
			m := fiber.Map{
				"test":  "this is a test",
				"Email": "this is a test",
				"NotifyAlert": struct {
					AlertType    string
					AlertTitle   string
					AlertMessage string
				}{"AlertSuccess", "title error", "message error"},
			}

			log.Println(m)

			return flash.WithData(c, m).Redirect("/signup")
		}

		// Handle request
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
