package aps


// Ranges inclusive of the starting time and exclusive of the end time to make it easy to represent times that start/end on hour boundaries.
//
// For example, `start: '17:00'` and `end: '24:00'` will begin at 17:00 and
// finish immediately before 24:00.
type TimeRange struct {
	// The end time, specified in the format 'HH:MM' using 24 hour time.
	End *string `field:"required" json:"end" yaml:"end"`
	// The start time, specified in the format 'HH:MM' using 24 hour time.
	Start *string `field:"required" json:"start" yaml:"start"`
}

