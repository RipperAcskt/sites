package main

import (
	"log"
	"sites/internal/app"
)

func main() {
	// generate.Gen()
	if err := app.Run(); err != nil {
		log.Fatalf("app run failed: %v", err)
	}
}
