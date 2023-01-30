package aps

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/vibe-io/cdk-extensions-go/cdkextensions/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2/awssns"
)

// Provides an interface for creating various alert manager destination objects that can receive notifications when an alert happens in Prometheus.
type AlertManagerDestination interface {
}

// The jsii proxy struct for AlertManagerDestination
type jsiiProxy_AlertManagerDestination struct {
	_ byte // padding
}

func NewAlertManagerDestination() AlertManagerDestination {
	_init_.Initialize()

	j := jsiiProxy_AlertManagerDestination{}

	_jsii_.Create(
		"cdk-extensions.aps.AlertManagerDestination",
		nil, // no parameters
		&j,
	)

	return &j
}

func NewAlertManagerDestination_Override(a AlertManagerDestination) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk-extensions.aps.AlertManagerDestination",
		nil, // no parameters
		a,
	)
}

// Creates an alert manager destination that sends alert notifications to an Amazon SNS topic.
//
// Returns: An alert manager destination that represents the SNS topic.
func AlertManagerDestination_SnsTopic(topic awssns.ITopic, options *AlertManagerSnsDestinationOptions) AlertManagerSnsDestination {
	_init_.Initialize()

	if err := validateAlertManagerDestination_SnsTopicParameters(topic, options); err != nil {
		panic(err)
	}
	var returns AlertManagerSnsDestination

	_jsii_.StaticInvoke(
		"cdk-extensions.aps.AlertManagerDestination",
		"snsTopic",
		[]interface{}{topic, options},
		&returns,
	)

	return returns
}

