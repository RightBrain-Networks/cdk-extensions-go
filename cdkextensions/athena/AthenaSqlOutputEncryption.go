package athena

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/vibe-io/cdk-extensions-go/cdkextensions/jsii"
)

type AthenaSqlOutputEncryption interface {
}

// The jsii proxy struct for AthenaSqlOutputEncryption
type jsiiProxy_AthenaSqlOutputEncryption struct {
	_ byte // padding
}

func NewAthenaSqlOutputEncryption() AthenaSqlOutputEncryption {
	_init_.Initialize()

	j := jsiiProxy_AthenaSqlOutputEncryption{}

	_jsii_.Create(
		"cdk-extensions.athena.AthenaSqlOutputEncryption",
		nil, // no parameters
		&j,
	)

	return &j
}

func NewAthenaSqlOutputEncryption_Override(a AthenaSqlOutputEncryption) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk-extensions.athena.AthenaSqlOutputEncryption",
		nil, // no parameters
		a,
	)
}

func AthenaSqlOutputEncryption_CseKms(options *AthenaResultKmsEncryptionOptions) IAthenaResultEncryption {
	_init_.Initialize()

	if err := validateAthenaSqlOutputEncryption_CseKmsParameters(options); err != nil {
		panic(err)
	}
	var returns IAthenaResultEncryption

	_jsii_.StaticInvoke(
		"cdk-extensions.athena.AthenaSqlOutputEncryption",
		"cseKms",
		[]interface{}{options},
		&returns,
	)

	return returns
}

func AthenaSqlOutputEncryption_SseKms(options *AthenaResultKmsEncryptionOptions) IAthenaResultEncryption {
	_init_.Initialize()

	if err := validateAthenaSqlOutputEncryption_SseKmsParameters(options); err != nil {
		panic(err)
	}
	var returns IAthenaResultEncryption

	_jsii_.StaticInvoke(
		"cdk-extensions.athena.AthenaSqlOutputEncryption",
		"sseKms",
		[]interface{}{options},
		&returns,
	)

	return returns
}

func AthenaSqlOutputEncryption_SseS3() IAthenaResultEncryption {
	_init_.Initialize()

	var returns IAthenaResultEncryption

	_jsii_.StaticInvoke(
		"cdk-extensions.athena.AthenaSqlOutputEncryption",
		"sseS3",
		nil, // no parameters
		&returns,
	)

	return returns
}

