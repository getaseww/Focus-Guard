package main

import (
	"focus-guard/db"
	"focus-guard/proxy"
	"focus-guard/schedule"
	"log"
)

func main() {
	// Initialize database and schedule management
	db.SetupDatabase("focus-guard.db")
	schedule.StartScheduleChecker()

	// Start the proxy server
	err := proxy.StartProxy()
	if err != nil {
		log.Fatalf("Failed to start proxy: %v", err)
	}
}
