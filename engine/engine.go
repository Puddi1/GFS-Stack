package engine

import (
	"log"

	"github.com/Puddi1/GFS-Stack/env"
	"github.com/Puddi1/GFS-Stack/handlers"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

// Init_engine creates the app, the view engine and adds static files.
// Note: for dev environments the app works around the source folder, while
// for production environments it uses the vite built for performances
// and stability.
func Init_engine() *fiber.App {
	if env.ENVs["DEVELOPMENT"] == "true" {
		// Init Fiber engine and app
		app := fiber.New(fiber.Config{
			Views:             init_engine(),
			PassLocalsToViews: true,
		})

		// Loading static files (css and js) on requests
		app.Static("/~style/", "./src")
		app.Static("/~script/", "./src")
		// Loading static public files (images) on requests
		app.Static("/", "./public")

		return app
	}

	// Init Fiber engine and app
	app := fiber.New(fiber.Config{
		Views:             init_engine(),
		PassLocalsToViews: true,
	})

	// Loading static files (css and js) on requests
	app.Static("/assets", "./dist/assets")
	// Loading static public files (images) on requests
	app.Static("/", "./public")

	return app
}

// SetRoutes is the function where you set all routes of the app
func SetRoutes(app *fiber.App) error {
	stripe := app.Group("/stripe")
	api := app.Group("/api")
	_ = stripe

	// // HTML Requests // //
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("index", fiber.Map{
			"pageTitle": "GFS - Stack",

			"stringFromBackend": "Ready to ship!",
		}, "layouts/main")
	})
	app.Get("/login", func(c *fiber.Ctx) error {
		return c.Render("login/index", fiber.Map{}, "layouts/main")
	})

	// // HTMX Requests // //
	app.Get("/login/user", func(c *fiber.Ctx) error {
		return c.Render("login/user", fiber.Map{})
	})

	// // API Requests // //
	api.Get("/login/redirect", func(c *fiber.Ctx) error {
		location, status := handlers.HandleLoginWithThirdPartyOAuth(handlers.GOOGLE, "")
		return c.Redirect(location, status)
	})

	// // STRIPE Requests // //
	stripe.Post("/user/create", func(c *fiber.Ctx) error {
		// Create user
		return nil
	})

	// http://localhost:3000/#access_token=eyJhbGciOiJIUzI1NiIsImtpZCI6IkVjL214ZS9yck1rWXNIb1kiLCJ0eXAiOiJKV1QifQ.eyJhdWQiOiJhdXRoZW50aWNhdGVkIiwiZXhwIjoxNjkwMDM4NDQ1LCJpYXQiOjE2ODk2Nzg0NDUsImlzcyI6Imh0dHBzOi8vaHR0cHM6Ly9yaHV4ZG50cWJqYmhyYXd0Y2dpaC5zdXBhYmFzZS5jby9hdXRoL3YxIiwic3ViIjoiMzcyZmUzODMtNzQ0ZS00NmVlLTgyY2YtODhiYmNiYzRlYzM4IiwiZW1haWwiOiJlbGlhLnJpdmEwMUBnbWFpbC5jb20iLCJwaG9uZSI6IiIsImFwcF9tZXRhZGF0YSI6eyJwcm92aWRlciI6Imdvb2dsZSIsInByb3ZpZGVycyI6WyJnb29nbGUiXX0sInVzZXJfbWV0YWRhdGEiOnsiYXZhdGFyX3VybCI6Imh0dHBzOi8vbGgzLmdvb2dsZXVzZXJjb250ZW50LmNvbS9hL0FBY0hUdGQ5bzZfU3NIbmlHZ3htM08xVlJ3RER1S0txbUhrR3B6OW1kaW5CSlJLUjdUZz1zOTYtYyIsImVtYWlsIjoiZWxpYS5yaXZhMDFAZ21haWwuY29tIiwiZW1haWxfdmVyaWZpZWQiOnRydWUsImZ1bGxfbmFtZSI6IkVsaWEgUml2YSIsImlzcyI6Imh0dHBzOi8vYWNjb3VudHMuZ29vZ2xlLmNvbSIsIm5hbWUiOiJFbGlhIFJpdmEiLCJwaWN0dXJlIjoiaHR0cHM6Ly9saDMuZ29vZ2xldXNlcmNvbnRlbnQuY29tL2EvQUFjSFR0ZDlvNl9Tc0huaUdneG0zTzFWUndERHVLS3FtSGtHcHo5bWRpbkJKUktSN1RnPXM5Ni1jIiwicHJvdmlkZXJfaWQiOiIxMTQ5NDk2MTg0MzQwNDI5NzYzOTIiLCJzdWIiOiIxMTQ5NDk2MTg0MzQwNDI5NzYzOTIifSwicm9sZSI6ImF1dGhlbnRpY2F0ZWQiLCJhYWwiOiJhYWwxIiwiYW1yIjpbeyJtZXRob2QiOiJvYXV0aCIsInRpbWVzdGFtcCI6MTY4OTY3ODQ0NX1dLCJzZXNzaW9uX2lkIjoiN2Y5ZDJlZDctOTY2NC00ODEyLTliMTMtNDQyNGJkODZjOGMxIn0.yFuviz8ndfGaUpSA8Kx6IvIbcl65Oi0bVvAw1DvKyj4&expires_in=360000&provider_token=ya29.a0AbVbY6Nbqbs3aJaLoxgwVv4gRVd9YjVaLRQrG8B11NC0hbN-QPes1Fl7gYyWkRAWnFYRvbY1MZyNBqC4h7T6XVoJO43EApCeVIAPedsHfGPu1hzl7qjziz41iKLDL_1qZ2YbQQQGrUQAy-K2lybQ84RwqMI7aCgYKAf8SARESFQFWKvPl3ExkxlQHplByA1gAiQeMKA0163&refresh_token=NM8kpIrGK_mBg42C1ZlA_Q&token_type=bearer

	return nil
}

// Listen makes the app listen to a port defined in the env variable, default is 3000
func Listen(app *fiber.App) {
	addr := func() string {
		if env.ENVs["PORT"] == "" {
			return "3000"
		}
		return env.ENVs["PORT"]
	}()
	log.Fatal(app.Listen(":" + addr))
}

func init_engine() *html.Engine {
	if env.ENVs["DEVELOPMENT"] == "true" {
		// Reload fiber templlates
		engine := html.New("./src", ".html")
		engine.Reload(true)
		return engine
	}
	// Static dist files
	engine := html.New("./dist", ".html")
	return engine
}
