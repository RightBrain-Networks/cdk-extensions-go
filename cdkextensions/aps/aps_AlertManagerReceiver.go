package aps

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/vibe-io/cdk-extensions-go/cdkextensions/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2/awssns"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/vibe-io/cdk-extensions-go/cdkextensions/aps/internal"
)

// Represents a notification integration within alert manager.
// See: [Receiver configuration](https://prometheus.io/docs/alerting/latest/configuration/#receiver)
//
type AlertManagerReceiver interface {
	constructs.Construct
	// Collection of destinations which define details for alerting providers where events routed to this receiver should be sent.
	Destinations() *[]IAlertManagerDestination
	// The name of the receiver which can be referenced in the other parts of the configuration.
	Name() *string
	// The tree node.
	Node() constructs.Node
	// Adds a new generic destination that should receive alerts that have been routed to this receiver.
	//
	// Returns: The receiver where the destination was added.
	AddDestination(destination IAlertManagerDestination) AlertManagerReceiver
	// Adds a new SNS destination that should receive alerts that have been routed to this receiver.
	//
	// Returns: The SnsDestination object that was added to the receiver.
	AddSnsTopic(topic awssns.ITopic, options *AlertManagerSnsDestinationOptions) AlertManagerSnsDestination
	// Associates the receiver with a construct that is handling the configuration of alert manager that will consume the configuration.
	//
	// Returns: An alert manager `receiver` configuration object.
	Bind(scope constructs.IConstruct) *map[string]interface{}
	// Returns a string representation of this construct.
	ToString() *string
}

// The jsii proxy struct for AlertManagerReceiver
type jsiiProxy_AlertManagerReceiver struct {
	internal.Type__constructsConstruct
}

func (j *jsiiProxy_AlertManagerReceiver) Destinations() *[]IAlertManagerDestination {
	var returns *[]IAlertManagerDestination
	_jsii_.Get(
		j,
		"destinations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertManagerReceiver) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertManagerReceiver) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


// Creates a new instance of the AlertManagerReceiver class.
func NewAlertManagerReceiver(scope AlertManagerConfiguration, id *string, props *AlertManagerReceiverProps) AlertManagerReceiver {
	_init_.Initialize()

	if err := validateNewAlertManagerReceiverParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_AlertManagerReceiver{}

	_jsii_.Create(
		"cdk-extensions.aps.AlertManagerReceiver",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Creates a new instance of the AlertManagerReceiver class.
func NewAlertManagerReceiver_Override(a AlertManagerReceiver, scope AlertManagerConfiguration, id *string, props *AlertManagerReceiverProps) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk-extensions.aps.AlertManagerReceiver",
		[]interface{}{scope, id, props},
		a,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead.
func AlertManagerReceiver_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAlertManagerReceiver_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"cdk-extensions.aps.AlertManagerReceiver",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlertManagerReceiver) AddDestination(destination IAlertManagerDestination) AlertManagerReceiver {
	if err := a.validateAddDestinationParameters(destination); err != nil {
		panic(err)
	}
	var returns AlertManagerReceiver

	_jsii_.Invoke(
		a,
		"addDestination",
		[]interface{}{destination},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlertManagerReceiver) AddSnsTopic(topic awssns.ITopic, options *AlertManagerSnsDestinationOptions) AlertManagerSnsDestination {
	if err := a.validateAddSnsTopicParameters(topic, options); err != nil {
		panic(err)
	}
	var returns AlertManagerSnsDestination

	_jsii_.Invoke(
		a,
		"addSnsTopic",
		[]interface{}{topic, options},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlertManagerReceiver) Bind(scope constructs.IConstruct) *map[string]interface{} {
	if err := a.validateBindParameters(scope); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"bind",
		[]interface{}{scope},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlertManagerReceiver) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

