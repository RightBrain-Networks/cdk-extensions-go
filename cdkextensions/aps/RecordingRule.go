package aps

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/vibe-io/cdk-extensions-go/cdkextensions/jsii"

	"github.com/aws/constructs-go/constructs/v10"
)

// Recording rules allow you to precompute frequently needed or computationally expensive expressions and save their result as a new set of time series.
//
// Querying the precomputed result will then often be much faster than
// executing the original expression every time it is needed. This is
// especially useful for dashboards, which need to query the same expression
// repeatedly every time they refresh.
// See: [Defining recording rules](https://prometheus.io/docs/prometheus/latest/configuration/recording_rules/)
//
type RecordingRule interface {
	IPrometheusRule
	// The PromQL expression to evaluate.
	//
	// Every evaluation cycle this is
	// evaluated at the current time, and the result recorded as a new set of
	// time series with the metric name as given by `record`.
	// See: [Querying prometheus](https://prometheus.io/docs/prometheus/latest/querying/basics/)
	//
	Expression() *string
	// Labels to add or overwrite before storing the result.
	Labels() *map[string]*string
	// The name of the time series to output to.
	//
	// Must be a valid metric name.
	Record() *string
	// Sets a new label and value that will be added or overridden before storing the result.
	//
	// Returns: The recording rule that the label was added to.
	AddLabel(label *string, value *string) RecordingRule
	// Associates the recording rule with a construct that is configuring an APS rule groups namespace.
	//
	// Returns: The rendered configuration for the rule as expected by and APS
	// rules config file.
	Bind(_scope constructs.IConstruct) *map[string]interface{}
}

// The jsii proxy struct for RecordingRule
type jsiiProxy_RecordingRule struct {
	jsiiProxy_IPrometheusRule
}

func (j *jsiiProxy_RecordingRule) Expression() *string {
	var returns *string
	_jsii_.Get(
		j,
		"expression",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RecordingRule) Labels() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"labels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RecordingRule) Record() *string {
	var returns *string
	_jsii_.Get(
		j,
		"record",
		&returns,
	)
	return returns
}


// Creates a new instance of the RecordingRule class.
func NewRecordingRule(props *RecordingRuleProps) RecordingRule {
	_init_.Initialize()

	if err := validateNewRecordingRuleParameters(props); err != nil {
		panic(err)
	}
	j := jsiiProxy_RecordingRule{}

	_jsii_.Create(
		"cdk-extensions.aps.RecordingRule",
		[]interface{}{props},
		&j,
	)

	return &j
}

// Creates a new instance of the RecordingRule class.
func NewRecordingRule_Override(r RecordingRule, props *RecordingRuleProps) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk-extensions.aps.RecordingRule",
		[]interface{}{props},
		r,
	)
}

func (r *jsiiProxy_RecordingRule) AddLabel(label *string, value *string) RecordingRule {
	if err := r.validateAddLabelParameters(label, value); err != nil {
		panic(err)
	}
	var returns RecordingRule

	_jsii_.Invoke(
		r,
		"addLabel",
		[]interface{}{label, value},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RecordingRule) Bind(_scope constructs.IConstruct) *map[string]interface{} {
	if err := r.validateBindParameters(_scope); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		r,
		"bind",
		[]interface{}{_scope},
		&returns,
	)

	return returns
}

