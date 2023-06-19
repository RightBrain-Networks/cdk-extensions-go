package resourcegroups

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/vibe-io/cdk-extensions-go/cdkextensions/jsii"
)

type GroupConfiguration interface {
}

// The jsii proxy struct for GroupConfiguration
type jsiiProxy_GroupConfiguration struct {
	_ byte // padding
}

func NewGroupConfiguration() GroupConfiguration {
	_init_.Initialize()

	j := jsiiProxy_GroupConfiguration{}

	_jsii_.Create(
		"cdk-extensions.resourcegroups.GroupConfiguration",
		nil, // no parameters
		&j,
	)

	return &j
}

func NewGroupConfiguration_Override(g GroupConfiguration) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk-extensions.resourcegroups.GroupConfiguration",
		nil, // no parameters
		g,
	)
}

func GroupConfiguration_CloudFormationStack(reference IStackReference, props *CloudFormationStackProps) CloudFormationStack {
	_init_.Initialize()

	if err := validateGroupConfiguration_CloudFormationStackParameters(reference, props); err != nil {
		panic(err)
	}
	var returns CloudFormationStack

	_jsii_.StaticInvoke(
		"cdk-extensions.resourcegroups.GroupConfiguration",
		"cloudFormationStack",
		[]interface{}{reference, props},
		&returns,
	)

	return returns
}

func GroupConfiguration_TagFilter(props TagFilterProps) TagFilter {
	_init_.Initialize()

	var returns TagFilter

	_jsii_.StaticInvoke(
		"cdk-extensions.resourcegroups.GroupConfiguration",
		"tagFilter",
		[]interface{}{props},
		&returns,
	)

	return returns
}

