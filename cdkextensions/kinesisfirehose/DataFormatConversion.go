package kinesisfirehose

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/vibe-io/cdk-extensions-go/cdkextensions/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awskinesisfirehose"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/vibe-io/cdk-extensions-go/cdkextensions/glue"
)

type DataFormatConversion interface {
	CatalogId() *string
	Database() glue.Database
	Enabled() *bool
	InputFormat() InputFormat
	OutputFormat() OutputFormat
	Region() *string
	Role() awsiam.IRole
	Table() glue.Table
	Version() TableVersion
	Bind(scope constructs.IConstruct) *awskinesisfirehose.CfnDeliveryStream_DataFormatConversionConfigurationProperty
}

// The jsii proxy struct for DataFormatConversion
type jsiiProxy_DataFormatConversion struct {
	_ byte // padding
}

func (j *jsiiProxy_DataFormatConversion) CatalogId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"catalogId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFormatConversion) Database() glue.Database {
	var returns glue.Database
	_jsii_.Get(
		j,
		"database",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFormatConversion) Enabled() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFormatConversion) InputFormat() InputFormat {
	var returns InputFormat
	_jsii_.Get(
		j,
		"inputFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFormatConversion) OutputFormat() OutputFormat {
	var returns OutputFormat
	_jsii_.Get(
		j,
		"outputFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFormatConversion) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFormatConversion) Role() awsiam.IRole {
	var returns awsiam.IRole
	_jsii_.Get(
		j,
		"role",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFormatConversion) Table() glue.Table {
	var returns glue.Table
	_jsii_.Get(
		j,
		"table",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFormatConversion) Version() TableVersion {
	var returns TableVersion
	_jsii_.Get(
		j,
		"version",
		&returns,
	)
	return returns
}


func NewDataFormatConversion(options *DataFormatConversionOptions) DataFormatConversion {
	_init_.Initialize()

	if err := validateNewDataFormatConversionParameters(options); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataFormatConversion{}

	_jsii_.Create(
		"cdk-extensions.kinesis_firehose.DataFormatConversion",
		[]interface{}{options},
		&j,
	)

	return &j
}

func NewDataFormatConversion_Override(d DataFormatConversion, options *DataFormatConversionOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk-extensions.kinesis_firehose.DataFormatConversion",
		[]interface{}{options},
		d,
	)
}

func (d *jsiiProxy_DataFormatConversion) Bind(scope constructs.IConstruct) *awskinesisfirehose.CfnDeliveryStream_DataFormatConversionConfigurationProperty {
	if err := d.validateBindParameters(scope); err != nil {
		panic(err)
	}
	var returns *awskinesisfirehose.CfnDeliveryStream_DataFormatConversionConfigurationProperty

	_jsii_.Invoke(
		d,
		"bind",
		[]interface{}{scope},
		&returns,
	)

	return returns
}

