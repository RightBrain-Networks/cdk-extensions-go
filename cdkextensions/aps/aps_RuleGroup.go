package aps

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/vibe-io/cdk-extensions-go/cdkextensions/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/vibe-io/cdk-extensions-go/cdkextensions/aps/internal"
)

// A group of alerting and recording rules for use inside an APS rule groups namespace configuration.
//
// Rules within a group are run sequentially at a
// regular interval, with the same evaluation time. The names of recording
// rules must be valid metric names. The names of alerting rules must be valid
// label values.
// See: [Prometheus rule_group specification](https://prometheus.io/docs/prometheus/latest/configuration/recording_rules/#rule_group)
//
type RuleGroup interface {
	constructs.Construct
	// How often rules in the group are evaluated.
	Interval() awscdk.Duration
	// Limit the number of alerts an alerting rule and series a recording rule can produce.
	//
	// 0 is no limit.
	Limit() *float64
	// The name of the group.
	//
	// Must be unique within the configuration.
	Name() *string
	// The tree node.
	Node() constructs.Node
	// Adds an alerting rule to the rule group.
	//
	// Returns: The alerting rule that was added.
	AddAlertingRule(options *AlertingRuleProps) AlertingRule
	// Adds a recording rule to the rule group.
	//
	// Returns: The recording rule that was added.
	AddRecordingRule(options *RecordingRuleProps) RecordingRule
	// Adds a Prometheus rule to the rule group.
	//
	// This method can be used to implement any rules that are created outside of
	// the the rule group that were created manually by calling their
	// constructors or for adding Prometheus rules that use their own custom
	// implementations.
	//
	// Returns: The rule group that the rule was added to.
	AddRule(rule IPrometheusRule) RuleGroup
	// Associates the rule group with a construct that is configuring an APS rule groups namespace.
	//
	// Returns: The rendered configuration for the rule group as expected by an
	// APS rules config file.
	Bind(scope constructs.IConstruct) *map[string]interface{}
	// Returns a string representation of this construct.
	ToString() *string
}

// The jsii proxy struct for RuleGroup
type jsiiProxy_RuleGroup struct {
	internal.Type__constructsConstruct
}

func (j *jsiiProxy_RuleGroup) Interval() awscdk.Duration {
	var returns awscdk.Duration
	_jsii_.Get(
		j,
		"interval",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RuleGroup) Limit() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"limit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RuleGroup) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RuleGroup) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


// Creates a new instance of the RuleGroup class.
func NewRuleGroup(scope constructs.IConstruct, id *string, props *RuleGroupProps) RuleGroup {
	_init_.Initialize()

	if err := validateNewRuleGroupParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_RuleGroup{}

	_jsii_.Create(
		"cdk-extensions.aps.RuleGroup",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Creates a new instance of the RuleGroup class.
func NewRuleGroup_Override(r RuleGroup, scope constructs.IConstruct, id *string, props *RuleGroupProps) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk-extensions.aps.RuleGroup",
		[]interface{}{scope, id, props},
		r,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead.
func RuleGroup_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateRuleGroup_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"cdk-extensions.aps.RuleGroup",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RuleGroup) AddAlertingRule(options *AlertingRuleProps) AlertingRule {
	if err := r.validateAddAlertingRuleParameters(options); err != nil {
		panic(err)
	}
	var returns AlertingRule

	_jsii_.Invoke(
		r,
		"addAlertingRule",
		[]interface{}{options},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RuleGroup) AddRecordingRule(options *RecordingRuleProps) RecordingRule {
	if err := r.validateAddRecordingRuleParameters(options); err != nil {
		panic(err)
	}
	var returns RecordingRule

	_jsii_.Invoke(
		r,
		"addRecordingRule",
		[]interface{}{options},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RuleGroup) AddRule(rule IPrometheusRule) RuleGroup {
	if err := r.validateAddRuleParameters(rule); err != nil {
		panic(err)
	}
	var returns RuleGroup

	_jsii_.Invoke(
		r,
		"addRule",
		[]interface{}{rule},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RuleGroup) Bind(scope constructs.IConstruct) *map[string]interface{} {
	if err := r.validateBindParameters(scope); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		r,
		"bind",
		[]interface{}{scope},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RuleGroup) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

