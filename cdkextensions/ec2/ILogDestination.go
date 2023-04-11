package ec2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/constructs-go/constructs/v10"
)

// Represents a resource that can act as a deliver endpoint for captured flow logs.
type ILogDestination interface {
	Bind(scope constructs.IConstruct) *FlowLogDestinationConfig
}

// The jsii proxy for ILogDestination
type jsiiProxy_ILogDestination struct {
	_ byte // padding
}

func (i *jsiiProxy_ILogDestination) Bind(scope constructs.IConstruct) *FlowLogDestinationConfig {
	if err := i.validateBindParameters(scope); err != nil {
		panic(err)
	}
	var returns *FlowLogDestinationConfig

	_jsii_.Invoke(
		i,
		"bind",
		[]interface{}{scope},
		&returns,
	)

	return returns
}

