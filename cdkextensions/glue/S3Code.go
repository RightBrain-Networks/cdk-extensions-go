package glue

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/vibe-io/cdk-extensions-go/cdkextensions/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
	"github.com/aws/constructs-go/constructs/v10"
)

// Glue job Code from an S3 bucket.
type S3Code interface {
	Code
	// Called when the Job is initialized to allow this object to bind.
	Bind(_scope constructs.Construct, grantable awsiam.IGrantable) *CodeConfig
}

// The jsii proxy struct for S3Code
type jsiiProxy_S3Code struct {
	jsiiProxy_Code
}

func NewS3Code(bucket awss3.IBucket, key *string) S3Code {
	_init_.Initialize()

	if err := validateNewS3CodeParameters(bucket, key); err != nil {
		panic(err)
	}
	j := jsiiProxy_S3Code{}

	_jsii_.Create(
		"cdk-extensions.glue.S3Code",
		[]interface{}{bucket, key},
		&j,
	)

	return &j
}

func NewS3Code_Override(s S3Code, bucket awss3.IBucket, key *string) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk-extensions.glue.S3Code",
		[]interface{}{bucket, key},
		s,
	)
}

// Job code from a local disk path.
func S3Code_FromAsset(path *string, options *awscdk.AssetOptions) AssetCode {
	_init_.Initialize()

	if err := validateS3Code_FromAssetParameters(path, options); err != nil {
		panic(err)
	}
	var returns AssetCode

	_jsii_.StaticInvoke(
		"cdk-extensions.glue.S3Code",
		"fromAsset",
		[]interface{}{path, options},
		&returns,
	)

	return returns
}

// Job code as an S3 object.
func S3Code_FromBucket(bucket awss3.IBucket, key *string) S3Code {
	_init_.Initialize()

	if err := validateS3Code_FromBucketParameters(bucket, key); err != nil {
		panic(err)
	}
	var returns S3Code

	_jsii_.StaticInvoke(
		"cdk-extensions.glue.S3Code",
		"fromBucket",
		[]interface{}{bucket, key},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_S3Code) Bind(_scope constructs.Construct, grantable awsiam.IGrantable) *CodeConfig {
	if err := s.validateBindParameters(_scope, grantable); err != nil {
		panic(err)
	}
	var returns *CodeConfig

	_jsii_.Invoke(
		s,
		"bind",
		[]interface{}{_scope, grantable},
		&returns,
	)

	return returns
}

