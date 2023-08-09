package config

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/constructs-go/constructs/v10"
)

type IRemediationTarget interface {
	Bind(scope constructs.IConstruct) *RemediationTargetConfiguration
}

// The jsii proxy for IRemediationTarget
type jsiiProxy_IRemediationTarget struct {
	_ byte // padding
}

func (i *jsiiProxy_IRemediationTarget) Bind(scope constructs.IConstruct) *RemediationTargetConfiguration {
	if err := i.validateBindParameters(scope); err != nil {
		panic(err)
	}
	var returns *RemediationTargetConfiguration

	_jsii_.Invoke(
		i,
		"bind",
		[]interface{}{scope},
		&returns,
	)

	return returns
}

