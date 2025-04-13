package main

import (
	"log"
	"shorty/internal/stats"
)

func main() {
	if err := stats.Run(); err != nil {
		log.Fatalf("failed to start stats service: %v", err)
	}
}
