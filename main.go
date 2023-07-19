package main

import (
	"fmt"
	"log"

	"github.com/Puddi1/GFS-Stack/database"
	"github.com/Puddi1/GFS-Stack/engine"
	"github.com/Puddi1/GFS-Stack/env"
	"github.com/Puddi1/GFS-Stack/stripe"
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
	stripe.Init_stripe()
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

	// Listen on PORT
	engine.Listen(app)
}
