package aps

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/vibe-io/cdk-extensions-go/cdkextensions/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/constructs-go/constructs/v10"
)

// Alerting rules allow you to define alert conditions based on Prometheus expression language expressions and to send notifications about firing alerts to an external service.
//
// Whenever the alert expression results in one
// or more vector elements at a given point in time, the alert counts as active
// for these elements' label sets.
// See: [Alerting rules](https://prometheus.io/docs/prometheus/latest/configuration/alerting_rules/)
//
type AlertingRule interface {
	IPrometheusRule
	// The name of the alert.
	//
	// Must be a valid label value.
	Alert() *string
	// The PromQL expression to evaluate.
	//
	// Every evaluation cycle this is
	// evaluated at the current time, and all resultant time series become
	// pending/firing alerts.
	// See: [Querying prometheus](https://prometheus.io/docs/prometheus/latest/querying/basics/)
	//
	Expression() *string
	// Alerts are considered firing once they have been returned for this long.
	//
	// Alerts which have not yet fired for long enough are considered pending.
	Period() awscdk.Duration
	// Sets a new annotation that will be added to each generated alert.
	//
	// Returns: The alerting rule that the annotation was added to.
	AddAnnotation(label *string, template *string) AlertingRule
	// Sets a new label that will be added or overridden for each generated alert.
	//
	// Returns: The alerting rule that the label was added to.
	AddLabel(label *string, template *string) AlertingRule
	// Associates the alerting rule with a construct that is configuring an APS rule groups namespace.
	//
	// Returns: The rendered configuration for the rule as expected by an APS
	// rules config file.
	Bind(_scope constructs.IConstruct) *map[string]interface{}
}

// The jsii proxy struct for AlertingRule
type jsiiProxy_AlertingRule struct {
	jsiiProxy_IPrometheusRule
}

func (j *jsiiProxy_AlertingRule) Alert() *string {
	var returns *string
	_jsii_.Get(
		j,
		"alert",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertingRule) Expression() *string {
	var returns *string
	_jsii_.Get(
		j,
		"expression",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertingRule) Period() awscdk.Duration {
	var returns awscdk.Duration
	_jsii_.Get(
		j,
		"period",
		&returns,
	)
	return returns
}


// Creates a new instance of the AlertingRule class.
func NewAlertingRule(props *AlertingRuleProps) AlertingRule {
	_init_.Initialize()

	if err := validateNewAlertingRuleParameters(props); err != nil {
		panic(err)
	}
	j := jsiiProxy_AlertingRule{}

	_jsii_.Create(
		"cdk-extensions.aps.AlertingRule",
		[]interface{}{props},
		&j,
	)

	return &j
}

// Creates a new instance of the AlertingRule class.
func NewAlertingRule_Override(a AlertingRule, props *AlertingRuleProps) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk-extensions.aps.AlertingRule",
		[]interface{}{props},
		a,
	)
}

func (a *jsiiProxy_AlertingRule) AddAnnotation(label *string, template *string) AlertingRule {
	if err := a.validateAddAnnotationParameters(label, template); err != nil {
		panic(err)
	}
	var returns AlertingRule

	_jsii_.Invoke(
		a,
		"addAnnotation",
		[]interface{}{label, template},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlertingRule) AddLabel(label *string, template *string) AlertingRule {
	if err := a.validateAddLabelParameters(label, template); err != nil {
		panic(err)
	}
	var returns AlertingRule

	_jsii_.Invoke(
		a,
		"addLabel",
		[]interface{}{label, template},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlertingRule) Bind(_scope constructs.IConstruct) *map[string]interface{} {
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

