package kinesisfirehose

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/vibe-io/cdk-extensions-go/cdkextensions/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2/awskinesisfirehose"
	"github.com/aws/constructs-go/constructs/v10"
)

type ParquetOutputSerDe interface {
	OutputFormat
	BlockSizeBytes() *float64
	Compression() ParquetCompressionFormat
	EnableDictionaryCompression() *bool
	MaxPaddingBytes() *float64
	PageSizeBytes() *float64
	WriterVersion() ParquetWriterVersion
	Bind(_scope constructs.IConstruct) *awskinesisfirehose.CfnDeliveryStream_OutputFormatConfigurationProperty
}

// The jsii proxy struct for ParquetOutputSerDe
type jsiiProxy_ParquetOutputSerDe struct {
	jsiiProxy_OutputFormat
}

func (j *jsiiProxy_ParquetOutputSerDe) BlockSizeBytes() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"blockSizeBytes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ParquetOutputSerDe) Compression() ParquetCompressionFormat {
	var returns ParquetCompressionFormat
	_jsii_.Get(
		j,
		"compression",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ParquetOutputSerDe) EnableDictionaryCompression() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"enableDictionaryCompression",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ParquetOutputSerDe) MaxPaddingBytes() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxPaddingBytes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ParquetOutputSerDe) PageSizeBytes() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"pageSizeBytes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ParquetOutputSerDe) WriterVersion() ParquetWriterVersion {
	var returns ParquetWriterVersion
	_jsii_.Get(
		j,
		"writerVersion",
		&returns,
	)
	return returns
}


func NewParquetOutputSerDe(options *ParquetOutputSerDeOptions) ParquetOutputSerDe {
	_init_.Initialize()

	if err := validateNewParquetOutputSerDeParameters(options); err != nil {
		panic(err)
	}
	j := jsiiProxy_ParquetOutputSerDe{}

	_jsii_.Create(
		"cdk-extensions.kinesis_firehose.ParquetOutputSerDe",
		[]interface{}{options},
		&j,
	)

	return &j
}

func NewParquetOutputSerDe_Override(p ParquetOutputSerDe, options *ParquetOutputSerDeOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk-extensions.kinesis_firehose.ParquetOutputSerDe",
		[]interface{}{options},
		p,
	)
}

func ParquetOutputSerDe_Orc(options *OrcOutputSerDeOptions) OrcOutputSerDe {
	_init_.Initialize()

	if err := validateParquetOutputSerDe_OrcParameters(options); err != nil {
		panic(err)
	}
	var returns OrcOutputSerDe

	_jsii_.StaticInvoke(
		"cdk-extensions.kinesis_firehose.ParquetOutputSerDe",
		"orc",
		[]interface{}{options},
		&returns,
	)

	return returns
}

func ParquetOutputSerDe_Parquet(options *ParquetOutputSerDeOptions) ParquetOutputSerDe {
	_init_.Initialize()

	if err := validateParquetOutputSerDe_ParquetParameters(options); err != nil {
		panic(err)
	}
	var returns ParquetOutputSerDe

	_jsii_.StaticInvoke(
		"cdk-extensions.kinesis_firehose.ParquetOutputSerDe",
		"parquet",
		[]interface{}{options},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_ParquetOutputSerDe) Bind(_scope constructs.IConstruct) *awskinesisfirehose.CfnDeliveryStream_OutputFormatConfigurationProperty {
	if err := p.validateBindParameters(_scope); err != nil {
		panic(err)
	}
	var returns *awskinesisfirehose.CfnDeliveryStream_OutputFormatConfigurationProperty

	_jsii_.Invoke(
		p,
		"bind",
		[]interface{}{_scope},
		&returns,
	)

	return returns
}

