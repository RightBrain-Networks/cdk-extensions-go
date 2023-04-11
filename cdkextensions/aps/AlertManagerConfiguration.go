package aps

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/vibe-io/cdk-extensions-go/cdkextensions/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2/awssns"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/vibe-io/cdk-extensions-go/cdkextensions/aps/internal"
)

// Represents a Prometheus alert manager configuration that can be used by Amazon APS to send generated alerts to one or more destinations.
//
// Currently
// the only destination type supported by APS is Amazon SNS.
// See: [Alert manager configuration specification](https://prometheus.io/docs/alerting/latest/configuration/)
//
type AlertManagerConfiguration interface {
	constructs.Construct
	IAlertManagerConfiguration
	// The receiver configuring the destinations where all alerts that are not matched by a child route in the routing tree will be sent.
	// See: [Receiver Official Documentation](https://prometheus.io/docs/alerting/latest/configuration/#receiver)
	//
	DefaultReceiver() AlertManagerReceiver
	// The default root route withing the routing tree.
	//
	// Serves as the default
	// configuration for all alerts. Sends alerts to the default receiver.
	//
	// Must match all alerts and cannot be muted or deactivated. To add more
	// advanced configurations create additional routes as children of the
	// default route.
	// See: [Route OfficialDocumentation](https://prometheus.io/docs/alerting/latest/configuration/#route)
	//
	DefaultRoute() AlertManagerRoute
	// The topic that the default receiver will send messages to if no destinations are specified as inputs (input destinations are null/undefined).
	//
	// The creation of this topic can be skipped to allow the defferred
	// configuration of default destination by providing an empty list of
	// destinations as inputs.
	DefaultTopic() awssns.ITopic
	// Collection of inhibit rules that mute an alert under scecific sets of circumstances.
	// See: [Inhibit Rule Official Documentation](https://prometheus.io/docs/alerting/latest/configuration/#inhibit_rule)
	//
	InhibitRules() *[]AlertManagerInhibitRule
	// The tree node.
	Node() constructs.Node
	// Collection of notification integrations that provide the details for how and where generated alerts should be sent.
	// See: [Receiver Official Documentation](https://prometheus.io/docs/alerting/latest/configuration/#inhibit_rule)
	//
	Receivers() *[]AlertManagerReceiver
	// Collection of alert manager templates that the configuration uses to format the alerts that it sends.
	Templates() *[]AlertManagerTemplate
	// Collection of timing configurations that can be used to control when specific alerts should be muted or activated.
	// See: [Time Interval Official Documentation](https://prometheus.io/docs/alerting/latest/configuration/#time_interval)
	//
	TimeIntervals() *[]TimeInterval
	// Adds a new inhibit rule that can mute alerts under specific circumstances.
	//
	// Returns: The inhibit rule that was added to the configuration.
	AddInhibitRule(id *string, options AlertManagerInhibitRuleProps) AlertManagerInhibitRule
	// Adds a new receiver to the configuration.
	//
	// The receiver can be used to
	// specify one or more destinations where alerts should be sent when matched
	// by a route in the routing tree.
	//
	// Returns: The receiver that was added to the configuration.
	AddReciever(id *string, options *AlertManagerReceiverProps) AlertManagerReceiver
	// Adds a new template to the configuration.
	//
	// The template can be referenced
	// within the configuration to control the formatting of the alerts being sent.
	//
	// Returns: The configuration object where the template was added.
	AddTemplate(template AlertManagerTemplate) IAlertManagerConfiguration
	// Adds a new time interval to the configuration.
	//
	// Time intervals can be used
	// to mute or activate groups of alerts under specific circumstances.
	//
	// Time intervals are referenced by routes in the routing tree to control the
	// behavior of the route. Time intervals to mute or activate alerts cannot be
	// added to the default route.
	//
	// Returns: The time interval that was added to the configuration.
	AddTimeInterval(id *string, options *TimeIntervalProps) TimeInterval
	// Associates the configuration with a construct that is handling the configuration of alert manager for an APS workspace.
	//
	// Returns: Alert manager configuration details.
	Bind(scope constructs.IConstruct) *AlertManagerConfigurationDetails
	// Generates a unique name for a template that is being referenced by the alert manager configuration.
	//
	// The names are only used to distinguish different templates as internally
	// each template is saved to a file with this name.
	//
	// Knowing the name is not needed to use the templates as template references
	// rely on the name given to the `define` keyword in the template content
	// itself rather than the name of the template as it is generated here.
	//
	// Returns: A unique name for the template.
	GenerateTemplateName(idx *float64) *string
	// Renders an alert manager configuration to be used for configuring notifications on an APS workspace.
	//
	// Returns: The string representation of the rendered Prometheus alert
	// manager configuration.
	RenderAlertManagerConfig(scope constructs.IConstruct) *string
	// Renders the configuration templates into the structure that is expected for an APS workspace alert manager configuration.
	//
	// Returns: An object that represents the templates contained in this
	// configuration or undefined if the configuration has no templates.
	RenderTemplates(_scope constructs.IConstruct) *map[string]*string
	// Returns a string representation of this construct.
	ToString() *string
}

// The jsii proxy struct for AlertManagerConfiguration
type jsiiProxy_AlertManagerConfiguration struct {
	internal.Type__constructsConstruct
	jsiiProxy_IAlertManagerConfiguration
}

func (j *jsiiProxy_AlertManagerConfiguration) DefaultReceiver() AlertManagerReceiver {
	var returns AlertManagerReceiver
	_jsii_.Get(
		j,
		"defaultReceiver",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertManagerConfiguration) DefaultRoute() AlertManagerRoute {
	var returns AlertManagerRoute
	_jsii_.Get(
		j,
		"defaultRoute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertManagerConfiguration) DefaultTopic() awssns.ITopic {
	var returns awssns.ITopic
	_jsii_.Get(
		j,
		"defaultTopic",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertManagerConfiguration) InhibitRules() *[]AlertManagerInhibitRule {
	var returns *[]AlertManagerInhibitRule
	_jsii_.Get(
		j,
		"inhibitRules",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertManagerConfiguration) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertManagerConfiguration) Receivers() *[]AlertManagerReceiver {
	var returns *[]AlertManagerReceiver
	_jsii_.Get(
		j,
		"receivers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertManagerConfiguration) Templates() *[]AlertManagerTemplate {
	var returns *[]AlertManagerTemplate
	_jsii_.Get(
		j,
		"templates",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertManagerConfiguration) TimeIntervals() *[]TimeInterval {
	var returns *[]TimeInterval
	_jsii_.Get(
		j,
		"timeIntervals",
		&returns,
	)
	return returns
}


// Creates a new instance of the AlertManagerConfiguration class.
func NewAlertManagerConfiguration(scope constructs.IConstruct, id *string, props *AlertManagerConfigurationProps) AlertManagerConfiguration {
	_init_.Initialize()

	if err := validateNewAlertManagerConfigurationParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_AlertManagerConfiguration{}

	_jsii_.Create(
		"cdk-extensions.aps.AlertManagerConfiguration",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Creates a new instance of the AlertManagerConfiguration class.
func NewAlertManagerConfiguration_Override(a AlertManagerConfiguration, scope constructs.IConstruct, id *string, props *AlertManagerConfigurationProps) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk-extensions.aps.AlertManagerConfiguration",
		[]interface{}{scope, id, props},
		a,
	)
}

// Imports an alert manager configuration from a raw string in the format expected by Amazon APS.
//
// The string should combine both the alert manager configuration and the
// alert template content into a single file. To import a configuration that
// has the alert manager configuration and templates being loaded from
// different files/sources use {@link fromSplitConfigurationContent} instead.
//
// Returns: An object that can be used to configure alert manager for APS.
// See: [Alert manager configuration specification](https://prometheus.io/docs/alerting/latest/configuration/)
//
func AlertManagerConfiguration_FromFullConfigurationContent(scope constructs.IConstruct, id *string, content *string) IAlertManagerConfiguration {
	_init_.Initialize()

	if err := validateAlertManagerConfiguration_FromFullConfigurationContentParameters(scope, id, content); err != nil {
		panic(err)
	}
	var returns IAlertManagerConfiguration

	_jsii_.StaticInvoke(
		"cdk-extensions.aps.AlertManagerConfiguration",
		"fromFullConfigurationContent",
		[]interface{}{scope, id, content},
		&returns,
	)

	return returns
}

// Imports an alert manager configuration using a file in the format expected by Amazon APS.
//
// The file should combine both the alert manager configuration and the alert
// template content into a single file. To import a configuration that has
// the alert manager configuration and templates being loaded from different
// files/sources use {@link fromSplitConfigurationFiles} instead.
//
// Returns: An object that can be used to configure alert manager for APS.
// See: [Alert manager configuration specification](https://prometheus.io/docs/alerting/latest/configuration/)
//
func AlertManagerConfiguration_FromFullConfigurationFile(scope constructs.IConstruct, id *string, path *string) IAlertManagerConfiguration {
	_init_.Initialize()

	if err := validateAlertManagerConfiguration_FromFullConfigurationFileParameters(scope, id, path); err != nil {
		panic(err)
	}
	var returns IAlertManagerConfiguration

	_jsii_.StaticInvoke(
		"cdk-extensions.aps.AlertManagerConfiguration",
		"fromFullConfigurationFile",
		[]interface{}{scope, id, path},
		&returns,
	)

	return returns
}

// Imports an alert manager configuration from a raw string.
//
// The string should be in YAML format and follow the alert manager
// configuration specifications. To import a configuration using a string
// that combines the alert manager configuration with the template content
// that is already in the format expected by Amazon APS use the import method
// {@link fromFullConfigurationContent} instead.
//
// If the configuration references any templates then each template the
// configuration references should be specified using the `templates`
// argument. Templates can be read from their own separate files on disk or
// as a string.
//
// Returns: An object that can be used to configure alert manager for APS.
// See: [Alert manager configuration specification](https://prometheus.io/docs/alerting/latest/configuration/)
//
func AlertManagerConfiguration_FromSplitConfigurationContent(scope constructs.IConstruct, id *string, content *string, templates *map[string]AlertManagerTemplate) IAlertManagerConfiguration {
	_init_.Initialize()

	if err := validateAlertManagerConfiguration_FromSplitConfigurationContentParameters(scope, id, content); err != nil {
		panic(err)
	}
	var returns IAlertManagerConfiguration

	_jsii_.StaticInvoke(
		"cdk-extensions.aps.AlertManagerConfiguration",
		"fromSplitConfigurationContent",
		[]interface{}{scope, id, content, templates},
		&returns,
	)

	return returns
}

// Imports an alert manager configuration from a file.
//
// The file is a YAML file that follows the alert manager configuration
// specifications. To import a configuration file that combines the alert
// manager configuration with the template content that is already in the
// format expected by Amazon APS use the import method
// {@link fromFullConfigurationFile} instead.
//
// If the configuration references any templates then each template the
// configuration references should be specified using the `templates`
// argument. Templates can be read from their own separate files on disk or
// as a string.
//
// Returns: An object that can be used to configure alert manager for APS.
// See: [Alert manager configuration specification](https://prometheus.io/docs/alerting/latest/configuration/)
//
func AlertManagerConfiguration_FromSplitConfigurationFiles(scope constructs.IConstruct, id *string, path *string, templates *map[string]AlertManagerTemplate) IAlertManagerConfiguration {
	_init_.Initialize()

	if err := validateAlertManagerConfiguration_FromSplitConfigurationFilesParameters(scope, id, path); err != nil {
		panic(err)
	}
	var returns IAlertManagerConfiguration

	_jsii_.StaticInvoke(
		"cdk-extensions.aps.AlertManagerConfiguration",
		"fromSplitConfigurationFiles",
		[]interface{}{scope, id, path, templates},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead.
func AlertManagerConfiguration_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAlertManagerConfiguration_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"cdk-extensions.aps.AlertManagerConfiguration",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlertManagerConfiguration) AddInhibitRule(id *string, options AlertManagerInhibitRuleProps) AlertManagerInhibitRule {
	if err := a.validateAddInhibitRuleParameters(id, options); err != nil {
		panic(err)
	}
	var returns AlertManagerInhibitRule

	_jsii_.Invoke(
		a,
		"addInhibitRule",
		[]interface{}{id, options},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlertManagerConfiguration) AddReciever(id *string, options *AlertManagerReceiverProps) AlertManagerReceiver {
	if err := a.validateAddRecieverParameters(id, options); err != nil {
		panic(err)
	}
	var returns AlertManagerReceiver

	_jsii_.Invoke(
		a,
		"addReciever",
		[]interface{}{id, options},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlertManagerConfiguration) AddTemplate(template AlertManagerTemplate) IAlertManagerConfiguration {
	if err := a.validateAddTemplateParameters(template); err != nil {
		panic(err)
	}
	var returns IAlertManagerConfiguration

	_jsii_.Invoke(
		a,
		"addTemplate",
		[]interface{}{template},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlertManagerConfiguration) AddTimeInterval(id *string, options *TimeIntervalProps) TimeInterval {
	if err := a.validateAddTimeIntervalParameters(id, options); err != nil {
		panic(err)
	}
	var returns TimeInterval

	_jsii_.Invoke(
		a,
		"addTimeInterval",
		[]interface{}{id, options},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlertManagerConfiguration) Bind(scope constructs.IConstruct) *AlertManagerConfigurationDetails {
	if err := a.validateBindParameters(scope); err != nil {
		panic(err)
	}
	var returns *AlertManagerConfigurationDetails

	_jsii_.Invoke(
		a,
		"bind",
		[]interface{}{scope},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlertManagerConfiguration) GenerateTemplateName(idx *float64) *string {
	if err := a.validateGenerateTemplateNameParameters(idx); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		a,
		"generateTemplateName",
		[]interface{}{idx},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlertManagerConfiguration) RenderAlertManagerConfig(scope constructs.IConstruct) *string {
	if err := a.validateRenderAlertManagerConfigParameters(scope); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		a,
		"renderAlertManagerConfig",
		[]interface{}{scope},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlertManagerConfiguration) RenderTemplates(_scope constructs.IConstruct) *map[string]*string {
	if err := a.validateRenderTemplatesParameters(_scope); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		a,
		"renderTemplates",
		[]interface{}{_scope},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlertManagerConfiguration) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

