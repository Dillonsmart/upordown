package main

import (
	"log"
	"os"
	"upordown/internal/storage"

	"github.com/joho/godotenv"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Printf("Warning: Could not load .env file: %v", err)
	}

	if len(os.Args) < 2 {
		os.Exit(1)
	}

	cmd := os.Args[1]

	storage.Init()

	switch cmd {
	case "track":
		HandleTrack(os.Args[2:])
	case "untrack":
		// Stop tracking the website URL
	}
}
