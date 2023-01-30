package aps

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/vibe-io/cdk-extensions-go/cdkextensions/jsii"
)

// Specifies the type of destination where alert manager can send notifications.
//
// Corresponds with a config block inside an alert manager receiver
// configuration.
type AlertManagerDestinationCategory interface {
	// The name of the category, as it would appear as a key in the configuration for an alert manager receiver.
	Name() *string
}

// The jsii proxy struct for AlertManagerDestinationCategory
type jsiiProxy_AlertManagerDestinationCategory struct {
	_ byte // padding
}

func (j *jsiiProxy_AlertManagerDestinationCategory) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}


// An escape hatch method that allows specifying arbitrary values for the type a receiver field a destination should be placed under.
//
// In the event that new destination types are added by alert manager, this
// can be used to implement custom destinations in the event that it hasn't
// had the chance to be implemented here.
//
// Whenever possible, it is recommended the static values provided be used.
//
// Returns: An object that can be used when building a receiver that
// specifies the given destination.
func AlertManagerDestinationCategory_Of(name *string) AlertManagerDestinationCategory {
	_init_.Initialize()

	if err := validateAlertManagerDestinationCategory_OfParameters(name); err != nil {
		panic(err)
	}
	var returns AlertManagerDestinationCategory

	_jsii_.StaticInvoke(
		"cdk-extensions.aps.AlertManagerDestinationCategory",
		"of",
		[]interface{}{name},
		&returns,
	)

	return returns
}

func AlertManagerDestinationCategory_EMAIL() AlertManagerDestinationCategory {
	_init_.Initialize()
	var returns AlertManagerDestinationCategory
	_jsii_.StaticGet(
		"cdk-extensions.aps.AlertManagerDestinationCategory",
		"EMAIL",
		&returns,
	)
	return returns
}

func AlertManagerDestinationCategory_OPSGENIE() AlertManagerDestinationCategory {
	_init_.Initialize()
	var returns AlertManagerDestinationCategory
	_jsii_.StaticGet(
		"cdk-extensions.aps.AlertManagerDestinationCategory",
		"OPSGENIE",
		&returns,
	)
	return returns
}

func AlertManagerDestinationCategory_PAGERDUTY() AlertManagerDestinationCategory {
	_init_.Initialize()
	var returns AlertManagerDestinationCategory
	_jsii_.StaticGet(
		"cdk-extensions.aps.AlertManagerDestinationCategory",
		"PAGERDUTY",
		&returns,
	)
	return returns
}

func AlertManagerDestinationCategory_PUSHOVER() AlertManagerDestinationCategory {
	_init_.Initialize()
	var returns AlertManagerDestinationCategory
	_jsii_.StaticGet(
		"cdk-extensions.aps.AlertManagerDestinationCategory",
		"PUSHOVER",
		&returns,
	)
	return returns
}

func AlertManagerDestinationCategory_SLACK() AlertManagerDestinationCategory {
	_init_.Initialize()
	var returns AlertManagerDestinationCategory
	_jsii_.StaticGet(
		"cdk-extensions.aps.AlertManagerDestinationCategory",
		"SLACK",
		&returns,
	)
	return returns
}

func AlertManagerDestinationCategory_SNS() AlertManagerDestinationCategory {
	_init_.Initialize()
	var returns AlertManagerDestinationCategory
	_jsii_.StaticGet(
		"cdk-extensions.aps.AlertManagerDestinationCategory",
		"SNS",
		&returns,
	)
	return returns
}

func AlertManagerDestinationCategory_TELEGRAM() AlertManagerDestinationCategory {
	_init_.Initialize()
	var returns AlertManagerDestinationCategory
	_jsii_.StaticGet(
		"cdk-extensions.aps.AlertManagerDestinationCategory",
		"TELEGRAM",
		&returns,
	)
	return returns
}

func AlertManagerDestinationCategory_VICTOROPS() AlertManagerDestinationCategory {
	_init_.Initialize()
	var returns AlertManagerDestinationCategory
	_jsii_.StaticGet(
		"cdk-extensions.aps.AlertManagerDestinationCategory",
		"VICTOROPS",
		&returns,
	)
	return returns
}

func AlertManagerDestinationCategory_WEBEX() AlertManagerDestinationCategory {
	_init_.Initialize()
	var returns AlertManagerDestinationCategory
	_jsii_.StaticGet(
		"cdk-extensions.aps.AlertManagerDestinationCategory",
		"WEBEX",
		&returns,
	)
	return returns
}

func AlertManagerDestinationCategory_WEBHOOK() AlertManagerDestinationCategory {
	_init_.Initialize()
	var returns AlertManagerDestinationCategory
	_jsii_.StaticGet(
		"cdk-extensions.aps.AlertManagerDestinationCategory",
		"WEBHOOK",
		&returns,
	)
	return returns
}

func AlertManagerDestinationCategory_WECHAT() AlertManagerDestinationCategory {
	_init_.Initialize()
	var returns AlertManagerDestinationCategory
	_jsii_.StaticGet(
		"cdk-extensions.aps.AlertManagerDestinationCategory",
		"WECHAT",
		&returns,
	)
	return returns
}

