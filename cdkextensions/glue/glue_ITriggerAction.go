package glue

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsglue"
)

type ITriggerAction interface {
	Bind(trigger Trigger) *awsglue.CfnTrigger_ActionProperty
}

// The jsii proxy for ITriggerAction
type jsiiProxy_ITriggerAction struct {
	_ byte // padding
}

func (i *jsiiProxy_ITriggerAction) Bind(trigger Trigger) *awsglue.CfnTrigger_ActionProperty {
	if err := i.validateBindParameters(trigger); err != nil {
		panic(err)
	}
	var returns *awsglue.CfnTrigger_ActionProperty

	_jsii_.Invoke(
		i,
		"bind",
		[]interface{}{trigger},
		&returns,
	)

	return returns
}

