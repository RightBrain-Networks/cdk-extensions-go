package aps


// A range specifying the numerical days in the month.
//
// Days begin at 1. Negative values are also accepted which begin at the end of
// the month, e.g. -1 during January would represent January 31. For example:
// `start: 1` and `end: 5` or `start:-3` and `end: -1` would both be valid
// ranges.
//
// Extending past the start or end of the month will cause it to be clamped.
// E.g. specifying `start: 1` and `end: 31` during February will clamp the
// actual end date to 28 or 29 depending on leap years.
//
// Inclusive on both ends.
type DayOfMonthRange struct {
	// The last day of the month for which the range should apply (inclusive).
	End *float64 `field:"required" json:"end" yaml:"end"`
	// The first day of the month for which the range should apply.
	Start *float64 `field:"required" json:"start" yaml:"start"`
}

