package logger

import (
	"fmt"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// NewLogger creates a new Zap logger instance based on the application environment.
// Pass "production" for JSON structured logging, otherwise assumes development mode
// with human-readable console output.
func NewLogger(appEnv string) (*zap.Logger, error) {
	var cfg zap.Config
	var err error

	// Configure logger based on environment
	if appEnv == "production" {
		cfg = zap.NewProductionConfig()
		// Optional: Customize production logging further
		// cfg.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder // Example
		// cfg.OutputPaths = []string{"stdout", "/var/log/bingo-backend.log"} // Example: Log to file too
		// cfg.ErrorOutputPaths = []string{"stderr", "/var/log/bingo-backend-error.log"} // Example
	} else {
		// Development configuration (more human-readable)
		cfg = zap.NewDevelopmentConfig()
		// Use console encoder with colors for development
		cfg.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	}

	// Optional: Set log level from environment variable if needed
	logLevel := os.Getenv("LOG_LEVEL")
	if logLevel != "" {
		var level zapcore.Level
		if err := level.Set(logLevel); err == nil {
			cfg.Level = zap.NewAtomicLevelAt(level)
		} else {
			// Log a warning if the level is invalid, but proceed with default
			fmt.Fprintf(os.Stderr, "Warning: Invalid LOG_LEVEL '%s'. Using default: %s\n", logLevel, cfg.Level.Level().String())
		}
	} else {
		// Default level if not set (usually Info for Production, Debug for Development)
		// The NewProductionConfig/NewDevelopmentConfig already set reasonable defaults.
	}

	// Build the logger
	logger, err := cfg.Build()
	if err != nil {
		return nil, fmt.Errorf("logger: failed to build zap logger: %w", err)
	}

	// Redirect Go's standard library logs to Zap (optional but recommended)
	// zap.RedirectStdLog(logger) // Be careful with this if other libraries also try to redirect

	return logger, nil
}

// SyncLogger flushes any buffered log entries.
// It's important to call this before application exit.
func SyncLogger(logger *zap.Logger) {
	if logger != nil {
		// Ignore error on sync as per Zap documentation recommendation in most cases
		_ = logger.Sync()
	}
}
