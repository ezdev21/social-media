package config

import (
    "fmt"
    "os"
)

type Config struct {
    DatabaseURL string
    JWTSecret   string
}

func Load() (*Config, error) {
    cfg := &Config{
        DatabaseURL: os.Getenv("DATABASE_URL"),
        JWTSecret:   os.Getenv("JWT_SECRET"),
    }

    if cfg.DatabaseURL == "" {
        return nil, fmt.Errorf("DATABASE_URL is required")
    }

    if len(cfg.JWTSecret) < 32 {
        return nil, fmt.Errorf("JWT_SECRET must be at least 32 bytes")
    }

    return cfg, nil
}