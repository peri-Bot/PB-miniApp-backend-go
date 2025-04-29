package main

import (
	"fmt"
	"net/http"
	"os"

	// Use godotenv for local development to load .env file
	"github.com/joho/godotenv"
	"go.uber.org/zap"

	// Adjust import paths based on your go module name
	"github.com/peri-Bot/PB-miniApp-backend-go/internal/infrastructure/config"
	apperrors "github.com/peri-Bot/PB-miniApp-backend-go/internal/infrastructure/errors" // Alias to avoid collision
	"github.com/peri-Bot/PB-miniApp-backend-go/internal/infrastructure/logger"
)

func main() {
	// --- Early Setup: Load .env for local dev ---
	// Load .env file if it exists. Ignore error if it doesn't.
	// In production, rely solely on actual environment variables.
	_ = godotenv.Load() // Call this BEFORE LoadConfig

	// --- Load Configuration ---
	cfg, err := config.LoadConfig()
	if err != nil {
		// Use standard log before custom logger is initialized
		fmt.Fprintf(os.Stderr, "FATAL: Could not load configuration: %v\n", err)
		os.Exit(1)
	}

	// --- Initialize Logger ---
	appLogger, err := logger.NewLogger(cfg.AppEnv) // Use AppEnv from config
	if err != nil {
		fmt.Fprintf(os.Stderr, "FATAL: Could not initialize logger: %v\n", err)
		os.Exit(1)
	}
	// Defer logger sync to ensure logs are flushed before exit
	defer logger.SyncLogger(appLogger)

	// Replace standard logger (optional, use Zap directly)
	// zap.ReplaceGlobals(appLogger) // Use with caution

	appLogger.Info("Logger initialized", zap.String("environment", cfg.AppEnv))
	appLogger.Info("Configuration loaded successfully", zap.String("port", cfg.ServerPort))
	// Avoid logging sensitive info in production
	// appLogger.Debug("Mongo URI", zap.String("uri", cfg.MongoURI))
	// appLogger.Debug("Redis URI", zap.String("uri", cfg.RedisURI))

	// --- Dependency Injection and Server Setup ---
	// ... Initialize DB, Redis, Repositories, Use Cases, Handlers ...
	// Pass appLogger and cfg to components that need them.

	// Example: Using a custom error
	if cfg.ServerPort == ":0" { // Example check leading to an error
		err := apperrors.ErrValidation // Use the custom error
		appLogger.Error("Invalid server port configuration", zap.Error(err))
		// Handle the error appropriately (e.g., exit or return from function)
		os.Exit(1)
	}

	// Example placeholder server start
	appLogger.Info(fmt.Sprintf("Starting server on port %s", cfg.ServerPort))
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		appLogger.Info("Received request", zap.String("path", r.URL.Path)) // Log requests
		fmt.Fprintln(w, "Bingo Backend Running!")
	})

	err = http.ListenAndServe(cfg.ServerPort, mux)
	if err != nil {
		appLogger.Fatal("Server failed to start", zap.Error(err)) // Use Fatal for critical startup errors
	}
}
