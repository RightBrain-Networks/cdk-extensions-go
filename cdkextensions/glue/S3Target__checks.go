//go:build !no_runtime_type_checking

package glue

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
)

func (s *jsiiProxy_S3Target) validateAddExclusionParameters(exclusion *string) error {
	if exclusion == nil {
		return fmt.Errorf("parameter exclusion is required, but nil was provided")
	}

	return nil
}

func (s *jsiiProxy_S3Target) validateBindParameters(crawler Crawler) error {
	if crawler == nil {
		return fmt.Errorf("parameter crawler is required, but nil was provided")
	}

	return nil
}

func validateNewS3TargetParameters(bucket awss3.IBucket, options *S3TargetOptions) error {
	if bucket == nil {
		return fmt.Errorf("parameter bucket is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

