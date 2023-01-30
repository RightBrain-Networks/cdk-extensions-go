package aps

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/constructs-go/constructs/v10"
)

// Represents a Prometheus rule that can be added to a rule group and used as part of an APS rules configuration.
type IPrometheusRule interface {
	// Associates the rule with a construct that is configuring an APS rule groups namespace.
	Bind(scope constructs.IConstruct) *map[string]interface{}
}

// The jsii proxy for IPrometheusRule
type jsiiProxy_IPrometheusRule struct {
	_ byte // padding
}

func (i *jsiiProxy_IPrometheusRule) Bind(scope constructs.IConstruct) *map[string]interface{} {
	if err := i.validateBindParameters(scope); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		i,
		"bind",
		[]interface{}{scope},
		&returns,
	)

	return returns
}

