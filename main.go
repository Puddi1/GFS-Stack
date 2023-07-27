package main

import (
	"fmt"
	"log"

	"github.com/Puddi1/GFS-Stack/database"
	"github.com/Puddi1/GFS-Stack/engine"
	"github.com/Puddi1/GFS-Stack/env"
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

	// Actions test

	// Listen on PORT
	engine.Listen(app)
}
