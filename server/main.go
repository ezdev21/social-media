package main

import (
    "log"

    "server/config"
    "server/database"
    "server/routes"
    "github.com/joho/godotenv"
)

func main() {
    godotenv.Load()
    cfg, err := config.Load()
    if err != nil {
        log.Fatal(err)
    }

    db, err := database.Connect(cfg.DatabaseURL)
    if err != nil {
        log.Fatal("Database connection failed: ", err)
    }

    router := routes.Setup(db, cfg.JWTSecret)

    if err := router.Run(":8080"); err != nil {
        log.Fatal(err)
    }
}