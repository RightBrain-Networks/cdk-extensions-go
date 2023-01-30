package aps


// The days of the week to be used in Prometheus alert manager configurations.
type Weekday string

const (
	// Sunday.
	Weekday_SUNDAY Weekday = "SUNDAY"
	// Monday.
	Weekday_MONDAY Weekday = "MONDAY"
	// Tuesday.
	Weekday_TUESDAY Weekday = "TUESDAY"
	// Wednesday.
	Weekday_WEDNESDAY Weekday = "WEDNESDAY"
	// Thursday.
	Weekday_THURSDAY Weekday = "THURSDAY"
	// Friday.
	Weekday_FRIDAY Weekday = "FRIDAY"
	// Saturday.
	Weekday_SATURDAY Weekday = "SATURDAY"
)

