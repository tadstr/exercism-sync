package booking

import (
	"fmt"
	"time"
)

// Schedule returns a time.Time from a string containing a date.
func Schedule(date string) time.Time {
	layout := "1/2/2006 15:04:05"
	t, err := time.Parse(layout, date)
	if err != nil {
		fmt.Printf("Error parsing date: %v\n", err)
	}
	return t
}

// HasPassed returns whether a date has passed.
func HasPassed(date string) bool {
    layout := "January _2, 2006 15:04:05"
    t, err := time.Parse(layout, date)
    
    now := time.Now()

    if err != nil {
		fmt.Printf("Error parsing date: %v\n", err)
	}

	return t.Before(now)
}

// IsAfternoonAppointment returns whether a time is in the afternoon.
func IsAfternoonAppointment(date string) bool {
	layout := "Monday, January 2, 2006 15:04:05"
	t, err := time.Parse(layout, date)
	if err != nil {
		return false
	}
	hour := t.Hour()
	return hour >= 12 && hour < 18
}

// Description returns a formatted string of the appointment time.
func Description(date string) string {
	t, err := time.Parse("1/2/2006 15:04:05", date)
	if err != nil {
		fmt.Printf("Error parsing date: %v\n", err)
	}
	formattedDate := t.Format("Monday, January 2, 2006, at 15:04.")
	return fmt.Sprintf("You have an appointment on %s", formattedDate)
}

// AnniversaryDate returns a Time with this year's anniversary.
func AnniversaryDate() time.Time {
	now := time.Now()
    year := now.Year()
    t := time.Date(year, time.September, 15, 0, 0, 0, 0,time.UTC)
    return t
}
