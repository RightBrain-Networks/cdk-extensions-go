package glue

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsglue"
)

type ITriggerPredicate interface {
	Bind(trigger Trigger) *awsglue.CfnTrigger_ConditionProperty
}

// The jsii proxy for ITriggerPredicate
type jsiiProxy_ITriggerPredicate struct {
	_ byte // padding
}

func (i *jsiiProxy_ITriggerPredicate) Bind(trigger Trigger) *awsglue.CfnTrigger_ConditionProperty {
	if err := i.validateBindParameters(trigger); err != nil {
		panic(err)
	}
	var returns *awsglue.CfnTrigger_ConditionProperty

	_jsii_.Invoke(
		i,
		"bind",
		[]interface{}{trigger},
		&returns,
	)

	return returns
}

