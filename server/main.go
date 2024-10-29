package server

import (
	"focus-guard/db"
	"log"
	"net/http"
)

func main() {
	// Initialize the database connection
	db.SetupDatabase("focus_guard.db")

	http.HandleFunc("/api/schedules", SchedulesHandler)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("server/static"))))

	log.Println("Starting server on :8081")
	if err := http.ListenAndServe(":8081", nil); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
