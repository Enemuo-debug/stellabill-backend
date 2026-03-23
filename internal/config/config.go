package config

import "os"

type Config struct {
	Env       string
	Port      string
	DBConn    string
	JWTSecret string
	// AllowedOrigins is a comma-separated list of permitted CORS origins used in
	// production and staging. Example: "https://app.stellarbill.com,https://admin.stellarbill.com"
	AllowedOrigins string
}

func Load() Config {
	return Config{
		Env:            getEnv("ENV", "development"),
		Port:           getEnv("PORT", "8080"),
		DBConn:         getEnv("DATABASE_URL", "postgres://localhost/stellarbill?sslmode=disable"),
		JWTSecret:      getEnv("JWT_SECRET", "change-me-in-production"),
		AllowedOrigins: getEnv("ALLOWED_ORIGINS", ""),
	}
}

func getEnv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}
