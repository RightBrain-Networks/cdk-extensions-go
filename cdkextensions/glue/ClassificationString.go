package glue

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/vibe-io/cdk-extensions-go/cdkextensions/jsii"
)

// Classification string given to tables with this data format.
// See: https://docs.aws.amazon.com/glue/latest/dg/add-classifier.html#classifier-built-in
//
type ClassificationString interface {
	Value() *string
}

// The jsii proxy struct for ClassificationString
type jsiiProxy_ClassificationString struct {
	_ byte // padding
}

func (j *jsiiProxy_ClassificationString) Value() *string {
	var returns *string
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


func NewClassificationString(value *string) ClassificationString {
	_init_.Initialize()

	if err := validateNewClassificationStringParameters(value); err != nil {
		panic(err)
	}
	j := jsiiProxy_ClassificationString{}

	_jsii_.Create(
		"cdk-extensions.glue.ClassificationString",
		[]interface{}{value},
		&j,
	)

	return &j
}

func NewClassificationString_Override(c ClassificationString, value *string) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk-extensions.glue.ClassificationString",
		[]interface{}{value},
		c,
	)
}

func ClassificationString_AVRO() ClassificationString {
	_init_.Initialize()
	var returns ClassificationString
	_jsii_.StaticGet(
		"cdk-extensions.glue.ClassificationString",
		"AVRO",
		&returns,
	)
	return returns
}

func ClassificationString_CSV() ClassificationString {
	_init_.Initialize()
	var returns ClassificationString
	_jsii_.StaticGet(
		"cdk-extensions.glue.ClassificationString",
		"CSV",
		&returns,
	)
	return returns
}

func ClassificationString_JSON() ClassificationString {
	_init_.Initialize()
	var returns ClassificationString
	_jsii_.StaticGet(
		"cdk-extensions.glue.ClassificationString",
		"JSON",
		&returns,
	)
	return returns
}

func ClassificationString_ORC() ClassificationString {
	_init_.Initialize()
	var returns ClassificationString
	_jsii_.StaticGet(
		"cdk-extensions.glue.ClassificationString",
		"ORC",
		&returns,
	)
	return returns
}

func ClassificationString_PARQUET() ClassificationString {
	_init_.Initialize()
	var returns ClassificationString
	_jsii_.StaticGet(
		"cdk-extensions.glue.ClassificationString",
		"PARQUET",
		&returns,
	)
	return returns
}

func ClassificationString_XML() ClassificationString {
	_init_.Initialize()
	var returns ClassificationString
	_jsii_.StaticGet(
		"cdk-extensions.glue.ClassificationString",
		"XML",
		&returns,
	)
	return returns
}

