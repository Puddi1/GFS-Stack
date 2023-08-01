package engine

import (
	"github.com/gofiber/fiber/v2"
)

// SetRoutes is the function where you set all routes of the app
func SetRoutes(app *fiber.App) error {
	rApi := app.Group("/api")
	rDb := rApi.Group("/database")
	rStripe := rApi.Group("/stripe")
	rStripeWe := rStripe.Group("/webhook_events")

	// // HTML Requests "url/..." // //
	htmlRequest(app)
	// // API Requests "url/api/..." // //
	apiRequest(rApi)
	// // DATABASE Requests "url/api/databse/..." // //
	databaseRequest(rDb)
	// // STRIPE Requests "url/api/stripe/..." // //
	stripeRequest(rStripe)
	// // STRIPE Webhooks "url/api/stripe/webhook_events/..." // //
	stripeWebhooks(rStripeWe)

	return nil
}
