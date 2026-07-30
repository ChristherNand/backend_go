package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/christhernand/backendgo/internal/env"
)

func main() {
	// load .env file (if present) into the process environment
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file loaded:", err)
	}

	cfg := config{
		addr: env.GetString("ADDR", ":8080"),
	}
	app := &application{
		config: cfg,
	}
	mux := app.mount()

	log.Fatal(app.run(mux))
}
