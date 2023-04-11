package aps

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/constructs-go/constructs/v10"
)

// Represents an alert manager configuration object that can be used by resources configuring alerting for Amazon APS.
type IAlertManagerConfiguration interface {
	// Associates the configuration with a construct that is handling the configuration of alert manager for an APS workspace.
	//
	// Returns: Alert manager configuration details.
	Bind(scope constructs.IConstruct) *AlertManagerConfigurationDetails
}

// The jsii proxy for IAlertManagerConfiguration
type jsiiProxy_IAlertManagerConfiguration struct {
	_ byte // padding
}

func (i *jsiiProxy_IAlertManagerConfiguration) Bind(scope constructs.IConstruct) *AlertManagerConfigurationDetails {
	if err := i.validateBindParameters(scope); err != nil {
		panic(err)
	}
	var returns *AlertManagerConfigurationDetails

	_jsii_.Invoke(
		i,
		"bind",
		[]interface{}{scope},
		&returns,
	)

	return returns
}

