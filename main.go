package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/Puddi1/GFS-Stack/database"
	"github.com/Puddi1/GFS-Stack/engine"
	"github.com/Puddi1/GFS-Stack/env"
	"github.com/Puddi1/GFS-Stack/stripe_gfs"
	"github.com/gofiber/fiber/v2"
)

func Init_GracefulShutdown(app *fiber.App) *sync.WaitGroup {
	// Creating channel that will be notified
	c := make(chan os.Signal, 1)
	// Adding signals to listen to for shutdown
	signal.Notify(c,
		os.Interrupt,
		syscall.SIGTERM,
		syscall.SIGINT,
	)
	// Creating sync group
	var serverShutdown sync.WaitGroup
	go func() {
		// Waiting for channel signal
		<-c // When using go 1.15 or older
		fmt.Println("Gracefully shutting down...")
		// Adding timeout logic: creating a sync that is waiting
		// in the main function for this defer to run after the timeout
		// is completed and the app is shut down.
		serverShutdown.Add(1)
		defer serverShutdown.Done()
		_ = app.ShutdownWithTimeout(60 * time.Second)
	}()

	return &serverShutdown
}

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
	// Init graceful shutdown
	serverShutdown := Init_GracefulShutdown(app)

	// ...

	// Listen on PORT
	engine.Listen(app)
	// Graceful Cleanup
	serverShutdown.Wait()
	fmt.Println("Running cleanup tasks...")
	// Your cleanup tasks go here
	// ...
}
