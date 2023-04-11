package aps

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/vibe-io/cdk-extensions-go/cdkextensions/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/vibe-io/cdk-extensions-go/cdkextensions/aps/internal"
)

// Represents a named interval of time that may be referenced by the alert manager router to mute/activate particular routes for particular times of day.
// See: [Time Interval Official Documentation](https://prometheus.io/docs/alerting/latest/configuration/#time_interval-0)
//
type TimeInterval interface {
	constructs.Construct
	// Collection of interval entries that defined the the full scope of the periods for which the time interval should apply.
	Intervals() *[]TimeIntervalEntry
	// The name of the time interval as it will be referenced throught the rest of the alert manager configuration.
	Name() *string
	// The tree node.
	Node() constructs.Node
	// Adds a new time interval entry to the time interval.
	//
	// Returns: The time interval which was added.
	AddInterval(interval TimeIntervalEntry) TimeIntervalEntry
	// Associates the time interval with a construct that is handling the configuration of alert manager.
	//
	// Returns: An alert manager `time_interval` configuration object.
	Bind(scope constructs.IConstruct) *map[string]interface{}
	// Returns a string representation of this construct.
	ToString() *string
}

// The jsii proxy struct for TimeInterval
type jsiiProxy_TimeInterval struct {
	internal.Type__constructsConstruct
}

func (j *jsiiProxy_TimeInterval) Intervals() *[]TimeIntervalEntry {
	var returns *[]TimeIntervalEntry
	_jsii_.Get(
		j,
		"intervals",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TimeInterval) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TimeInterval) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


// Creates a new instance of the TimeInterval class.
func NewTimeInterval(scope AlertManagerConfiguration, id *string, options *TimeIntervalProps) TimeInterval {
	_init_.Initialize()

	if err := validateNewTimeIntervalParameters(scope, id, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_TimeInterval{}

	_jsii_.Create(
		"cdk-extensions.aps.TimeInterval",
		[]interface{}{scope, id, options},
		&j,
	)

	return &j
}

// Creates a new instance of the TimeInterval class.
func NewTimeInterval_Override(t TimeInterval, scope AlertManagerConfiguration, id *string, options *TimeIntervalProps) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk-extensions.aps.TimeInterval",
		[]interface{}{scope, id, options},
		t,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead.
func TimeInterval_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateTimeInterval_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"cdk-extensions.aps.TimeInterval",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TimeInterval) AddInterval(interval TimeIntervalEntry) TimeIntervalEntry {
	if err := t.validateAddIntervalParameters(interval); err != nil {
		panic(err)
	}
	var returns TimeIntervalEntry

	_jsii_.Invoke(
		t,
		"addInterval",
		[]interface{}{interval},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TimeInterval) Bind(scope constructs.IConstruct) *map[string]interface{} {
	if err := t.validateBindParameters(scope); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		t,
		"bind",
		[]interface{}{scope},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TimeInterval) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

