package main

import (
	"fmt"
	"log"

	"github.com/Puddi1/GFS-Stack/database"
	"github.com/Puddi1/GFS-Stack/engine"
	"github.com/Puddi1/GFS-Stack/env"
	"github.com/Puddi1/GFS-Stack/handlers"
	"github.com/Puddi1/GFS-Stack/stripe_gfs"
)

// func gracefulShutdown() {

// }

func main() {
	// Execution started
	fmt.Println("\nRun server")
	// Init env
	env.Init_env()

	if env.ENVs["DEVELOPMENT"] == "true" {
		// Templates are based on the src directory and reload is managed by Fiber
		fmt.Println("\nDevelopment Environment")
	}

	// Init db
	database.Init_db()
	// Init stripe
	stripe_gfs.Init_stripe()
	// Init fiber
	app := engine.Init_engine()
	err := engine.SetRoutes(app)
	if err != nil {
		log.Fatalf("Error occured during routes creation: %s", err)
	}

	// database.DB.AutoMigrate(&data.Subscription{})
	// database.DB.Create(&data.Subscription{Free: 1, Basic: 2, Enterprise: 5})
	// database.DB.Create(&data.Subscription{Free: 2, Basic: 134, Enterprise: 593})
	// database.DB.Create(&data.Subscription{Free: 3, Basic: 1342, Enterprise: 556})

	// database.SignUpUserWithEmail("signup@user.com", "passordSuperSecure")

	// s, err := handlers.HandleCheckoutSessionCreation(
	// 	&stripe.CheckoutSessionParams{
	// 		SuccessURL: stripe.String("https://pizza.com"),
	// 		Mode:       stripe.String(stripe_gfs.PAYMENT),
	// 		LineItems: []*stripe.CheckoutSessionLineItemParams{{
	// 			Price:    stripe.String("price_1NXoBZHqfgifxJ03TMZLKskv"),
	// 			Quantity: stripe.Int64(2),
	// 		}},
	// 	},
	// )
	// _ = err

	// fmt.Println(s)
	s, _ := handlers.HandleCustomerPortalSessionCreation("cus_OKTgoI3cQYJLVV", "https://youtube.com")
	fmt.Print(s)

	// Listen on PORT
	engine.Listen(app)
}
