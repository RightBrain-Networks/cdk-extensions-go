package glue

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/vibe-io/cdk-extensions-go/cdkextensions/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
)

type S3Target interface {
	ICrawlerTarget
	Bucket() awss3.IBucket
	Connection() Connection
	Exclusions() *[]*string
	KeyPrefix() *string
	SampleSize() *string
	AddExclusion(exclusion *string)
	Bind(crawler Crawler) *CrawlerTargetCollection
}

// The jsii proxy struct for S3Target
type jsiiProxy_S3Target struct {
	jsiiProxy_ICrawlerTarget
}

func (j *jsiiProxy_S3Target) Bucket() awss3.IBucket {
	var returns awss3.IBucket
	_jsii_.Get(
		j,
		"bucket",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3Target) Connection() Connection {
	var returns Connection
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3Target) Exclusions() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"exclusions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3Target) KeyPrefix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyPrefix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3Target) SampleSize() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sampleSize",
		&returns,
	)
	return returns
}


func NewS3Target(bucket awss3.IBucket, options *S3TargetOptions) S3Target {
	_init_.Initialize()

	if err := validateNewS3TargetParameters(bucket, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_S3Target{}

	_jsii_.Create(
		"cdk-extensions.glue.S3Target",
		[]interface{}{bucket, options},
		&j,
	)

	return &j
}

func NewS3Target_Override(s S3Target, bucket awss3.IBucket, options *S3TargetOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk-extensions.glue.S3Target",
		[]interface{}{bucket, options},
		s,
	)
}

func (s *jsiiProxy_S3Target) AddExclusion(exclusion *string) {
	if err := s.validateAddExclusionParameters(exclusion); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"addExclusion",
		[]interface{}{exclusion},
	)
}

func (s *jsiiProxy_S3Target) Bind(crawler Crawler) *CrawlerTargetCollection {
	if err := s.validateBindParameters(crawler); err != nil {
		panic(err)
	}
	var returns *CrawlerTargetCollection

	_jsii_.Invoke(
		s,
		"bind",
		[]interface{}{crawler},
		&returns,
	)

	return returns
}

