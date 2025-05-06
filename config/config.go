package config

import (
	"log"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

type Config struct {
	Port          string
	MongoURI      string
	RedisURI      string
	JWTSecret     string
	BotToken      string // Telegram Bot Token
	FrontendURL   string
	JWTExpire     time.Duration
	RedisKeyTTL   time.Duration // General TTL for keys like game state, room players
	SessionKeyTTL time.Duration // TTL for user sessions
	CacheKeyTTL   time.Duration // TTL for general caching (like rooms list)
}

var AppConfig *Config

func LoadConfig() (*Config, error) {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, relying on environment variables")
	}

	port := getEnv("PORT", "5000")
	mongoURI := getEnv("MONGODB_URI", "")
	redisURI := getEnv("REDIS_URI", "")
	jwtSecret := getEnv("JWT_SECRET", "default_super_secret_key") // Use a strong default ONLY for dev
	botToken := getEnv("BOT_TOKEN", "")
	frontendURL := getEnv("FRONTEND_URL", "http://localhost:5173") // Default to Vite dev server

	if mongoURI == "" {
		log.Fatal("MONGODB_URI environment variable is required")
	}
	if redisURI == "" {
		log.Fatal("REDIS_URI environment variable is required")
	}
	if botToken == "" {
		log.Fatal("BOT_TOKEN environment variable is required")
	}

	jwtExpireMinutes, _ := strconv.Atoi(getEnv("JWT_EXPIRE_MINUTES", "60"))             // e.g., 1 hour
	redisKeyTTLSeconds, _ := strconv.Atoi(getEnv("REDIS_KEY_TTL_SECONDS", "7200"))      // 2 hours
	sessionKeyTTLSeconds, _ := strconv.Atoi(getEnv("SESSION_KEY_TTL_SECONDS", "86400")) // 24 hours
	cacheKeyTTLSeconds, _ := strconv.Atoi(getEnv("CACHE_KEY_TTL_SECONDS", "600"))       // 10 minutes

	AppConfig = &Config{
		Port:          port,
		MongoURI:      mongoURI,
		RedisURI:      redisURI,
		JWTSecret:     jwtSecret,
		BotToken:      botToken,
		FrontendURL:   frontendURL,
		JWTExpire:     time.Duration(jwtExpireMinutes) * time.Minute,
		RedisKeyTTL:   time.Duration(redisKeyTTLSeconds) * time.Second,
		SessionKeyTTL: time.Duration(sessionKeyTTLSeconds) * time.Second,
		CacheKeyTTL:   time.Duration(cacheKeyTTLSeconds) * time.Second,
	}

	log.Printf("Configuration loaded. Port: %s", AppConfig.Port)
	return AppConfig, nil
}

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}
