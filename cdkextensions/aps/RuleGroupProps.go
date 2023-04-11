package aps

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Options needed to configure a Prometheus rule group for use with an APS rule groups namespace configuration.
type RuleGroupProps struct {
	// How often rules in the group are evaluated.
	Interval awscdk.Duration `field:"optional" json:"interval" yaml:"interval"`
	// Limit the number of alerts an alerting rule and series a recording rule can produce.
	//
	// 0 is no limit.
	Limit *float64 `field:"optional" json:"limit" yaml:"limit"`
	// The name of the group.
	//
	// Must be unique within the configuration.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The rules to be evaluated per the rule group's configuration.
	Rules *[]IPrometheusRule `field:"optional" json:"rules" yaml:"rules"`
}

