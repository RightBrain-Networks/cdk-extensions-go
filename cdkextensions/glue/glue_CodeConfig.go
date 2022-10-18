package glue

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
)

// Result of binding `Code` into a `Job`.
type CodeConfig struct {
	// The location of the code in S3.
	S3Location *awss3.Location `field:"required" json:"s3Location" yaml:"s3Location"`
}

