package aps


// A numerical range of years.
//
// Ranges to cover multiple years are accepted. For example, `start: 2020` and
// `end: 2022`.
//
// Inclusive on both ends.
type YearRange struct {
	// The year where the range should start.
	Start *float64 `field:"required" json:"start" yaml:"start"`
	// The year where the range should end.
	//
	// If not specified, the range will end at the end of the year specified by
	// `start`.
	End *float64 `field:"optional" json:"end" yaml:"end"`
}

