package kinesisfirehose

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/vibe-io/cdk-extensions-go/cdkextensions/jsii"
)

type TableVersion interface {
	Version() *string
}

// The jsii proxy struct for TableVersion
type jsiiProxy_TableVersion struct {
	_ byte // padding
}

func (j *jsiiProxy_TableVersion) Version() *string {
	var returns *string
	_jsii_.Get(
		j,
		"version",
		&returns,
	)
	return returns
}


func TableVersion_Fixed(version *float64) TableVersion {
	_init_.Initialize()

	if err := validateTableVersion_FixedParameters(version); err != nil {
		panic(err)
	}
	var returns TableVersion

	_jsii_.StaticInvoke(
		"cdk-extensions.kinesis_firehose.TableVersion",
		"fixed",
		[]interface{}{version},
		&returns,
	)

	return returns
}

func TableVersion_LATEST() TableVersion {
	_init_.Initialize()
	var returns TableVersion
	_jsii_.StaticGet(
		"cdk-extensions.kinesis_firehose.TableVersion",
		"LATEST",
		&returns,
	)
	return returns
}

