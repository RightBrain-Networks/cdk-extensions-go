package kinesisfirehose

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/vibe-io/cdk-extensions-go/cdkextensions/jsii"
)

type SubRecordType interface {
	Name() *string
}

// The jsii proxy struct for SubRecordType
type jsiiProxy_SubRecordType struct {
	_ byte // padding
}

func (j *jsiiProxy_SubRecordType) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}


func SubRecordType_Of(name *string) SubRecordType {
	_init_.Initialize()

	if err := validateSubRecordType_OfParameters(name); err != nil {
		panic(err)
	}
	var returns SubRecordType

	_jsii_.StaticInvoke(
		"cdk-extensions.kinesis_firehose.SubRecordType",
		"of",
		[]interface{}{name},
		&returns,
	)

	return returns
}

func SubRecordType_DELIMITED() SubRecordType {
	_init_.Initialize()
	var returns SubRecordType
	_jsii_.StaticGet(
		"cdk-extensions.kinesis_firehose.SubRecordType",
		"DELIMITED",
		&returns,
	)
	return returns
}

func SubRecordType_JSON() SubRecordType {
	_init_.Initialize()
	var returns SubRecordType
	_jsii_.StaticGet(
		"cdk-extensions.kinesis_firehose.SubRecordType",
		"JSON",
		&returns,
	)
	return returns
}

