package k8saws

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/vibe-io/cdk-extensions-go/cdkextensions/jsii"

	"github.com/vibe-io/cdk-extensions-go/cdkextensions/core"
)

// Represents the size of the OpenSeach output buffer to be used by Fluent Bit.
type OpenSearchOutputBufferSize interface {
	// The value to use for the OpenSearch buffer output property.
	Value() *string
}

// The jsii proxy struct for OpenSearchOutputBufferSize
type jsiiProxy_OpenSearchOutputBufferSize struct {
	_ byte // padding
}

func (j *jsiiProxy_OpenSearchOutputBufferSize) Value() *string {
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
func OpenSearchOutputBufferSize_Bytes(size core.DataSize) OpenSearchOutputBufferSize {
	_init_.Initialize()

	if err := validateOpenSearchOutputBufferSize_BytesParameters(size); err != nil {
		panic(err)
	}
	var returns OpenSearchOutputBufferSize

	_jsii_.StaticInvoke(
		"cdk-extensions.k8s_aws.OpenSearchOutputBufferSize",
		"bytes",
		[]interface{}{size},
		&returns,
	)

	return returns
}

// An escape hatch that allows an arbitrary value to be set for the OpenSearch buffer output property.
//
// Returns: A `OpenSearchOutputBufferSize` object representing the passed
// value.
func OpenSearchOutputBufferSize_Of(value *string) OpenSearchOutputBufferSize {
	_init_.Initialize()

	if err := validateOpenSearchOutputBufferSize_OfParameters(value); err != nil {
		panic(err)
	}
	var returns OpenSearchOutputBufferSize

	_jsii_.StaticInvoke(
		"cdk-extensions.k8s_aws.OpenSearchOutputBufferSize",
		"of",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func OpenSearchOutputBufferSize_UNLIMITED() OpenSearchOutputBufferSize {
	_init_.Initialize()
	var returns OpenSearchOutputBufferSize
	_jsii_.StaticGet(
		"cdk-extensions.k8s_aws.OpenSearchOutputBufferSize",
		"UNLIMITED",
		&returns,
	)
	return returns
}

