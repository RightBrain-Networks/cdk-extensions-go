package k8saws

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Represents a Fluent Bit plugin that parses inbound records to populate fields.
type IFluentBitParserPlugin interface {
	IFluentBitPlugin
	Format() *string
}

// The jsii proxy for IFluentBitParserPlugin
type jsiiProxy_IFluentBitParserPlugin struct {
	jsiiProxy_IFluentBitPlugin
}

func (j *jsiiProxy_IFluentBitParserPlugin) Format() *string {
	var returns *string
	_jsii_.Get(
		j,
		"format",
		&returns,
	)
	return returns
}

