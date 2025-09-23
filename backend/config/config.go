package config

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	Server ServerConfig
	DB     DatabaseConfig
	JWT    JWTConfig
	CORS   CORSConfig
}

type ServerConfig struct {
	Port string
	Mode string
}

type DatabaseConfig struct {
	Type     string
	Host     string
	Port     string
	Name     string
	User     string
	Password string
}

type JWTConfig struct {
	Secret      string
	ExpireHours int
}

type CORSConfig struct {
	AllowOrigins string
}

func Load() *Config {
	// 加载环境变量文件
	_ = godotenv.Load()

	return &Config{
		Server: ServerConfig{
			Port: getEnv("SERVER_PORT", "8080"),
			Mode: getEnv("SERVER_MODE", "debug"),
		},
		DB: DatabaseConfig{
			Type:     getEnv("DB_TYPE", "sqlite"),
			Host:     getEnv("DB_HOST", "localhost"),
			Port:     getEnv("DB_PORT", "3306"),
			Name:     getEnv("DB_NAME", "progress_wall"),
			User:     getEnv("DB_USER", "root"),
			Password: getEnv("DB_PASSWORD", ""),
		},
		JWT: JWTConfig{
			Secret:      getEnv("JWT_SECRET", "default-secret-key"),
			ExpireHours: getEnvAsInt("JWT_EXPIRE_HOURS", 24),
		},
		CORS: CORSConfig{
			AllowOrigins: getEnv("CORS_ALLOW_ORIGINS", "http://localhost:3000,http://localhost:5173"),
		},
	}
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func getEnvAsInt(key string, defaultValue int) int {
	if value := os.Getenv(key); value != "" {
		if intValue, err := strconv.Atoi(value); err == nil {
			return intValue
		}
	}
	return defaultValue
}
