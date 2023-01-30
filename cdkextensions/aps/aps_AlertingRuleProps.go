package aps

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Options needed to configure a Prometheus alerting rule inside an APS rules configuration.
type AlertingRuleProps struct {
	// The name of the alert.
	//
	// Must be a valid label value.
	Alert *string `field:"required" json:"alert" yaml:"alert"`
	// The PromQL expression to evaluate.
	//
	// Every evaluation cycle this is
	// evaluated at the current time, and all resultant time series become
	// pending/firing alerts.
	// See: [Querying prometheus](https://prometheus.io/docs/prometheus/latest/querying/basics/)
	//
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Annotations to add to each alert.
	//
	// Supports templating.
	Annotations *map[string]*string `field:"optional" json:"annotations" yaml:"annotations"`
	// Labels to add or overwrite for each alert.
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// Alerts are considered firing once they have been returned for this long.
	//
	// Alerts which have not yet fired for long enough are considered pending.
	Period awscdk.Duration `field:"optional" json:"period" yaml:"period"`
}

