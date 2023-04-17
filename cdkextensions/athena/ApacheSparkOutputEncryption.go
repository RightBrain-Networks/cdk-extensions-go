package athena

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/vibe-io/cdk-extensions-go/cdkextensions/jsii"
)

type ApacheSparkOutputEncryption interface {
}

// The jsii proxy struct for ApacheSparkOutputEncryption
type jsiiProxy_ApacheSparkOutputEncryption struct {
	_ byte // padding
}

func NewApacheSparkOutputEncryption() ApacheSparkOutputEncryption {
	_init_.Initialize()

	j := jsiiProxy_ApacheSparkOutputEncryption{}

	_jsii_.Create(
		"cdk-extensions.athena.ApacheSparkOutputEncryption",
		nil, // no parameters
		&j,
	)

	return &j
}

func NewApacheSparkOutputEncryption_Override(a ApacheSparkOutputEncryption) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk-extensions.athena.ApacheSparkOutputEncryption",
		nil, // no parameters
		a,
	)
}

func ApacheSparkOutputEncryption_SseKms(options *AthenaResultKmsEncryptionOptions) IAthenaResultEncryption {
	_init_.Initialize()

	if err := validateApacheSparkOutputEncryption_SseKmsParameters(options); err != nil {
		panic(err)
	}
	var returns IAthenaResultEncryption

	_jsii_.StaticInvoke(
		"cdk-extensions.athena.ApacheSparkOutputEncryption",
		"sseKms",
		[]interface{}{options},
		&returns,
	)

	return returns
}

func ApacheSparkOutputEncryption_SseS3() IAthenaResultEncryption {
	_init_.Initialize()

	var returns IAthenaResultEncryption

	_jsii_.StaticInvoke(
		"cdk-extensions.athena.ApacheSparkOutputEncryption",
		"sseS3",
		nil, // no parameters
		&returns,
	)

	return returns
}

