package aps

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/constructs-go/constructs/v10"
)

// Represents a destination where alert manager can send notifications.
type IAlertManagerDestination interface {
	// Associates the destination with a construct that is handling the configuration of alert manager.
	//
	// Returns: An object representing the configured destination.
	Bind(scope constructs.IConstruct) *map[string]interface{}
	// The destination type being configured.
	//
	// Represents a config block in an alert manager receiver configuration.
	Category() AlertManagerDestinationCategory
}

// The jsii proxy for IAlertManagerDestination
type jsiiProxy_IAlertManagerDestination struct {
	_ byte // padding
}

func (i *jsiiProxy_IAlertManagerDestination) Bind(scope constructs.IConstruct) *map[string]interface{} {
	if err := i.validateBindParameters(scope); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		i,
		"bind",
		[]interface{}{scope},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_IAlertManagerDestination) Category() AlertManagerDestinationCategory {
	var returns AlertManagerDestinationCategory
	_jsii_.Get(
		j,
		"category",
		&returns,
	)
	return returns
}

