package aps


// A range of calendar months identified by number, where January = 1.
//
// Ranges are also accepted by specifying `end` inclusive on both ends.
type MonthRange struct {
	// The month where the range should start.
	Start *float64 `field:"required" json:"start" yaml:"start"`
	// The month at the end of the range if the range should cover multiple months (inclusive).
	End *float64 `field:"optional" json:"end" yaml:"end"`
}

