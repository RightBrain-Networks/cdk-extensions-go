package k8saws

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Options for configuring the Throttle Fluent Bit filter plugin.
// See: [Throttle Plugin Documention](https://docs.fluentbit.io/manual/pipeline/filters/throttle)
//
type FluentBitThrottleFilterOptions struct {
	// The pattern to match for records that this output should apply to.
	Match FluentBitMatch `field:"optional" json:"match" yaml:"match"`
	// Time interval.
	Interval awscdk.Duration `field:"optional" json:"interval" yaml:"interval"`
	// Whether to print status messages with current rate and the limits to information logs.
	PrintStatus *bool `field:"optional" json:"printStatus" yaml:"printStatus"`
	// Amount of messages for the time.
	Rate *float64 `field:"optional" json:"rate" yaml:"rate"`
	// Amount of intervals to calculate average over.
	// Default: 5.
	//
	Window *float64 `field:"optional" json:"window" yaml:"window"`
}

