package athena

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/vibe-io/cdk-extensions-go/cdkextensions/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awskms"
)

type ApacheSparkEngineOptions interface {
	EncryptionKey() awskms.IKey
	EngineVersion() ApacheSparkEngineVersion
	Output() *AnalyticsEngineOutputOptions
	PublishMetrics() *bool
	Role() awsiam.IRole
}

// The jsii proxy struct for ApacheSparkEngineOptions
type jsiiProxy_ApacheSparkEngineOptions struct {
	_ byte // padding
}

func (j *jsiiProxy_ApacheSparkEngineOptions) EncryptionKey() awskms.IKey {
	var returns awskms.IKey
	_jsii_.Get(
		j,
		"encryptionKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApacheSparkEngineOptions) EngineVersion() ApacheSparkEngineVersion {
	var returns ApacheSparkEngineVersion
	_jsii_.Get(
		j,
		"engineVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApacheSparkEngineOptions) Output() *AnalyticsEngineOutputOptions {
	var returns *AnalyticsEngineOutputOptions
	_jsii_.Get(
		j,
		"output",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApacheSparkEngineOptions) PublishMetrics() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"publishMetrics",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApacheSparkEngineOptions) Role() awsiam.IRole {
	var returns awsiam.IRole
	_jsii_.Get(
		j,
		"role",
		&returns,
	)
	return returns
}


func NewApacheSparkEngineOptions() ApacheSparkEngineOptions {
	_init_.Initialize()

	j := jsiiProxy_ApacheSparkEngineOptions{}

	_jsii_.Create(
		"cdk-extensions.athena.ApacheSparkEngineOptions",
		nil, // no parameters
		&j,
	)

	return &j
}

func NewApacheSparkEngineOptions_Override(a ApacheSparkEngineOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk-extensions.athena.ApacheSparkEngineOptions",
		nil, // no parameters
		a,
	)
}

