package config

import "os"

type Config struct {
	Port           string
	NodeServiceURL string
	AppName        string
	JWTToken       string
}

func LoadConfig() *Config {
	return &Config{
		Port:           getEnv("PORT", "8080"),
		NodeServiceURL: getEnv("NODE_SERVICE_URL", "http://node:3000"),
		AppName:        getEnv("APP_NAME", "Go QR Factorization API"),
		JWTToken:       getEnv("JWT_TOKEN", ""),
	}
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
