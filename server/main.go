package main

import (
	"encoding/json"
	"focus-guard/db"
	"focus-guard/schedule"
	"log"
	"net/http"
	"time"
)

func CORSMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Set CORS headers
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000") // Replace with your frontend URL
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		w.Header().Set("Access-Control-Allow-Credentials", "true")

		// Handle preflight (OPTIONS) requests
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func SchedulesHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		GetSchedules(w)
	case "POST":
		CreateSchedule(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func GetSchedules(w http.ResponseWriter) {
	rows, err := db.DB.Query("SELECT id, url, start_time, end_time, day_of_week FROM schedules")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var schedules []schedule.Schedule
	for rows.Next() {
		var s schedule.Schedule
		var tempStartTime, tempEndTime string
		if err := rows.Scan(&s.ID, &s.URL, &tempStartTime, &tempEndTime, &s.DayOfWeek); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		s.StartTime, err = time.Parse("2006-01-02 15:04:05", tempStartTime)
		s.EndTime, err = time.Parse("2006-01-02 15:04:05", tempEndTime)
		schedules = append(schedules, s)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(schedules)
}

func CreateSchedule(w http.ResponseWriter, r *http.Request) {
	var s schedule.Schedule
	if err := json.NewDecoder(r.Body).Decode(&s); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	result, err := db.DB.Exec(
		"INSERT INTO schedules (url, start_time, end_time, day_of_week) VALUES (?, ?, ?, ?)",
		s.URL, s.StartTime, s.EndTime, s.DayOfWeek,
	)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	id, _ := result.LastInsertId()
	s.ID = int(id)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(s)
}

func main() {
	// Initialize the database connection
	db.SetupDatabase("focus_guard.db")

	fs := http.FileServer(http.Dir("./frontend/build"))
	http.Handle("/", fs)
	// Wrap the SchedulesHandler with CORSMiddleware
	http.Handle("/api/schedules", CORSMiddleware(http.HandlerFunc(SchedulesHandler)))

	log.Println("Starting server on :8081")
	if err := http.ListenAndServe(":8081", nil); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
