package alerting

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/vibe-io/cdk-extensions-go/cdkextensions/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssecretsmanager"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsstepfunctions"
	"github.com/aws/constructs-go/constructs/v10"
)

// A standardized implementation that allows Jira tickets to be created in response to events detected in AWS.
//
// Intended for use with the `IssueManager` state machine which allows
// arbitrary types of events to be processed into standard values and then
// output or one of more issue tracking services.
// See: [AWS-CreateJiraIssue](https://docs.aws.amazon.com/systems-manager-automation-runbooks/latest/userguide/automation-aws-createjiraissue.html)
//
type JiraTicket interface {
	IssuePluginBase
	IIssueHandler
	// Destination pointing to a Jira instance where tickets are to be created.
	ApiDestination() awsevents.ApiDestination
	// The default assignee that issues should be created with if no other assignee is specified by the event that triggered the issue creation.
	Assignee() *string
	// API connection providing details of how to communicate with the configured Jira instance.
	Connection() awsevents.Connection
	// The credentials to be used for connecting to Jira. The secret should be in JSON format and contain the key:.
	//
	// username: The name of the user issues should be created as.
	// password: A password or API key for the user specified in `username`.
	Credentials() awssecretsmanager.ISecret
	// The environment this resource belongs to.
	//
	// For resources that are created and managed by the CDK
	// (generally, those created by creating new class instances like Role, Bucket, etc.),
	// this is always the same as the environment of the stack they belong to;
	// however, for imported resources
	// (those obtained from static methods like fromRoleArn, fromBucketName, etc.),
	// that might be different than the stack they were imported into.
	Env() *awscdk.ResourceEnvironment
	// The event bus to use to trigger writes to the Jira instance.
	//
	// This integration formats a Jira API response and then sends it to a Jira
	// instance by means of an EventBridge Destination API and a specially
	// crafted event pattern. This is the event bus where the rule to trigger the
	// API will be added and the trigger event will be sent.
	EventBus() awsevents.IEventBus
	// The State Machine that handles creating a Jira ticket for a passed issue.
	//
	// Internally this state machine uses the AWS managed `AWS-CreateJiraIssue`
	// SSM Automation document.
	Handler() awsstepfunctions.IStateMachine
	// The default issue type that issues should be created as if no other type is specified by the event that triggered the issue creation.
	IssueType() *string
	// The URL of the Jira instance where tickets should be created.
	JiraUrl() *string
	Logging() *StateMachineLogging
	// The human friendly name that can be used to identify the plugin.
	Name() *string
	// The tree node.
	Node() constructs.Node
	// Returns a string-encoded token that resolves to the physical name that should be passed to the CloudFormation resource.
	//
	// This value will resolve to one of the following:
	// - a concrete value (e.g. `"my-awesome-bucket"`)
	// - `undefined`, when a name should be generated by CloudFormation
	// - a concrete name generated automatically during synthesis, in
	//   cross-environment scenarios.
	PhysicalName() *string
	// A mapping of the standard severities supported by issue manager to priority levels supported by the destination Jira instance.
	PriorityMap() *JiraTicketPriorityMap
	// The name of the default project to use for creating issues if no other project is specified by the event that triggered the issue creation.
	Project() *string
	// The stack in which this resource is defined.
	Stack() awscdk.Stack
	// The length of time that the State Machine that handles creation of Jira tickets is allowed to run before timing out.
	Timeout() awscdk.Duration
	// Apply the given removal policy to this resource.
	//
	// The Removal Policy controls what happens to this resource when it stops
	// being managed by CloudFormation, either because you've removed it from the
	// CDK application or because you've made a change that requires the resource
	// to be replaced.
	//
	// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
	// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy)
	BuildEventOverrides(options *JiraTicketOverrideOptions) IssueHandlerOverride
	BuildLogging() *awsstepfunctions.LogOptions
	BuildSeverityMap() awsstepfunctions.Chain
	GeneratePhysicalName() *string
	// Returns an environment-sensitive token that should be used for the resource's "ARN" attribute (e.g. `bucket.bucketArn`).
	//
	// Normally, this token will resolve to `arnAttr`, but if the resource is
	// referenced across environments, `arnComponents` will be used to synthesize
	// a concrete ARN with the resource's physical name. Make sure to reference
	// `this.physicalName` in `arnComponents`.
	GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string
	// Returns an environment-sensitive token that should be used for the resource's "name" attribute (e.g. `bucket.bucketName`).
	//
	// Normally, this token will resolve to `nameAttr`, but if the resource is
	// referenced across environments, it will be resolved to `this.physicalName`,
	// which will be a concrete name.
	GetResourceNameAttribute(nameAttr *string) *string
	// Returns a string representation of this construct.
	ToString() *string
}

// The jsii proxy struct for JiraTicket
type jsiiProxy_JiraTicket struct {
	jsiiProxy_IssuePluginBase
	jsiiProxy_IIssueHandler
}

func (j *jsiiProxy_JiraTicket) ApiDestination() awsevents.ApiDestination {
	var returns awsevents.ApiDestination
	_jsii_.Get(
		j,
		"apiDestination",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JiraTicket) Assignee() *string {
	var returns *string
	_jsii_.Get(
		j,
		"assignee",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JiraTicket) Connection() awsevents.Connection {
	var returns awsevents.Connection
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JiraTicket) Credentials() awssecretsmanager.ISecret {
	var returns awssecretsmanager.ISecret
	_jsii_.Get(
		j,
		"credentials",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JiraTicket) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JiraTicket) EventBus() awsevents.IEventBus {
	var returns awsevents.IEventBus
	_jsii_.Get(
		j,
		"eventBus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JiraTicket) Handler() awsstepfunctions.IStateMachine {
	var returns awsstepfunctions.IStateMachine
	_jsii_.Get(
		j,
		"handler",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JiraTicket) IssueType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"issueType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JiraTicket) JiraUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"jiraUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JiraTicket) Logging() *StateMachineLogging {
	var returns *StateMachineLogging
	_jsii_.Get(
		j,
		"logging",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JiraTicket) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JiraTicket) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JiraTicket) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JiraTicket) PriorityMap() *JiraTicketPriorityMap {
	var returns *JiraTicketPriorityMap
	_jsii_.Get(
		j,
		"priorityMap",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JiraTicket) Project() *string {
	var returns *string
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JiraTicket) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JiraTicket) Timeout() awscdk.Duration {
	var returns awscdk.Duration
	_jsii_.Get(
		j,
		"timeout",
		&returns,
	)
	return returns
}


// Creates a new instance of the JiraTicket class.
func NewJiraTicket(scope constructs.IConstruct, id *string, props *JiraTicketProps) JiraTicket {
	_init_.Initialize()

	if err := validateNewJiraTicketParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_JiraTicket{}

	_jsii_.Create(
		"cdk-extensions.alerting.JiraTicket",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Creates a new instance of the JiraTicket class.
func NewJiraTicket_Override(j JiraTicket, scope constructs.IConstruct, id *string, props *JiraTicketProps) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk-extensions.alerting.JiraTicket",
		[]interface{}{scope, id, props},
		j,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead.
func JiraTicket_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateJiraTicket_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"cdk-extensions.alerting.JiraTicket",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Returns true if the construct was created by CDK, and false otherwise.
func JiraTicket_IsOwnedResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	if err := validateJiraTicket_IsOwnedResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"cdk-extensions.alerting.JiraTicket",
		"isOwnedResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
func JiraTicket_IsResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	if err := validateJiraTicket_IsResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"cdk-extensions.alerting.JiraTicket",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

func JiraTicket_DEFAULT_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"cdk-extensions.alerting.JiraTicket",
		"DEFAULT_NAME",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JiraTicket) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	if err := j.validateApplyRemovalPolicyParameters(policy); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (j *jsiiProxy_JiraTicket) BuildEventOverrides(options *JiraTicketOverrideOptions) IssueHandlerOverride {
	if err := j.validateBuildEventOverridesParameters(options); err != nil {
		panic(err)
	}
	var returns IssueHandlerOverride

	_jsii_.Invoke(
		j,
		"buildEventOverrides",
		[]interface{}{options},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JiraTicket) BuildLogging() *awsstepfunctions.LogOptions {
	var returns *awsstepfunctions.LogOptions

	_jsii_.Invoke(
		j,
		"buildLogging",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JiraTicket) BuildSeverityMap() awsstepfunctions.Chain {
	var returns awsstepfunctions.Chain

	_jsii_.Invoke(
		j,
		"buildSeverityMap",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JiraTicket) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		j,
		"generatePhysicalName",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JiraTicket) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	if err := j.validateGetResourceArnAttributeParameters(arnAttr, arnComponents); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		j,
		"getResourceArnAttribute",
		[]interface{}{arnAttr, arnComponents},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JiraTicket) GetResourceNameAttribute(nameAttr *string) *string {
	if err := j.validateGetResourceNameAttributeParameters(nameAttr); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		j,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JiraTicket) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		j,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

