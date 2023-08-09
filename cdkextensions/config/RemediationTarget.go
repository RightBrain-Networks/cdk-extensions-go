package config

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/vibe-io/cdk-extensions-go/cdkextensions/jsii"
)

type RemediationTarget interface {
}

// The jsii proxy struct for RemediationTarget
type jsiiProxy_RemediationTarget struct {
	_ byte // padding
}

func NewRemediationTarget() RemediationTarget {
	_init_.Initialize()

	j := jsiiProxy_RemediationTarget{}

	_jsii_.Create(
		"cdk-extensions.config.RemediationTarget",
		nil, // no parameters
		&j,
	)

	return &j
}

func NewRemediationTarget_Override(r RemediationTarget) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk-extensions.config.RemediationTarget",
		nil, // no parameters
		r,
	)
}

func RemediationTarget_AutomationDocument(props *AutomationDocumentRemediationProps) IRemediationTarget {
	_init_.Initialize()

	if err := validateRemediationTarget_AutomationDocumentParameters(props); err != nil {
		panic(err)
	}
	var returns IRemediationTarget

	_jsii_.StaticInvoke(
		"cdk-extensions.config.RemediationTarget",
		"automationDocument",
		[]interface{}{props},
		&returns,
	)

	return returns
}

