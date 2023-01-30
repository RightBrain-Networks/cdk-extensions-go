package aps

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Configuration options that control the default behavior of alerts.
//
// These are the settings that will be used for any alerts that do not match a
// child node in the default routing tree or alerts that do match a child node
// that doesn't override the settings provded here.
//
// The default route does not support all the settings supported by other
// (child) routes.
type DefaultRouteOptions struct {
	// The labels by which incoming alerts are grouped together.
	//
	// For example,
	// multiple alerts coming in for cluster=A and alertname=LatencyHigh would be
	// batched into a single group.
	GroupByLabels *[]*string `field:"optional" json:"groupByLabels" yaml:"groupByLabels"`
	// How long to wait before sending a notification about new alerts that are added to a group of alerts for which an initial notification has already been sent (usually ~5m or more).
	GroupInterval awscdk.Duration `field:"optional" json:"groupInterval" yaml:"groupInterval"`
	// How long to initially wait to send a notification for a group of alerts.
	//
	// Allows to wait for an inhibiting alert to arrive or collect more initial
	// alerts for the same group (usually ~0s to few minutes).
	GroupWait awscdk.Duration `field:"optional" json:"groupWait" yaml:"groupWait"`
	// A list of matchers that an alert has to fulfill to match the node.
	Matchers *[]AlertManagerMatcher `field:"optional" json:"matchers" yaml:"matchers"`
	// How long to wait before sending a notification again if it has already been sent successfully for an alert.
	//
	// (Usually ~3h or more).
	RepeatInterval awscdk.Duration `field:"optional" json:"repeatInterval" yaml:"repeatInterval"`
}

