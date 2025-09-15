package main

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"GoATTHStart/internal/config"
	"GoATTHStart/internal/database"
	"GoATTHStart/internal/handlers"
	"GoATTHStart/internal/server"
	"GoATTHStart/internal/services"
)

func gracefulShutdown(apiServer *http.Server, done chan bool) {
	// Create context that listens for the interrupt signal from the OS.
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	// Listen for the interrupt signal.
	<-ctx.Done()

	log.Println("shutting down gracefully, press Ctrl+C again to force")
	stop() // Allow Ctrl+C to force shutdown

	// The context is used to inform the server it has 5 seconds to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := apiServer.Shutdown(ctx); err != nil {
		log.Printf("Server forced to shutdown with error: %v", err)
	}

	log.Println("Server exiting")

	// Notify the main goroutine that the shutdown is complete
	done <- true
}

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelInfo,
	}))

	cfg, err := config.Load(logger)
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	db, err := database.NewDBConnexion(&cfg.DBConfig, logger)
	if err != nil {
		logger.Error("failed to initialize database", "error", err)
		os.Exit(1)
	}

	// cache, err := cache.NewRedisClient(&cfg.CacheConfig)
	// if err != nil {
	// 	logger.Error("failed to initialize redis", "error", err)
	// 	os.Exit(1)
	// }

	healthService := services.NewHealthService(db)
	healthHandler := handlers.NewHealthHandler(healthService, logger)

	handlerStruct := &server.Handlers{Health: healthHandler}

	server := server.New(cfg, logger, handlerStruct)

	// Create a done channel to signal when the shutdown is complete
	done := make(chan bool, 1)

	// Run graceful shutdown in a separate goroutine
	go gracefulShutdown(server.GetHTTPServer(), done)

	err = server.Start(context.Background())
	if err != nil && err != http.ErrServerClosed {
		panic(fmt.Sprintf("http server error: %s", err))
	}

	// Wait for the graceful shutdown to complete
	<-done
	log.Println("Graceful shutdown complete.")
}
