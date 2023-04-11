package glue

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsglue"
	"github.com/aws/constructs-go/constructs/v10"
)

// Represents a precondition that must be satisfied in order for a trigger to be executed.
type ITriggerPredicate interface {
	Bind(scope constructs.IConstruct) *awsglue.CfnTrigger_ConditionProperty
}

// The jsii proxy for ITriggerPredicate
type jsiiProxy_ITriggerPredicate struct {
	_ byte // padding
}

func (i *jsiiProxy_ITriggerPredicate) Bind(scope constructs.IConstruct) *awsglue.CfnTrigger_ConditionProperty {
	if err := i.validateBindParameters(scope); err != nil {
		panic(err)
	}
	var returns *awsglue.CfnTrigger_ConditionProperty

	_jsii_.Invoke(
		i,
		"bind",
		[]interface{}{scope},
		&returns,
	)

	return returns
}

