package k8saws

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/vibe-io/cdk-extensions-go/cdkextensions/jsii"
)

// Define a buffering mechanism for the new records created by the rewrite tag Fluent Bit filter plugin.
type EmitterStorageType interface {
	// The name of the emitter storage type as it should appear in the plugin configuration file.
	Name() *string
}

// The jsii proxy struct for EmitterStorageType
type jsiiProxy_EmitterStorageType struct {
	_ byte // padding
}

func (j *jsiiProxy_EmitterStorageType) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}


// An escape hatch that allows for specifying a custom value for the rewrite tag plugin's `Emitter_Storage.type` field.
func EmitterStorageType_Of(name *string) EmitterStorageType {
	_init_.Initialize()

	if err := validateEmitterStorageType_OfParameters(name); err != nil {
		panic(err)
	}
	var returns EmitterStorageType

	_jsii_.StaticInvoke(
		"cdk-extensions.k8s_aws.EmitterStorageType",
		"of",
		[]interface{}{name},
		&returns,
	)

	return returns
}

func EmitterStorageType_FILESYSTEM() EmitterStorageType {
	_init_.Initialize()
	var returns EmitterStorageType
	_jsii_.StaticGet(
		"cdk-extensions.k8s_aws.EmitterStorageType",
		"FILESYSTEM",
		&returns,
	)
	return returns
}

func EmitterStorageType_MEMORY() EmitterStorageType {
	_init_.Initialize()
	var returns EmitterStorageType
	_jsii_.StaticGet(
		"cdk-extensions.k8s_aws.EmitterStorageType",
		"MEMORY",
		&returns,
	)
	return returns
}

