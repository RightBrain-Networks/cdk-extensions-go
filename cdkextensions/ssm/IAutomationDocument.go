package ssm

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

type IAutomationDocument interface {
	IDocument
	ArnForAutomationDefinitionVersion(version *string) *string
	AutomationDefinitionArn() *string
}

// The jsii proxy for IAutomationDocument
type jsiiProxy_IAutomationDocument struct {
	jsiiProxy_IDocument
}

func (i *jsiiProxy_IAutomationDocument) ArnForAutomationDefinitionVersion(version *string) *string {
	if err := i.validateArnForAutomationDefinitionVersionParameters(version); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		i,
		"arnForAutomationDefinitionVersion",
		[]interface{}{version},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_IAutomationDocument) AutomationDefinitionArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"automationDefinitionArn",
		&returns,
	)
	return returns
}

