package glue

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsglue"
	"github.com/aws/constructs-go/constructs/v10"
)

// Represents an action that should be taken when a trigger is executed.
type ITriggerAction interface {
	Bind(scope constructs.IConstruct) *awsglue.CfnTrigger_ActionProperty
}

// The jsii proxy for ITriggerAction
type jsiiProxy_ITriggerAction struct {
	_ byte // padding
}

func (i *jsiiProxy_ITriggerAction) Bind(scope constructs.IConstruct) *awsglue.CfnTrigger_ActionProperty {
	if err := i.validateBindParameters(scope); err != nil {
		panic(err)
	}
	var returns *awsglue.CfnTrigger_ActionProperty

	_jsii_.Invoke(
		i,
		"bind",
		[]interface{}{scope},
		&returns,
	)

	return returns
}

