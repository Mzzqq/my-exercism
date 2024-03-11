package main

import "time"

// Schedule returns a time.Time from a string containing a date.
// "7/25/2019 13:45:00"
func Schedule(date string) time.Time {
	t, err := time.Parse("1/02/2006 15:04:05", date)
	if err != nil {
		panic(err)
	}
	return t
}
