package aps

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/vibe-io/cdk-extensions-go/cdkextensions/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2/awssns"
	"github.com/aws/constructs-go/constructs/v10"
)

// An alert manager destination that provides details for sending alert notifications to an Amazon SNS topic.
// See: [Alert manager SNS configuration](https://prometheus.io/docs/alerting/latest/configuration/#sns_config)
//
type AlertManagerSnsDestination interface {
	IAlertManagerDestination
	// The SNS API URL i.e. https://sns.us-east-2.amazonaws.com.
	//
	// If not specified, the SNS API URL from the SNS SDK will be used.
	ApiUrl() *string
	// Attributes to be applied to SNS messages.
	Attributes() *map[string]*string
	// The destination type being configured.
	//
	// Represents a config block in an alert manager receiver configuration.
	Category() AlertManagerDestinationCategory
	// The message content of the SNS notification.
	Message() *string
	// Controls whether to notify about resolved alerts.
	SendResolved() *bool
	// Subject line when the message is delivered to email endpoints.
	Subject() *string
	// SNS topic where alerts will be sent.
	//
	// If you are using a FIFO SNS topic you should set a message group interval
	// longer than 5 minutes to prevent messages with the same group key being
	// deduplicated by the SNS default deduplication window.
	Topic() awssns.ITopic
	// Registers a new attribute to be added to sent SNS messages.
	//
	// Returns: The sns destination object the attribute was added to.
	AddAttribute(key *string, value *string) AlertManagerSnsDestination
	// Associates the destination with a construct that is handling the configuration of alert manager.
	//
	// Returns: An object representing the configured destination.
	Bind(scope constructs.IConstruct) *map[string]interface{}
}

// The jsii proxy struct for AlertManagerSnsDestination
type jsiiProxy_AlertManagerSnsDestination struct {
	jsiiProxy_IAlertManagerDestination
}

func (j *jsiiProxy_AlertManagerSnsDestination) ApiUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertManagerSnsDestination) Attributes() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"attributes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertManagerSnsDestination) Category() AlertManagerDestinationCategory {
	var returns AlertManagerDestinationCategory
	_jsii_.Get(
		j,
		"category",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertManagerSnsDestination) Message() *string {
	var returns *string
	_jsii_.Get(
		j,
		"message",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertManagerSnsDestination) SendResolved() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"sendResolved",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertManagerSnsDestination) Subject() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subject",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertManagerSnsDestination) Topic() awssns.ITopic {
	var returns awssns.ITopic
	_jsii_.Get(
		j,
		"topic",
		&returns,
	)
	return returns
}


// Creates a new instance of the AlertManagerSnsDestination class.
func NewAlertManagerSnsDestination(options *AlertManagerSnsDestinationProps) AlertManagerSnsDestination {
	_init_.Initialize()

	if err := validateNewAlertManagerSnsDestinationParameters(options); err != nil {
		panic(err)
	}
	j := jsiiProxy_AlertManagerSnsDestination{}

	_jsii_.Create(
		"cdk-extensions.aps.AlertManagerSnsDestination",
		[]interface{}{options},
		&j,
	)

	return &j
}

// Creates a new instance of the AlertManagerSnsDestination class.
func NewAlertManagerSnsDestination_Override(a AlertManagerSnsDestination, options *AlertManagerSnsDestinationProps) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk-extensions.aps.AlertManagerSnsDestination",
		[]interface{}{options},
		a,
	)
}

func (a *jsiiProxy_AlertManagerSnsDestination) AddAttribute(key *string, value *string) AlertManagerSnsDestination {
	if err := a.validateAddAttributeParameters(key, value); err != nil {
		panic(err)
	}
	var returns AlertManagerSnsDestination

	_jsii_.Invoke(
		a,
		"addAttribute",
		[]interface{}{key, value},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlertManagerSnsDestination) Bind(scope constructs.IConstruct) *map[string]interface{} {
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

