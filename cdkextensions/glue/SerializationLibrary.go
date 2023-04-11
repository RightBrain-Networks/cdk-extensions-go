package glue

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/vibe-io/cdk-extensions-go/cdkextensions/jsii"
)

// Serialization library to use when serializing/deserializing (SerDe) table records.
// See: https://cwiki.apache.org/confluence/display/Hive/SerDe
//
type SerializationLibrary interface {
	ClassName() *string
}

// The jsii proxy struct for SerializationLibrary
type jsiiProxy_SerializationLibrary struct {
	_ byte // padding
}

func (j *jsiiProxy_SerializationLibrary) ClassName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"className",
		&returns,
	)
	return returns
}


func NewSerializationLibrary(className *string) SerializationLibrary {
	_init_.Initialize()

	if err := validateNewSerializationLibraryParameters(className); err != nil {
		panic(err)
	}
	j := jsiiProxy_SerializationLibrary{}

	_jsii_.Create(
		"cdk-extensions.glue.SerializationLibrary",
		[]interface{}{className},
		&j,
	)

	return &j
}

func NewSerializationLibrary_Override(s SerializationLibrary, className *string) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk-extensions.glue.SerializationLibrary",
		[]interface{}{className},
		s,
	)
}

func SerializationLibrary_AVRO() SerializationLibrary {
	_init_.Initialize()
	var returns SerializationLibrary
	_jsii_.StaticGet(
		"cdk-extensions.glue.SerializationLibrary",
		"AVRO",
		&returns,
	)
	return returns
}

func SerializationLibrary_CLOUDTRAIL() SerializationLibrary {
	_init_.Initialize()
	var returns SerializationLibrary
	_jsii_.StaticGet(
		"cdk-extensions.glue.SerializationLibrary",
		"CLOUDTRAIL",
		&returns,
	)
	return returns
}

func SerializationLibrary_GROK() SerializationLibrary {
	_init_.Initialize()
	var returns SerializationLibrary
	_jsii_.StaticGet(
		"cdk-extensions.glue.SerializationLibrary",
		"GROK",
		&returns,
	)
	return returns
}

func SerializationLibrary_HIVE_JSON() SerializationLibrary {
	_init_.Initialize()
	var returns SerializationLibrary
	_jsii_.StaticGet(
		"cdk-extensions.glue.SerializationLibrary",
		"HIVE_JSON",
		&returns,
	)
	return returns
}

func SerializationLibrary_LAZY_SIMPLE() SerializationLibrary {
	_init_.Initialize()
	var returns SerializationLibrary
	_jsii_.StaticGet(
		"cdk-extensions.glue.SerializationLibrary",
		"LAZY_SIMPLE",
		&returns,
	)
	return returns
}

func SerializationLibrary_OPEN_CSV() SerializationLibrary {
	_init_.Initialize()
	var returns SerializationLibrary
	_jsii_.StaticGet(
		"cdk-extensions.glue.SerializationLibrary",
		"OPEN_CSV",
		&returns,
	)
	return returns
}

func SerializationLibrary_OPENX_JSON() SerializationLibrary {
	_init_.Initialize()
	var returns SerializationLibrary
	_jsii_.StaticGet(
		"cdk-extensions.glue.SerializationLibrary",
		"OPENX_JSON",
		&returns,
	)
	return returns
}

func SerializationLibrary_ORC() SerializationLibrary {
	_init_.Initialize()
	var returns SerializationLibrary
	_jsii_.StaticGet(
		"cdk-extensions.glue.SerializationLibrary",
		"ORC",
		&returns,
	)
	return returns
}

func SerializationLibrary_PARQUET() SerializationLibrary {
	_init_.Initialize()
	var returns SerializationLibrary
	_jsii_.StaticGet(
		"cdk-extensions.glue.SerializationLibrary",
		"PARQUET",
		&returns,
	)
	return returns
}

func SerializationLibrary_REGEXP() SerializationLibrary {
	_init_.Initialize()
	var returns SerializationLibrary
	_jsii_.StaticGet(
		"cdk-extensions.glue.SerializationLibrary",
		"REGEXP",
		&returns,
	)
	return returns
}

