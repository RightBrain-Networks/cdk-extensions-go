package k8saws

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/vibe-io/cdk-extensions-go/cdkextensions/jsii"
)

// Operations with exclusive options that can be performed by the Fluent Bit Nest filter plugin.
type NestFilterOperation interface {
	INestFilterOperation
	// The fields representing configuration options for the operation.
	Fields() *map[string]*[]*string
	// The type of operation to be performed.
	Operation() NestFilterOperationType
}

// The jsii proxy struct for NestFilterOperation
type jsiiProxy_NestFilterOperation struct {
	jsiiProxy_INestFilterOperation
}

func (j *jsiiProxy_NestFilterOperation) Fields() *map[string]*[]*string {
	var returns *map[string]*[]*string
	_jsii_.Get(
		j,
		"fields",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NestFilterOperation) Operation() NestFilterOperationType {
	var returns NestFilterOperationType
	_jsii_.Get(
		j,
		"operation",
		&returns,
	)
	return returns
}


func NestFilterOperation_Lift(options *LiftOptions) INestFilterOperation {
	_init_.Initialize()

	if err := validateNestFilterOperation_LiftParameters(options); err != nil {
		panic(err)
	}
	var returns INestFilterOperation

	_jsii_.StaticInvoke(
		"cdk-extensions.k8s_aws.NestFilterOperation",
		"lift",
		[]interface{}{options},
		&returns,
	)

	return returns
}

func NestFilterOperation_Nest(options *NestOptions) INestFilterOperation {
	_init_.Initialize()

	if err := validateNestFilterOperation_NestParameters(options); err != nil {
		panic(err)
	}
	var returns INestFilterOperation

	_jsii_.StaticInvoke(
		"cdk-extensions.k8s_aws.NestFilterOperation",
		"nest",
		[]interface{}{options},
		&returns,
	)

	return returns
}

