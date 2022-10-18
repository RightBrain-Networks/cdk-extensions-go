package kinesisfirehose

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/vibe-io/cdk-extensions-go/cdkextensions/jsii"
)

type MetaDataExtractionQuery interface {
	Query() *string
	SetQuery(val *string)
	Render() *string
}

// The jsii proxy struct for MetaDataExtractionQuery
type jsiiProxy_MetaDataExtractionQuery struct {
	_ byte // padding
}

func (j *jsiiProxy_MetaDataExtractionQuery) Query() *string {
	var returns *string
	_jsii_.Get(
		j,
		"query",
		&returns,
	)
	return returns
}


func NewMetaDataExtractionQuery(query *string) MetaDataExtractionQuery {
	_init_.Initialize()

	if err := validateNewMetaDataExtractionQueryParameters(query); err != nil {
		panic(err)
	}
	j := jsiiProxy_MetaDataExtractionQuery{}

	_jsii_.Create(
		"cdk-extensions.kinesis_firehose.MetaDataExtractionQuery",
		[]interface{}{query},
		&j,
	)

	return &j
}

func NewMetaDataExtractionQuery_Override(m MetaDataExtractionQuery, query *string) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk-extensions.kinesis_firehose.MetaDataExtractionQuery",
		[]interface{}{query},
		m,
	)
}

func (j *jsiiProxy_MetaDataExtractionQuery)SetQuery(val *string) {
	if err := j.validateSetQueryParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"query",
		val,
	)
}

func MetaDataExtractionQuery_Jq(fields *map[string]*string) JsonQuery {
	_init_.Initialize()

	if err := validateMetaDataExtractionQuery_JqParameters(fields); err != nil {
		panic(err)
	}
	var returns JsonQuery

	_jsii_.StaticInvoke(
		"cdk-extensions.kinesis_firehose.MetaDataExtractionQuery",
		"jq",
		[]interface{}{fields},
		&returns,
	)

	return returns
}

func MetaDataExtractionQuery_Of(query *string) MetaDataExtractionQuery {
	_init_.Initialize()

	if err := validateMetaDataExtractionQuery_OfParameters(query); err != nil {
		panic(err)
	}
	var returns MetaDataExtractionQuery

	_jsii_.StaticInvoke(
		"cdk-extensions.kinesis_firehose.MetaDataExtractionQuery",
		"of",
		[]interface{}{query},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MetaDataExtractionQuery) Render() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"render",
		nil, // no parameters
		&returns,
	)

	return returns
}

