package engine

import (
	"errors"
	"fmt"
	"log"

	"github.com/Puddi1/GFS-Stack/env"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

type ErrorSpecifications struct {
	// ecc to create custom err messages on the fly
}

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
			ErrorHandler:      handleErrors(),
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
		ErrorHandler:      handleErrors(),
	})

	// Loading static files (css and js) on requests
	app.Static("/assets", "./dist/assets")
	// Loading static public files (images) on requests
	app.Static("/", "./public")

	return app
}

// Listen makes the app listen to a port defined in the env variable, default is 3000
func Listen(app *fiber.App) {
	addr := func() string {
		if env.ENVs["PORT"] == "" {
			return "3000"
		}
		return env.ENVs["PORT"]
	}()
	if err := app.Listen(":" + addr); err != nil {
		log.Panic(err)
	}
}

// Initialize the correct engine
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

// handle Fiber Errors
func handleErrors() func(*fiber.Ctx, error) error {
	return func(c *fiber.Ctx, err error) error {
		// Status code defaults to 500
		code := fiber.StatusInternalServerError
		// Retrieve the custom status code if it's a *fiber.Error
		var e *fiber.Error
		if errors.As(err, &e) {
			code = e.Code
		}

		// Custom error actions
		switch code {
		case 403:
			log.Println("Error 403 triggered")
		case 404:
			log.Println("Error 404 triggered")
		}

		// Render error page
		log.Println("Rendering page")
		name := fmt.Sprintf("errors/%d", code)
		pageTitle := fmt.Sprintf("GFS Stack - %d error", code)
		errRender := c.Render(name, fiber.Map{
			"pageTitle": pageTitle,
		}, "layouts/main")

		if errRender != nil {
			// handle error
		}

		return nil
	}
}
