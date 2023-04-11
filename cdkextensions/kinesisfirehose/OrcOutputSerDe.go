package kinesisfirehose

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/vibe-io/cdk-extensions-go/cdkextensions/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2/awskinesisfirehose"
	"github.com/aws/constructs-go/constructs/v10"
)

type OrcOutputSerDe interface {
	OutputFormat
	BlockSizeBytes() *float64
	BloomFilterColumns() *[]*string
	BloomFilterFalsePositiveProbability() *float64
	Compression() OrcCompressionFormat
	DictionaryKeyThreshold() *float64
	EnablePadding() *bool
	FormatVersion() OrcFormatVersion
	PaddingTolerance() *float64
	RowIndexStride() *float64
	StripeSizeBytes() *float64
	AddBloomFilterColumn(column *string) OrcOutputSerDe
	Bind(_scope constructs.IConstruct) *awskinesisfirehose.CfnDeliveryStream_OutputFormatConfigurationProperty
}

// The jsii proxy struct for OrcOutputSerDe
type jsiiProxy_OrcOutputSerDe struct {
	jsiiProxy_OutputFormat
}

func (j *jsiiProxy_OrcOutputSerDe) BlockSizeBytes() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"blockSizeBytes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrcOutputSerDe) BloomFilterColumns() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"bloomFilterColumns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrcOutputSerDe) BloomFilterFalsePositiveProbability() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"bloomFilterFalsePositiveProbability",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrcOutputSerDe) Compression() OrcCompressionFormat {
	var returns OrcCompressionFormat
	_jsii_.Get(
		j,
		"compression",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrcOutputSerDe) DictionaryKeyThreshold() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"dictionaryKeyThreshold",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrcOutputSerDe) EnablePadding() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"enablePadding",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrcOutputSerDe) FormatVersion() OrcFormatVersion {
	var returns OrcFormatVersion
	_jsii_.Get(
		j,
		"formatVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrcOutputSerDe) PaddingTolerance() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"paddingTolerance",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrcOutputSerDe) RowIndexStride() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"rowIndexStride",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrcOutputSerDe) StripeSizeBytes() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"stripeSizeBytes",
		&returns,
	)
	return returns
}


func NewOrcOutputSerDe(options *OrcOutputSerDeOptions) OrcOutputSerDe {
	_init_.Initialize()

	if err := validateNewOrcOutputSerDeParameters(options); err != nil {
		panic(err)
	}
	j := jsiiProxy_OrcOutputSerDe{}

	_jsii_.Create(
		"cdk-extensions.kinesis_firehose.OrcOutputSerDe",
		[]interface{}{options},
		&j,
	)

	return &j
}

func NewOrcOutputSerDe_Override(o OrcOutputSerDe, options *OrcOutputSerDeOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk-extensions.kinesis_firehose.OrcOutputSerDe",
		[]interface{}{options},
		o,
	)
}

func OrcOutputSerDe_Orc(options *OrcOutputSerDeOptions) OrcOutputSerDe {
	_init_.Initialize()

	if err := validateOrcOutputSerDe_OrcParameters(options); err != nil {
		panic(err)
	}
	var returns OrcOutputSerDe

	_jsii_.StaticInvoke(
		"cdk-extensions.kinesis_firehose.OrcOutputSerDe",
		"orc",
		[]interface{}{options},
		&returns,
	)

	return returns
}

func OrcOutputSerDe_Parquet(options *ParquetOutputSerDeOptions) ParquetOutputSerDe {
	_init_.Initialize()

	if err := validateOrcOutputSerDe_ParquetParameters(options); err != nil {
		panic(err)
	}
	var returns ParquetOutputSerDe

	_jsii_.StaticInvoke(
		"cdk-extensions.kinesis_firehose.OrcOutputSerDe",
		"parquet",
		[]interface{}{options},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OrcOutputSerDe) AddBloomFilterColumn(column *string) OrcOutputSerDe {
	if err := o.validateAddBloomFilterColumnParameters(column); err != nil {
		panic(err)
	}
	var returns OrcOutputSerDe

	_jsii_.Invoke(
		o,
		"addBloomFilterColumn",
		[]interface{}{column},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OrcOutputSerDe) Bind(_scope constructs.IConstruct) *awskinesisfirehose.CfnDeliveryStream_OutputFormatConfigurationProperty {
	if err := o.validateBindParameters(_scope); err != nil {
		panic(err)
	}
	var returns *awskinesisfirehose.CfnDeliveryStream_OutputFormatConfigurationProperty

	_jsii_.Invoke(
		o,
		"bind",
		[]interface{}{_scope},
		&returns,
	)

	return returns
}

