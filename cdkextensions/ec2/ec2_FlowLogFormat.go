package ec2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/vibe-io/cdk-extensions-go/cdkextensions/jsii"
)

type FlowLogFormat interface {
	// The fields that make up the flow log format, in the order that they should appear in the log entries.
	Fields() *[]FlowLogField
	// The rendered format string in the format expected by AWS when creating a new Flow Log.
	Template() *string
	// Adds a new field to the flow log output.
	//
	// New fields are added at the
	// end of a log entry after all the other fields that came before it.
	AddField(field FlowLogField)
}

// The jsii proxy struct for FlowLogFormat
type jsiiProxy_FlowLogFormat struct {
	_ byte // padding
}

func (j *jsiiProxy_FlowLogFormat) Fields() *[]FlowLogField {
	var returns *[]FlowLogField
	_jsii_.Get(
		j,
		"fields",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FlowLogFormat) Template() *string {
	var returns *string
	_jsii_.Get(
		j,
		"template",
		&returns,
	)
	return returns
}


// Creates a new instance of the FlowLogFormat class.
func NewFlowLogFormat(fields ...FlowLogField) FlowLogFormat {
	_init_.Initialize()

	args := []interface{}{}
	for _, a := range fields {
		args = append(args, a)
	}

	j := jsiiProxy_FlowLogFormat{}

	_jsii_.Create(
		"cdk-extensions.ec2.FlowLogFormat",
		args,
		&j,
	)

	return &j
}

// Creates a new instance of the FlowLogFormat class.
func NewFlowLogFormat_Override(f FlowLogFormat, fields ...FlowLogField) {
	_init_.Initialize()

	args := []interface{}{}
	for _, a := range fields {
		args = append(args, a)
	}

	_jsii_.Create(
		"cdk-extensions.ec2.FlowLogFormat",
		args,
		f,
	)
}

func FlowLogFormat_V2() FlowLogFormat {
	_init_.Initialize()
	var returns FlowLogFormat
	_jsii_.StaticGet(
		"cdk-extensions.ec2.FlowLogFormat",
		"V2",
		&returns,
	)
	return returns
}

func FlowLogFormat_V3() FlowLogFormat {
	_init_.Initialize()
	var returns FlowLogFormat
	_jsii_.StaticGet(
		"cdk-extensions.ec2.FlowLogFormat",
		"V3",
		&returns,
	)
	return returns
}

func FlowLogFormat_V4() FlowLogFormat {
	_init_.Initialize()
	var returns FlowLogFormat
	_jsii_.StaticGet(
		"cdk-extensions.ec2.FlowLogFormat",
		"V4",
		&returns,
	)
	return returns
}

func FlowLogFormat_V5() FlowLogFormat {
	_init_.Initialize()
	var returns FlowLogFormat
	_jsii_.StaticGet(
		"cdk-extensions.ec2.FlowLogFormat",
		"V5",
		&returns,
	)
	return returns
}

func (f *jsiiProxy_FlowLogFormat) AddField(field FlowLogField) {
	if err := f.validateAddFieldParameters(field); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"addField",
		[]interface{}{field},
	)
}

