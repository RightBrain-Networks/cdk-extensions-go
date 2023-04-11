package aps


// A day of the week, where the week begins on Sunday and ends on Saturday.
//
// For convenience, ranges are also accepted by specifying `end` and are
// inclusive on both ends.
type WeekdayRange struct {
	// The day of the week where the range should start.
	Start Weekday `field:"required" json:"start" yaml:"start"`
	// The day of the week where the range should end.
	//
	// If not specified, the range will end at the end of the day specified by
	// `start`.
	End Weekday `field:"optional" json:"end" yaml:"end"`
}

