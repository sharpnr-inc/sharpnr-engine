package main

import (
	"log"

	"github.com/sharpnr-inc/sharpnr-engine/config"
)

// main is the entry point of the application.
// It prints a greeting message to the console.
func main() {
	config.LoadEnv()
	log.Println("App started...")
}
