package aps

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/vibe-io/cdk-extensions-go/cdkextensions/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/vibe-io/cdk-extensions-go/cdkextensions/aps/internal"
)

// Represents a rules file definition that can be consumed by Amazon Managed Service for Prometheus.
// See: [Creating a rules file](https://docs.aws.amazon.com/prometheus/latest/userguide/AMP-ruler-rulesfile.html)
//
type RuleGroupConfiguration interface {
	constructs.Construct
	IRuleGroupConfiguration
	// The tree node.
	Node() constructs.Node
	// Collection of rule groups that are part of this configuration.
	RuleGroups() *[]RuleGroup
	// Adds a new rule group to the configuration.
	//
	// Returns: The rule group that was added.
	AddRuleGroup(id *string, options *RuleGroupProps) RuleGroup
	// Associates the configuration with a resource that is handling the creation of an APS rule groups namespace.
	//
	// Returns: Rule group configuration details.
	Bind(scope constructs.IConstruct) *RuleGroupConfigurationDetails
	// Returns a string representation of this construct.
	ToString() *string
}

// The jsii proxy struct for RuleGroupConfiguration
type jsiiProxy_RuleGroupConfiguration struct {
	internal.Type__constructsConstruct
	jsiiProxy_IRuleGroupConfiguration
}

func (j *jsiiProxy_RuleGroupConfiguration) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RuleGroupConfiguration) RuleGroups() *[]RuleGroup {
	var returns *[]RuleGroup
	_jsii_.Get(
		j,
		"ruleGroups",
		&returns,
	)
	return returns
}


// Creates a new instance of the RuleGroupConfiguration class.
func NewRuleGroupConfiguration(scope constructs.IConstruct, id *string, _props *RuleGroupConfigurationProps) RuleGroupConfiguration {
	_init_.Initialize()

	if err := validateNewRuleGroupConfigurationParameters(scope, id, _props); err != nil {
		panic(err)
	}
	j := jsiiProxy_RuleGroupConfiguration{}

	_jsii_.Create(
		"cdk-extensions.aps.RuleGroupConfiguration",
		[]interface{}{scope, id, _props},
		&j,
	)

	return &j
}

// Creates a new instance of the RuleGroupConfiguration class.
func NewRuleGroupConfiguration_Override(r RuleGroupConfiguration, scope constructs.IConstruct, id *string, _props *RuleGroupConfigurationProps) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk-extensions.aps.RuleGroupConfiguration",
		[]interface{}{scope, id, _props},
		r,
	)
}

// Imports an APS rules configuration using a raw string.
//
// The string should be in YAML format and follow the specification expected
// by the `aps:CreateRuleGroupsNamespace` API call.
//
// Returns: An object that can be used to configure a rule groups namespace
// for APS.
// See: [Rules file specification](https://docs.aws.amazon.com/prometheus/latest/userguide/AMP-ruler-rulesfile.html)
//
func RuleGroupConfiguration_FromContent(scope constructs.IConstruct, id *string, content *string) IRuleGroupConfiguration {
	_init_.Initialize()

	if err := validateRuleGroupConfiguration_FromContentParameters(scope, id, content); err != nil {
		panic(err)
	}
	var returns IRuleGroupConfiguration

	_jsii_.StaticInvoke(
		"cdk-extensions.aps.RuleGroupConfiguration",
		"fromContent",
		[]interface{}{scope, id, content},
		&returns,
	)

	return returns
}

// Imports an APS rules file from the local filesystem.
//
// The file should be in YAML format and follow the specification expected by
// the `aps:CreateRuleGroupsNamespace` API call.
//
// Returns: An object that can be used to configure a rule groups namespace
// for APS.
// See: [Rules file specification](https://docs.aws.amazon.com/prometheus/latest/userguide/AMP-ruler-rulesfile.html)
//
func RuleGroupConfiguration_FromRulesFile(scope constructs.IConstruct, id *string, path *string) IRuleGroupConfiguration {
	_init_.Initialize()

	if err := validateRuleGroupConfiguration_FromRulesFileParameters(scope, id, path); err != nil {
		panic(err)
	}
	var returns IRuleGroupConfiguration

	_jsii_.StaticInvoke(
		"cdk-extensions.aps.RuleGroupConfiguration",
		"fromRulesFile",
		[]interface{}{scope, id, path},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead.
func RuleGroupConfiguration_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateRuleGroupConfiguration_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"cdk-extensions.aps.RuleGroupConfiguration",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RuleGroupConfiguration) AddRuleGroup(id *string, options *RuleGroupProps) RuleGroup {
	if err := r.validateAddRuleGroupParameters(id, options); err != nil {
		panic(err)
	}
	var returns RuleGroup

	_jsii_.Invoke(
		r,
		"addRuleGroup",
		[]interface{}{id, options},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RuleGroupConfiguration) Bind(scope constructs.IConstruct) *RuleGroupConfigurationDetails {
	if err := r.validateBindParameters(scope); err != nil {
		panic(err)
	}
	var returns *RuleGroupConfigurationDetails

	_jsii_.Invoke(
		r,
		"bind",
		[]interface{}{scope},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RuleGroupConfiguration) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

