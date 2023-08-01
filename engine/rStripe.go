package engine

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

// All requests to stripe backend side
func stripeRequest(r fiber.Router) {
	r.Post("/user/create", func(c *fiber.Ctx) error {
		// Create user
		return nil
	})
}

// All requests sent by stripe webhook
func stripeWebhooks(r fiber.Router) {
	// All webhook requests are handled by a signle endpoint for simplicity
	// cases not handled will return a 503 status error and a json custom message

	// add Idempotent webhooks, verify signature
	r.Post("/:source", func(c *fiber.Ctx) error {
		c.Accepts("application/json") // "application/json"

		// rb := new(stripe_gfs.)
		// c.Body()

		// Idempotent webhooks
		// if 0 != 0 {
		// Return request already seen json
		// }

		// Handle if new requests
		switch c.Params("source") {
		// case string(stripe.CheckoutSessionStatusComplete):
		// 	fmt.Printf("yay")
		// 	return nil
		default:
			fmt.Print(c.Params("source"))
			return fiber.NewError(fiber.StatusServiceUnavailable, `{"message": "Stripe webhook not supported"}`)
		}
	})
}
