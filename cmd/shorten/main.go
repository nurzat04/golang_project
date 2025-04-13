package main

import (
	"log"
	"shorty/internal/shorten"
)

func main() {
	if err := shorten.Run(); err != nil {
		log.Fatalf("failed to start shorten service: %v", err)
	}
}
