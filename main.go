package main

import (
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
	"github.com/Puddi1/GFS-Stack/utils"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/exp/slog"
)

func Init_GracefulShutdown(app *fiber.App) *sync.WaitGroup {
	// Creating channel that will be notified
	c := make(chan os.Signal, 1)
	// Adding signals to listen to for shutdown
	signal.Notify(c,
		os.Interrupt,
		syscall.SIGTERM,
		syscall.SIGINT,
		syscall.SIGSEGV,
	)
	// Creating sync group
	var serverShutdown sync.WaitGroup
	go func() {
		// Waiting for channel signal
		<-c // When using go 1.15 or older
		log.Println("Gracefully shutting down...")
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
	log.Println("Run server")
	// Init env
	env.Init_env()
	if env.ENVs["DEVELOPMENT"] == "true" {
		// Templates are based on the src directory and reload is managed by Fiber
		log.Println("Development Environment")
	}
	// Init Logger
	errLog := utils.Init_LoggerGFS(env.ENVs["LOG_FILE_PATH"], env.ENVs["WRITE_LOGS"])
	if errLog != nil {
		log.Panicf("Logger not initialized correctly: %e", errLog)
	}
	lg := slog.Default()
	lg.Info("Log initialized, log instance: ", lg)
	// Init db
	database.Init_db()
	// Init stripe
	stripe_gfs.Init_stripe()
	// Init fiber
	app := engine.Init_engine()
	err := engine.SetRoutes(app)
	if err != nil {
		log.Panicf("Error occured during routes creation: %s", err)
	}
	// Init graceful shutdown
	serverShutdown := Init_GracefulShutdown(app)

	// ...

	// Listen on PORT
	engine.Listen(app)
	// Graceful Cleanup
	serverShutdown.Wait()
	log.Print("Running cleanup tasks...")
	// Your cleanup tasks go here
	// ...
}
