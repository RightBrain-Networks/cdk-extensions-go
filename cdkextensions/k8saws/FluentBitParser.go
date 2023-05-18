package k8saws

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/vibe-io/cdk-extensions-go/cdkextensions/jsii"
)

// Standard parse comfigurations which can be applied to Fluent Bit to allow for parsing data from incoming records.
//
// The records to which parsers are applied is controlled using the parser
// filter.
type FluentBitParser interface {
}

// The jsii proxy struct for FluentBitParser
type jsiiProxy_FluentBitParser struct {
	_ byte // padding
}

func NewFluentBitParser() FluentBitParser {
	_init_.Initialize()

	j := jsiiProxy_FluentBitParser{}

	_jsii_.Create(
		"cdk-extensions.k8s_aws.FluentBitParser",
		nil, // no parameters
		&j,
	)

	return &j
}

func NewFluentBitParser_Override(f FluentBitParser) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk-extensions.k8s_aws.FluentBitParser",
		nil, // no parameters
		f,
	)
}

// Creates a parser that processes records that are formatted in JSON.
//
// Returns: A parser object that can be applied to the Fluent Bit
// configuration.
func FluentBitParser_Json(name *string) IFluentBitParserPlugin {
	_init_.Initialize()

	if err := validateFluentBitParser_JsonParameters(name); err != nil {
		panic(err)
	}
	var returns IFluentBitParserPlugin

	_jsii_.StaticInvoke(
		"cdk-extensions.k8s_aws.FluentBitParser",
		"json",
		[]interface{}{name},
		&returns,
	)

	return returns
}

// Creates a parser that processes records that are formatted using the `logfmt` standard.
//
// Returns: A parser object that can be applied to the Fluent Bit
// configuration.
// See: [Golang logfmt documentation](https://pkg.go.dev/github.com/kr/logfmt)
//
func FluentBitParser_Logfmt(name *string) IFluentBitParserPlugin {
	_init_.Initialize()

	if err := validateFluentBitParser_LogfmtParameters(name); err != nil {
		panic(err)
	}
	var returns IFluentBitParserPlugin

	_jsii_.StaticInvoke(
		"cdk-extensions.k8s_aws.FluentBitParser",
		"logfmt",
		[]interface{}{name},
		&returns,
	)

	return returns
}

// Creates a parser that processes records that are formatted using the `ltsv` standard.
//
// Returns: A parser object that can be applied to the Fluent Bit
// configuration.
// See: [LTSV](http://ltsv.org/)
//
func FluentBitParser_Ltsv(name *string) IFluentBitParserPlugin {
	_init_.Initialize()

	if err := validateFluentBitParser_LtsvParameters(name); err != nil {
		panic(err)
	}
	var returns IFluentBitParserPlugin

	_jsii_.StaticInvoke(
		"cdk-extensions.k8s_aws.FluentBitParser",
		"ltsv",
		[]interface{}{name},
		&returns,
	)

	return returns
}

// Creates a parser that uses regular expressions to parse incoming records.
//
// Returns: A parser object that can be applied to the Fluent Bit
// configuration.
func FluentBitParser_Regex(name *string, regex *string) IFluentBitParserPlugin {
	_init_.Initialize()

	if err := validateFluentBitParser_RegexParameters(name, regex); err != nil {
		panic(err)
	}
	var returns IFluentBitParserPlugin

	_jsii_.StaticInvoke(
		"cdk-extensions.k8s_aws.FluentBitParser",
		"regex",
		[]interface{}{name, regex},
		&returns,
	)

	return returns
}

