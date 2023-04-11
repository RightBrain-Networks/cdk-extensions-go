package kinesisfirehose

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/vibe-io/cdk-extensions-go/cdkextensions/jsii"
)

type JsonQuery interface {
	MetaDataExtractionQuery
	Query() *string
	SetQuery(val *string)
	AddField(name *string, query *string) JsonQuery
	Render() *string
}

// The jsii proxy struct for JsonQuery
type jsiiProxy_JsonQuery struct {
	jsiiProxy_MetaDataExtractionQuery
}

func (j *jsiiProxy_JsonQuery) Query() *string {
	var returns *string
	_jsii_.Get(
		j,
		"query",
		&returns,
	)
	return returns
}


func NewJsonQuery(fields *map[string]*string) JsonQuery {
	_init_.Initialize()

	j := jsiiProxy_JsonQuery{}

	_jsii_.Create(
		"cdk-extensions.kinesis_firehose.JsonQuery",
		[]interface{}{fields},
		&j,
	)

	return &j
}

func NewJsonQuery_Override(j JsonQuery, fields *map[string]*string) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk-extensions.kinesis_firehose.JsonQuery",
		[]interface{}{fields},
		j,
	)
}

func (j *jsiiProxy_JsonQuery)SetQuery(val *string) {
	if err := j.validateSetQueryParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"query",
		val,
	)
}

func JsonQuery_Jq(fields *map[string]*string) JsonQuery {
	_init_.Initialize()

	if err := validateJsonQuery_JqParameters(fields); err != nil {
		panic(err)
	}
	var returns JsonQuery

	_jsii_.StaticInvoke(
		"cdk-extensions.kinesis_firehose.JsonQuery",
		"jq",
		[]interface{}{fields},
		&returns,
	)

	return returns
}

func JsonQuery_Of(query *string) MetaDataExtractionQuery {
	_init_.Initialize()

	if err := validateJsonQuery_OfParameters(query); err != nil {
		panic(err)
	}
	var returns MetaDataExtractionQuery

	_jsii_.StaticInvoke(
		"cdk-extensions.kinesis_firehose.JsonQuery",
		"of",
		[]interface{}{query},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JsonQuery) AddField(name *string, query *string) JsonQuery {
	if err := j.validateAddFieldParameters(name, query); err != nil {
		panic(err)
	}
	var returns JsonQuery

	_jsii_.Invoke(
		j,
		"addField",
		[]interface{}{name, query},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JsonQuery) Render() *string {
	var returns *string

	_jsii_.Invoke(
		j,
		"render",
		nil, // no parameters
		&returns,
	)

	return returns
}

