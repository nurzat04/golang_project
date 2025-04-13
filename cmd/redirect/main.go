package main

import (
	"log"
	"shorty/internal/redirect"
)

func main() {
	if err := redirect.Run(); err != nil {
		log.Fatalf("failed to start redirect service: %v", err)
	}
}
