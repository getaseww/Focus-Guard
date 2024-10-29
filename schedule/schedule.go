package schedule

import (
	"focus-guard/db"
	"log"
	"time"
)

// Schedule represents a blocking schedule
type Schedule struct {
	ID        int
	URL       string
	StartTime time.Time
	EndTime   time.Time
	DayOfWeek int
}

// StartScheduleChecker periodically updates block rules
func StartScheduleChecker() {
	go func() {
		for {
			time.Sleep(1 * time.Minute)
			updateBlockRules()
		}
	}()
}

// Check if a URL is blocked based on the current time
func IsBlocked(url string) bool {
	// Retrieve current time
	now := time.Now()

	// Query the DB for schedules matching the URL and current time
	rows, err := db.DB.Query(`
        SELECT url FROM schedules WHERE url = ? AND day_of_week = ? 
        AND start_time <= ? AND end_time >= ?;
    `, url, int(now.Weekday()), now.Format("15:04"), now.Format("15:04"))

	if err != nil {
		log.Println("Failed to query schedules:", err)
		return false
	}
	defer rows.Close()

	return rows.Next() // true if there's a matching row
}

// Update block rules (future expansion for in-memory rule updates)
func updateBlockRules() {
	// Example function to update in-memory block list
}
