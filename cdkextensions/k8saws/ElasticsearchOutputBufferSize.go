package k8saws

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/vibe-io/cdk-extensions-go/cdkextensions/jsii"

	"github.com/vibe-io/cdk-extensions-go/cdkextensions/core"
)

// Represents the size of the Elasticsearch output buffer to be used by Fluent Bit.
type ElasticsearchOutputBufferSize interface {
	// The value to use for the Elasticsearch buffer output property.
	Value() *string
}

// The jsii proxy struct for ElasticsearchOutputBufferSize
type jsiiProxy_ElasticsearchOutputBufferSize struct {
	_ byte // padding
}

func (j *jsiiProxy_ElasticsearchOutputBufferSize) Value() *string {
	var returns *string
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


// Set the output buffer to a specified data size.
//
// Returns: An output buffer size object representing the specified buffer
// size.
func ElasticsearchOutputBufferSize_Bytes(size core.DataSize) ElasticsearchOutputBufferSize {
	_init_.Initialize()

	if err := validateElasticsearchOutputBufferSize_BytesParameters(size); err != nil {
		panic(err)
	}
	var returns ElasticsearchOutputBufferSize

	_jsii_.StaticInvoke(
		"cdk-extensions.k8s_aws.ElasticsearchOutputBufferSize",
		"bytes",
		[]interface{}{size},
		&returns,
	)

	return returns
}

// An escape hatch that allows an arbitrary value to be set for the Elasticsearch buffer output property.
//
// Returns: A `ElasticsearchOutputBufferSize` object representing the
// passed value.
func ElasticsearchOutputBufferSize_Of(value *string) ElasticsearchOutputBufferSize {
	_init_.Initialize()

	if err := validateElasticsearchOutputBufferSize_OfParameters(value); err != nil {
		panic(err)
	}
	var returns ElasticsearchOutputBufferSize

	_jsii_.StaticInvoke(
		"cdk-extensions.k8s_aws.ElasticsearchOutputBufferSize",
		"of",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func ElasticsearchOutputBufferSize_UNLIMITED() ElasticsearchOutputBufferSize {
	_init_.Initialize()
	var returns ElasticsearchOutputBufferSize
	_jsii_.StaticGet(
		"cdk-extensions.k8s_aws.ElasticsearchOutputBufferSize",
		"UNLIMITED",
		&returns,
	)
	return returns
}

