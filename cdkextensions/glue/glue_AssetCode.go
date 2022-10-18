package glue

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/vibe-io/cdk-extensions-go/cdkextensions/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
	"github.com/aws/constructs-go/constructs/v10"
)

// Job Code from a local file.
type AssetCode interface {
	Code
	// Called when the Job is initialized to allow this object to bind.
	Bind(scope constructs.Construct, grantable awsiam.IGrantable) *CodeConfig
}

// The jsii proxy struct for AssetCode
type jsiiProxy_AssetCode struct {
	jsiiProxy_Code
}

func NewAssetCode(path *string, options *awscdk.AssetOptions) AssetCode {
	_init_.Initialize()

	if err := validateNewAssetCodeParameters(path, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_AssetCode{}

	_jsii_.Create(
		"cdk-extensions.glue.AssetCode",
		[]interface{}{path, options},
		&j,
	)

	return &j
}

func NewAssetCode_Override(a AssetCode, path *string, options *awscdk.AssetOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk-extensions.glue.AssetCode",
		[]interface{}{path, options},
		a,
	)
}

// Job code from a local disk path.
func AssetCode_FromAsset(path *string, options *awscdk.AssetOptions) AssetCode {
	_init_.Initialize()

	if err := validateAssetCode_FromAssetParameters(path, options); err != nil {
		panic(err)
	}
	var returns AssetCode

	_jsii_.StaticInvoke(
		"cdk-extensions.glue.AssetCode",
		"fromAsset",
		[]interface{}{path, options},
		&returns,
	)

	return returns
}

// Job code as an S3 object.
func AssetCode_FromBucket(bucket awss3.IBucket, key *string) S3Code {
	_init_.Initialize()

	if err := validateAssetCode_FromBucketParameters(bucket, key); err != nil {
		panic(err)
	}
	var returns S3Code

	_jsii_.StaticInvoke(
		"cdk-extensions.glue.AssetCode",
		"fromBucket",
		[]interface{}{bucket, key},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AssetCode) Bind(scope constructs.Construct, grantable awsiam.IGrantable) *CodeConfig {
	if err := a.validateBindParameters(scope, grantable); err != nil {
		panic(err)
	}
	var returns *CodeConfig

	_jsii_.Invoke(
		a,
		"bind",
		[]interface{}{scope, grantable},
		&returns,
	)

	return returns
}

