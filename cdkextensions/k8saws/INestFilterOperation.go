package k8saws

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Represents an operation with excludive options that can be performed by the Fluent Bit Nest filter plugin.
type INestFilterOperation interface {
	Fields() *map[string]*[]*string
	Operation() NestFilterOperationType
}

// The jsii proxy for INestFilterOperation
type jsiiProxy_INestFilterOperation struct {
	_ byte // padding
}

func (j *jsiiProxy_INestFilterOperation) Fields() *map[string]*[]*string {
	var returns *map[string]*[]*string
	_jsii_.Get(
		j,
		"fields",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_INestFilterOperation) Operation() NestFilterOperationType {
	var returns NestFilterOperationType
	_jsii_.Get(
		j,
		"operation",
		&returns,
	)
	return returns
}

