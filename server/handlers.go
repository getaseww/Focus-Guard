package main

import (
	"encoding/json"
	"focus-guard/db"
	"focus-guard/schedule"
	"net/http"
)

// func GetSchedules(w http.ResponseWriter) {
// 	rows, err := db.DB.Query("SELECT id, url, start_time, end_time, day_of_week FROM schedules")
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}
// 	defer rows.Close()

// 	var schedules []schedule.Schedule
// 	for rows.Next() {
// 		var s schedule.Schedule
// 		if err := rows.Scan(&s.ID, &s.URL, &s.StartTime, &s.EndTime, &s.DayOfWeek); err != nil {
// 			http.Error(w, err.Error(), http.StatusInternalServerError)
// 			return
// 		}
// 		schedules = append(schedules, s)
// 	}
// 	w.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(schedules)
// }

// func CreateSchedule(w http.ResponseWriter, r *http.Request) {
// 	var s schedule.Schedule
// 	if err := json.NewDecoder(r.Body).Decode(&s); err != nil {
// 		http.Error(w, err.Error(), http.StatusBadRequest)
// 		return
// 	}

// 	result, err := db.DB.Exec(
// 		"INSERT INTO schedules (url, start_time, end_time, day_of_week) VALUES (?, ?, ?, ?)",
// 		s.URL, s.StartTime, s.EndTime, s.DayOfWeek,
// 	)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}
// 	id, _ := result.LastInsertId()
// 	s.ID = int(id) // Convert int64 to int
// 	w.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(s)
// }

func DeleteSchedule(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "Missing id parameter", http.StatusBadRequest)
		return
	}

	_, err := db.DB.Exec("DELETE FROM schedules WHERE id = ?", id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func UpdateSchedule(w http.ResponseWriter, r *http.Request) {
	var s schedule.Schedule
	if err := json.NewDecoder(r.Body).Decode(&s); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	_, err := db.DB.Exec(
		"UPDATE schedules SET url = ?, start_time = ?, end_time = ?, day_of_week = ? WHERE id = ?",
		s.URL, s.StartTime, s.EndTime, s.DayOfWeek, s.ID,
	)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(s)
}
