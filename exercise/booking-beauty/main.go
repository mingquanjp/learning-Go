package main

import (
	"fmt"
	"time"
)

func main() {
	date := "7/2/2026 15:04:03"
	fmt.Println(Schedule(date))
	fmt.Println(Description(date))
	fmt.Println("Anniversary:", AnniversaryDate())
}

// Schedule returns a time.Time from a string containing a date.
func Schedule(date string) time.Time {
	t, err := time.Parse("1/2/2006 15:04:03", date)
	if err != nil {
		panic(err)
	}
	return t
}

// HasPassed returns whether a date has passed.
func HasPassed(date string) bool {
	t, err := time.Parse("July 2, 2006 15:04:03", date)
	if err != nil {
		panic(err)
	}
	if t.Before(time.Now()) {
		return true
	}
	return false
}

// IsAfternoonAppointment returns whether a time is in the afternoon.
func IsAfternoonAppointment(date string) bool {
	t, err := time.Parse("Monday, January 2, 2006 15:04:05", date)
	if err != nil {
		panic(err)
	}
	return t.Hour() >= 12 && t.Hour() < 18
}

// Description returns a formatted string of the appointment time.
func Description(date string) string {
	t, err := time.Parse("1/2/2006 15:04:05", date)
	if err != nil {
		panic(err)
	}

	return fmt.Sprintf("You have an appointment on %s", t.Format("Monday, January 2, 2006, at 15:04"))
}

// AnniversaryDate returns a Time with this year's anniversary.
func AnniversaryDate() time.Time {
	timeNow := time.Now()
	return time.Date(timeNow.Year(), timeNow.Month(), timeNow.Day(), 0, 0, 0, 0, timeNow.Location())
}
