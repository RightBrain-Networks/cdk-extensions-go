package k8saws

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/vibe-io/cdk-extensions-go/cdkextensions/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Standard filter options which can be applied to Fluent Bit to control the output and formatting of logs.
//
// Filters change the structure of log records by doing things like adding
// metadata to a record, restructuring a record, or adding and removing fields.
type FluentBitFilter interface {
}

// The jsii proxy struct for FluentBitFilter
type jsiiProxy_FluentBitFilter struct {
	_ byte // padding
}

func NewFluentBitFilter() FluentBitFilter {
	_init_.Initialize()

	j := jsiiProxy_FluentBitFilter{}

	_jsii_.Create(
		"cdk-extensions.k8s_aws.FluentBitFilter",
		nil, // no parameters
		&j,
	)

	return &j
}

func NewFluentBitFilter_Override(f FluentBitFilter) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk-extensions.k8s_aws.FluentBitFilter",
		nil, // no parameters
		f,
	)
}

// Creates a filter that adds fields to a record that matches the given pattern.
//
// Returns: A filter object that can be applied to the Fluent Bit
// configuration.
func FluentBitFilter_AppendFields(match FluentBitMatch, records ...*AppendedRecord) IFluentBitFilterPlugin {
	_init_.Initialize()

	if err := validateFluentBitFilter_AppendFieldsParameters(match, &records); err != nil {
		panic(err)
	}
	args := []interface{}{match}
	for _, a := range records {
		args = append(args, a)
	}

	var returns IFluentBitFilterPlugin

	_jsii_.StaticInvoke(
		"cdk-extensions.k8s_aws.FluentBitFilter",
		"appendFields",
		args,
		&returns,
	)

	return returns
}

// Creates a filter that removes a set of fields from any records that match a given pattern.
//
// Returns: A filter object that can be applied to the Fluent Bit
// configuration.
func FluentBitFilter_BlacklistFields(match FluentBitMatch, fields ...*string) IFluentBitFilterPlugin {
	_init_.Initialize()

	if err := validateFluentBitFilter_BlacklistFieldsParameters(match); err != nil {
		panic(err)
	}
	args := []interface{}{match}
	for _, a := range fields {
		args = append(args, a)
	}

	var returns IFluentBitFilterPlugin

	_jsii_.StaticInvoke(
		"cdk-extensions.k8s_aws.FluentBitFilter",
		"blacklistFields",
		args,
		&returns,
	)

	return returns
}

// Filters log entries based on a pattern.
//
// Log entries can be removed and
// not forwarded based on whether they match or do not match the given
// pattern.
//
// Returns: A filter object that can be applied to the Fluent Bit
// configuration.
func FluentBitFilter_Grep(match FluentBitMatch, pattern *FluentBitGrepRegex) IFluentBitFilterPlugin {
	_init_.Initialize()

	if err := validateFluentBitFilter_GrepParameters(match, pattern); err != nil {
		panic(err)
	}
	var returns IFluentBitFilterPlugin

	_jsii_.StaticInvoke(
		"cdk-extensions.k8s_aws.FluentBitFilter",
		"grep",
		[]interface{}{match, pattern},
		&returns,
	)

	return returns
}

// Adds Kubernetes metadata to output records including pod information, labels, etc..
//
// Returns: A filter object that can be applied to the Fluent Bit
// configuration.
func FluentBitFilter_Kubernetes(match FluentBitMatch) IFluentBitFilterPlugin {
	_init_.Initialize()

	if err := validateFluentBitFilter_KubernetesParameters(match); err != nil {
		panic(err)
	}
	var returns IFluentBitFilterPlugin

	_jsii_.StaticInvoke(
		"cdk-extensions.k8s_aws.FluentBitFilter",
		"kubernetes",
		[]interface{}{match},
		&returns,
	)

	return returns
}

// Lifts nested fields in a record up to their parent object.
//
// Returns: A filter object that can be applied to the Fluent Bit
// configuration.
func FluentBitFilter_Lift(match FluentBitMatch, nestedUnder *string) IFluentBitFilterPlugin {
	_init_.Initialize()

	if err := validateFluentBitFilter_LiftParameters(match, nestedUnder); err != nil {
		panic(err)
	}
	var returns IFluentBitFilterPlugin

	_jsii_.StaticInvoke(
		"cdk-extensions.k8s_aws.FluentBitFilter",
		"lift",
		[]interface{}{match, nestedUnder},
		&returns,
	)

	return returns
}

// Applies various transformations to matched records including adding, removing, copying, and renaming fields.
//
// Returns: A filter object that can be applied to the Fluent Bit
// configuration.
func FluentBitFilter_Modify(match FluentBitMatch, operations ...ModifyOperation) IFluentBitFilterPlugin {
	_init_.Initialize()

	if err := validateFluentBitFilter_ModifyParameters(match); err != nil {
		panic(err)
	}
	args := []interface{}{match}
	for _, a := range operations {
		args = append(args, a)
	}

	var returns IFluentBitFilterPlugin

	_jsii_.StaticInvoke(
		"cdk-extensions.k8s_aws.FluentBitFilter",
		"modify",
		args,
		&returns,
	)

	return returns
}

// Nests a set of fields in a record under into a specified object.
//
// Returns: A filter object that can be applied to the Fluent Bit
// configuration.
func FluentBitFilter_Nest(match FluentBitMatch, nestUnder *string, fields ...*string) IFluentBitFilterPlugin {
	_init_.Initialize()

	if err := validateFluentBitFilter_NestParameters(match, nestUnder); err != nil {
		panic(err)
	}
	args := []interface{}{match, nestUnder}
	for _, a := range fields {
		args = append(args, a)
	}

	var returns IFluentBitFilterPlugin

	_jsii_.StaticInvoke(
		"cdk-extensions.k8s_aws.FluentBitFilter",
		"nest",
		args,
		&returns,
	)

	return returns
}

// Applies a set of parsers to matched records.
//
// The parser is used to read the input record and set structured fields in
// the output.
//
// Returns: A filter object that can be applied to the Fluent Bit
// configuration.
func FluentBitFilter_Parser(match FluentBitMatch, key *string, parsers ...IFluentBitParserPlugin) IFluentBitFilterPlugin {
	_init_.Initialize()

	if err := validateFluentBitFilter_ParserParameters(match, key); err != nil {
		panic(err)
	}
	args := []interface{}{match, key}
	for _, a := range parsers {
		args = append(args, a)
	}

	var returns IFluentBitFilterPlugin

	_jsii_.StaticInvoke(
		"cdk-extensions.k8s_aws.FluentBitFilter",
		"parser",
		args,
		&returns,
	)

	return returns
}

// Allows modification of tags set by the input configuration to affect the routing of when records are output.
//
// Returns: A filter object that can be applied to the Fluent Bit
// configuration.
func FluentBitFilter_RewriteTag(match FluentBitMatch, rules ...*RewriteTagRule) IFluentBitFilterPlugin {
	_init_.Initialize()

	if err := validateFluentBitFilter_RewriteTagParameters(match, &rules); err != nil {
		panic(err)
	}
	args := []interface{}{match}
	for _, a := range rules {
		args = append(args, a)
	}

	var returns IFluentBitFilterPlugin

	_jsii_.StaticInvoke(
		"cdk-extensions.k8s_aws.FluentBitFilter",
		"rewriteTag",
		args,
		&returns,
	)

	return returns
}

// Sets an average rate of messages that are allowed to be output over a configured period of time.
//
// When the rate of messages surpasses the configured limits messages will
// be dropped.
//
// Returns: A filter object that can be applied to the Fluent Bit
// configuration.
func FluentBitFilter_Throttle(match FluentBitMatch, interval awscdk.Duration, rate *float64, window *float64) IFluentBitFilterPlugin {
	_init_.Initialize()

	if err := validateFluentBitFilter_ThrottleParameters(match, interval, rate, window); err != nil {
		panic(err)
	}
	var returns IFluentBitFilterPlugin

	_jsii_.StaticInvoke(
		"cdk-extensions.k8s_aws.FluentBitFilter",
		"throttle",
		[]interface{}{match, interval, rate, window},
		&returns,
	)

	return returns
}

// Creates a filter that removes all fields in a record that are not approved.
//
// Returns: A filter object that can be applied to the Fluent Bit
// configuration.
func FluentBitFilter_WhitelistFields(match FluentBitMatch, fields ...*string) IFluentBitFilterPlugin {
	_init_.Initialize()

	if err := validateFluentBitFilter_WhitelistFieldsParameters(match); err != nil {
		panic(err)
	}
	args := []interface{}{match}
	for _, a := range fields {
		args = append(args, a)
	}

	var returns IFluentBitFilterPlugin

	_jsii_.StaticInvoke(
		"cdk-extensions.k8s_aws.FluentBitFilter",
		"whitelistFields",
		args,
		&returns,
	)

	return returns
}

