package aps

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/vibe-io/cdk-extensions-go/cdkextensions/aps/internal"
)

// Represents a rules configuration that can be consumed by Amazon Managed Service for Prometheus when creating a rule groups namespace.
type IRuleGroupConfiguration interface {
	constructs.IConstruct
	// Associates the configuration with a resource that is handling the creation of an APS rule groups namespace.
	//
	// Returns: Rule group configuration details.
	Bind(scope constructs.IConstruct) *RuleGroupConfigurationDetails
}

// The jsii proxy for IRuleGroupConfiguration
type jsiiProxy_IRuleGroupConfiguration struct {
	internal.Type__constructsIConstruct
}

func (i *jsiiProxy_IRuleGroupConfiguration) Bind(scope constructs.IConstruct) *RuleGroupConfigurationDetails {
	if err := i.validateBindParameters(scope); err != nil {
		panic(err)
	}
	var returns *RuleGroupConfigurationDetails

	_jsii_.Invoke(
		i,
		"bind",
		[]interface{}{scope},
		&returns,
	)

	return returns
}

