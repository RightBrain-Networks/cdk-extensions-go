package aps

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/vibe-io/cdk-extensions-go/cdkextensions/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/vibe-io/cdk-extensions-go/cdkextensions/aps/internal"
)

// An inhibition rule mutes an alert (target) matching a set of matchers when an alert (source) exists that matches another set of matchers.
//
// Both target
// and source alerts can optionally be required to have the same label values
// for a specified list of label names.
//
// Semantically, a missing label and a label with an empty value are the same
// thing. Therefore, if all the label names given for `equalLabels` are missing
// from both the source and target alerts, the inhibition rule will apply.
//
// To prevent an alert from inhibiting itself, an alert that matches both the
// target and the source side of a rule cannot be inhibited by alerts for which
// the same is true (including itself). However, we recommend to choose target
// and source matchers in a way that alerts never match both sides. It is much
// easier to reason about and does not trigger this special case.
// See: [Inhibit Rule Official Documentation](https://prometheus.io/docs/alerting/latest/configuration/#inhibit_rule)
//
type AlertManagerInhibitRule interface {
	constructs.Construct
	// Collection of labels that must be equal for both matched source and target alerts in order for the inhibition to take effect.
	EqualLabels() *[]*string
	// The tree node.
	Node() constructs.Node
	// Collection of matchers for which one or more alerts have to exist for the inhibition to take effect.
	SourceMatchers() *[]AlertManagerMatcher
	// Collection of matchers that have to be fulfilled by the target alerts to be muted.
	TargetMatchers() *[]AlertManagerMatcher
	// Adds a label that must be equal for both matched source and target alerts in order for the inhibition to take effect.
	//
	// Returns: The inhibit rule to which the label was added.
	AddEqualLabel(label *string) AlertManagerInhibitRule
	// Adds a matcher for which one or more alerts have to exist in order for inhibition to take effect.
	//
	// Returns: The inhibit rule to which the matcher was added.
	AddSourceMatcher(matcher AlertManagerMatcher) AlertManagerInhibitRule
	// Adds a matcher that has to be fulfilled by a target alert in order to be muted.
	//
	// Returns: The inhibit rule to which the matcher was added.
	AddTargetMatcher(matcher AlertManagerMatcher) AlertManagerInhibitRule
	// Associates the inhibit rule with a construct that is handling the configuration of alert manager.
	//
	// Returns: An alert manager _inhibit_rule` object.
	Bind(_scope constructs.IConstruct) *map[string]interface{}
	// Returns a string representation of this construct.
	ToString() *string
}

// The jsii proxy struct for AlertManagerInhibitRule
type jsiiProxy_AlertManagerInhibitRule struct {
	internal.Type__constructsConstruct
}

func (j *jsiiProxy_AlertManagerInhibitRule) EqualLabels() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"equalLabels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertManagerInhibitRule) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertManagerInhibitRule) SourceMatchers() *[]AlertManagerMatcher {
	var returns *[]AlertManagerMatcher
	_jsii_.Get(
		j,
		"sourceMatchers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertManagerInhibitRule) TargetMatchers() *[]AlertManagerMatcher {
	var returns *[]AlertManagerMatcher
	_jsii_.Get(
		j,
		"targetMatchers",
		&returns,
	)
	return returns
}


// Creates a new instance of the AlertManagerInhibitRule class.
func NewAlertManagerInhibitRule(scope AlertManagerConfiguration, id *string, options AlertManagerInhibitRuleProps) AlertManagerInhibitRule {
	_init_.Initialize()

	if err := validateNewAlertManagerInhibitRuleParameters(scope, id, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_AlertManagerInhibitRule{}

	_jsii_.Create(
		"cdk-extensions.aps.AlertManagerInhibitRule",
		[]interface{}{scope, id, options},
		&j,
	)

	return &j
}

// Creates a new instance of the AlertManagerInhibitRule class.
func NewAlertManagerInhibitRule_Override(a AlertManagerInhibitRule, scope AlertManagerConfiguration, id *string, options AlertManagerInhibitRuleProps) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk-extensions.aps.AlertManagerInhibitRule",
		[]interface{}{scope, id, options},
		a,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead.
func AlertManagerInhibitRule_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAlertManagerInhibitRule_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"cdk-extensions.aps.AlertManagerInhibitRule",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlertManagerInhibitRule) AddEqualLabel(label *string) AlertManagerInhibitRule {
	if err := a.validateAddEqualLabelParameters(label); err != nil {
		panic(err)
	}
	var returns AlertManagerInhibitRule

	_jsii_.Invoke(
		a,
		"addEqualLabel",
		[]interface{}{label},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlertManagerInhibitRule) AddSourceMatcher(matcher AlertManagerMatcher) AlertManagerInhibitRule {
	if err := a.validateAddSourceMatcherParameters(matcher); err != nil {
		panic(err)
	}
	var returns AlertManagerInhibitRule

	_jsii_.Invoke(
		a,
		"addSourceMatcher",
		[]interface{}{matcher},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlertManagerInhibitRule) AddTargetMatcher(matcher AlertManagerMatcher) AlertManagerInhibitRule {
	if err := a.validateAddTargetMatcherParameters(matcher); err != nil {
		panic(err)
	}
	var returns AlertManagerInhibitRule

	_jsii_.Invoke(
		a,
		"addTargetMatcher",
		[]interface{}{matcher},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlertManagerInhibitRule) Bind(_scope constructs.IConstruct) *map[string]interface{} {
	if err := a.validateBindParameters(_scope); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"bind",
		[]interface{}{_scope},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlertManagerInhibitRule) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

