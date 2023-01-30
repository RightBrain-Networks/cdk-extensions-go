package aps


// Configuration for the alert manager time interval.
type TimeIntervalProps struct {
	// The interval definitions that define the periods of time that the time interval should apply for.
	Intervals *[]TimeIntervalEntry `field:"optional" json:"intervals" yaml:"intervals"`
	// The name of the time interval as it will be referenced throught the rest of the alert manager configuration.
	Name *string `field:"optional" json:"name" yaml:"name"`
}

