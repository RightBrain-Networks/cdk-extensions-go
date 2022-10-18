package ec2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/vibe-io/cdk-extensions-go/cdkextensions/jsii"
)

type FlowLogField interface {
	// The name of the Flow Log field, as it should be used when building a format string.
	Name() *string
	// The data type of the field as it would appear in Parquet.
	//
	// For
	// information on the type for various files, see documentation on the
	// [available fields](https://docs.aws.amazon.com/vpc/latest/userguide/flow-logs.html#flow-logs-fields).
	Type() FlowLogDataType
}

// The jsii proxy struct for FlowLogField
type jsiiProxy_FlowLogField struct {
	_ byte // padding
}

func (j *jsiiProxy_FlowLogField) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FlowLogField) Type() FlowLogDataType {
	var returns FlowLogDataType
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}


// Creates a new instance of the FlowLogField class.
func NewFlowLogField(name *string, type_ FlowLogDataType) FlowLogField {
	_init_.Initialize()

	if err := validateNewFlowLogFieldParameters(name, type_); err != nil {
		panic(err)
	}
	j := jsiiProxy_FlowLogField{}

	_jsii_.Create(
		"cdk-extensions.ec2.FlowLogField",
		[]interface{}{name, type_},
		&j,
	)

	return &j
}

// Creates a new instance of the FlowLogField class.
func NewFlowLogField_Override(f FlowLogField, name *string, type_ FlowLogDataType) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk-extensions.ec2.FlowLogField",
		[]interface{}{name, type_},
		f,
	)
}

func FlowLogField_ACCOUNT_ID() FlowLogField {
	_init_.Initialize()
	var returns FlowLogField
	_jsii_.StaticGet(
		"cdk-extensions.ec2.FlowLogField",
		"ACCOUNT_ID",
		&returns,
	)
	return returns
}

func FlowLogField_ACTION() FlowLogField {
	_init_.Initialize()
	var returns FlowLogField
	_jsii_.StaticGet(
		"cdk-extensions.ec2.FlowLogField",
		"ACTION",
		&returns,
	)
	return returns
}

func FlowLogField_AZ_ID() FlowLogField {
	_init_.Initialize()
	var returns FlowLogField
	_jsii_.StaticGet(
		"cdk-extensions.ec2.FlowLogField",
		"AZ_ID",
		&returns,
	)
	return returns
}

func FlowLogField_BYTES() FlowLogField {
	_init_.Initialize()
	var returns FlowLogField
	_jsii_.StaticGet(
		"cdk-extensions.ec2.FlowLogField",
		"BYTES",
		&returns,
	)
	return returns
}

func FlowLogField_DSTADDR() FlowLogField {
	_init_.Initialize()
	var returns FlowLogField
	_jsii_.StaticGet(
		"cdk-extensions.ec2.FlowLogField",
		"DSTADDR",
		&returns,
	)
	return returns
}

func FlowLogField_DSTPORT() FlowLogField {
	_init_.Initialize()
	var returns FlowLogField
	_jsii_.StaticGet(
		"cdk-extensions.ec2.FlowLogField",
		"DSTPORT",
		&returns,
	)
	return returns
}

func FlowLogField_END() FlowLogField {
	_init_.Initialize()
	var returns FlowLogField
	_jsii_.StaticGet(
		"cdk-extensions.ec2.FlowLogField",
		"END",
		&returns,
	)
	return returns
}

func FlowLogField_FLOW_DIRECTION() FlowLogField {
	_init_.Initialize()
	var returns FlowLogField
	_jsii_.StaticGet(
		"cdk-extensions.ec2.FlowLogField",
		"FLOW_DIRECTION",
		&returns,
	)
	return returns
}

func FlowLogField_INSTANCE_ID() FlowLogField {
	_init_.Initialize()
	var returns FlowLogField
	_jsii_.StaticGet(
		"cdk-extensions.ec2.FlowLogField",
		"INSTANCE_ID",
		&returns,
	)
	return returns
}

func FlowLogField_INTERFACE_ID() FlowLogField {
	_init_.Initialize()
	var returns FlowLogField
	_jsii_.StaticGet(
		"cdk-extensions.ec2.FlowLogField",
		"INTERFACE_ID",
		&returns,
	)
	return returns
}

func FlowLogField_LOG_STATUS() FlowLogField {
	_init_.Initialize()
	var returns FlowLogField
	_jsii_.StaticGet(
		"cdk-extensions.ec2.FlowLogField",
		"LOG_STATUS",
		&returns,
	)
	return returns
}

func FlowLogField_PACKETS() FlowLogField {
	_init_.Initialize()
	var returns FlowLogField
	_jsii_.StaticGet(
		"cdk-extensions.ec2.FlowLogField",
		"PACKETS",
		&returns,
	)
	return returns
}

func FlowLogField_PKT_DST_AWS_SERVICE() FlowLogField {
	_init_.Initialize()
	var returns FlowLogField
	_jsii_.StaticGet(
		"cdk-extensions.ec2.FlowLogField",
		"PKT_DST_AWS_SERVICE",
		&returns,
	)
	return returns
}

func FlowLogField_PKT_DSTADDR() FlowLogField {
	_init_.Initialize()
	var returns FlowLogField
	_jsii_.StaticGet(
		"cdk-extensions.ec2.FlowLogField",
		"PKT_DSTADDR",
		&returns,
	)
	return returns
}

func FlowLogField_PKT_SRC_AWS_SERVICE() FlowLogField {
	_init_.Initialize()
	var returns FlowLogField
	_jsii_.StaticGet(
		"cdk-extensions.ec2.FlowLogField",
		"PKT_SRC_AWS_SERVICE",
		&returns,
	)
	return returns
}

func FlowLogField_PKT_SRCADDR() FlowLogField {
	_init_.Initialize()
	var returns FlowLogField
	_jsii_.StaticGet(
		"cdk-extensions.ec2.FlowLogField",
		"PKT_SRCADDR",
		&returns,
	)
	return returns
}

func FlowLogField_PROTOCOL() FlowLogField {
	_init_.Initialize()
	var returns FlowLogField
	_jsii_.StaticGet(
		"cdk-extensions.ec2.FlowLogField",
		"PROTOCOL",
		&returns,
	)
	return returns
}

func FlowLogField_REGION() FlowLogField {
	_init_.Initialize()
	var returns FlowLogField
	_jsii_.StaticGet(
		"cdk-extensions.ec2.FlowLogField",
		"REGION",
		&returns,
	)
	return returns
}

func FlowLogField_SRCADDR() FlowLogField {
	_init_.Initialize()
	var returns FlowLogField
	_jsii_.StaticGet(
		"cdk-extensions.ec2.FlowLogField",
		"SRCADDR",
		&returns,
	)
	return returns
}

func FlowLogField_SRCPORT() FlowLogField {
	_init_.Initialize()
	var returns FlowLogField
	_jsii_.StaticGet(
		"cdk-extensions.ec2.FlowLogField",
		"SRCPORT",
		&returns,
	)
	return returns
}

func FlowLogField_START() FlowLogField {
	_init_.Initialize()
	var returns FlowLogField
	_jsii_.StaticGet(
		"cdk-extensions.ec2.FlowLogField",
		"START",
		&returns,
	)
	return returns
}

func FlowLogField_SUBLOCATION_ID() FlowLogField {
	_init_.Initialize()
	var returns FlowLogField
	_jsii_.StaticGet(
		"cdk-extensions.ec2.FlowLogField",
		"SUBLOCATION_ID",
		&returns,
	)
	return returns
}

func FlowLogField_SUBLOCATION_TYPE() FlowLogField {
	_init_.Initialize()
	var returns FlowLogField
	_jsii_.StaticGet(
		"cdk-extensions.ec2.FlowLogField",
		"SUBLOCATION_TYPE",
		&returns,
	)
	return returns
}

func FlowLogField_SUBNET_ID() FlowLogField {
	_init_.Initialize()
	var returns FlowLogField
	_jsii_.StaticGet(
		"cdk-extensions.ec2.FlowLogField",
		"SUBNET_ID",
		&returns,
	)
	return returns
}

func FlowLogField_TCP_FLAGS() FlowLogField {
	_init_.Initialize()
	var returns FlowLogField
	_jsii_.StaticGet(
		"cdk-extensions.ec2.FlowLogField",
		"TCP_FLAGS",
		&returns,
	)
	return returns
}

func FlowLogField_TRAFFIC_PATH() FlowLogField {
	_init_.Initialize()
	var returns FlowLogField
	_jsii_.StaticGet(
		"cdk-extensions.ec2.FlowLogField",
		"TRAFFIC_PATH",
		&returns,
	)
	return returns
}

func FlowLogField_TYPE() FlowLogField {
	_init_.Initialize()
	var returns FlowLogField
	_jsii_.StaticGet(
		"cdk-extensions.ec2.FlowLogField",
		"TYPE",
		&returns,
	)
	return returns
}

func FlowLogField_VERSION() FlowLogField {
	_init_.Initialize()
	var returns FlowLogField
	_jsii_.StaticGet(
		"cdk-extensions.ec2.FlowLogField",
		"VERSION",
		&returns,
	)
	return returns
}

func FlowLogField_VPC_ID() FlowLogField {
	_init_.Initialize()
	var returns FlowLogField
	_jsii_.StaticGet(
		"cdk-extensions.ec2.FlowLogField",
		"VPC_ID",
		&returns,
	)
	return returns
}

