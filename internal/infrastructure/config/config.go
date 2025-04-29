package config

import (
	"fmt"
	"os"
	"strconv" // For parsing port if needed, though string is often fine
	"strings"
	// "time" // Import if you add time.Duration fields later
)

// Config holds all configuration for the application.
// Values are loaded directly from environment variables.
type Config struct {
	// Database
	MongoURI string
	// MongoDatabaseName string // Example: Add if needed
	// MongoTimeout      time.Duration // Example: Add if needed

	// Redis
	RedisURI string
	// RedisTimeout time.Duration // Example: Add if needed

	// Authentication & Security
	JWTSecret        string
	TelegramBotToken string
	// JWTExpiration    time.Duration // Example: Add if needed

	// Server
	ServerPort string
	// ReadTimeout  time.Duration // Example: Add if needed
	// WriteTimeout time.Duration // Example: Add if needed
	// IdleTimeout  time.Duration // Example: Add if needed

	// Application Environment (useful for logger setup, etc.)
	AppEnv string // e.g., "development", "production"
}

// LoadConfig reads configuration directly from environment variables.
// It assumes godotenv.Load() might have been called beforehand in main.go
// to load a .env file for local development.
func LoadConfig() (*Config, error) {
	cfg := &Config{}

	// Load values using os.Getenv
	cfg.MongoURI = os.Getenv("MONGODB_URI")
	cfg.RedisURI = os.Getenv("REDIS_URI")
	cfg.JWTSecret = os.Getenv("JWT_SECRET")
	cfg.TelegramBotToken = os.Getenv("BOT_TOKEN")
	cfg.ServerPort = os.Getenv("SERVER_PORT")
	cfg.AppEnv = os.Getenv("APP_ENV") // Environment variable for app mode

	// --- Set Defaults ---
	if cfg.ServerPort == "" {
		cfg.ServerPort = "8080" // Default port
	}
	if cfg.AppEnv == "" {
		cfg.AppEnv = "development" // Default environment
	}

	// --- Validation (Essential) ---
	if cfg.MongoURI == "" {
		return nil, fmt.Errorf("config: MONGODB_URI environment variable is required")
	}
	if cfg.RedisURI == "" {
		return nil, fmt.Errorf("config: REDIS_URI environment variable is required")
	}
	if cfg.JWTSecret == "" {
		return nil, fmt.Errorf("config: JWT_SECRET environment variable is required")
	}
	if cfg.TelegramBotToken == "" {
		return nil, fmt.Errorf("config: BOT_TOKEN environment variable is required")
	}

	// --- Post-processing ---
	// Ensure ServerPort starts with a colon if it's just a number
	if port := cfg.ServerPort; port != "" && !strings.HasPrefix(port, ":") {
		// Optional: Validate if it's a valid port number
		if _, err := strconv.Atoi(port); err == nil {
			cfg.ServerPort = ":" + port
		} else {
			// Handle case where ServerPort might be a host:port string already
			// Or return an error if format is unexpected
			// For simplicity here, we assume it's either ":port" or "port"
		}
	}

	return cfg, nil
}
