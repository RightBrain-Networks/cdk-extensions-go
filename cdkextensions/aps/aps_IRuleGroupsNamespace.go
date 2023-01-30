package aps

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Represents an APS rule groups namespace in AWS.
type IRuleGroupsNamespace interface {
	// The Amazon Resource Name (ARN) of the APS rule groups namespace.
	RulesGroupsNamespaceArn() *string
}

// The jsii proxy for IRuleGroupsNamespace
type jsiiProxy_IRuleGroupsNamespace struct {
	_ byte // padding
}

func (j *jsiiProxy_IRuleGroupsNamespace) RulesGroupsNamespaceArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rulesGroupsNamespaceArn",
		&returns,
	)
	return returns
}

