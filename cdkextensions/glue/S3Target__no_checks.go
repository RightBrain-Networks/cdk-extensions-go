//go:build no_runtime_type_checking

package glue

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_S3Target) validateAddExclusionParameters(exclusion *string) error {
	return nil
}

func (s *jsiiProxy_S3Target) validateBindParameters(crawler Crawler) error {
	return nil
}

func validateNewS3TargetParameters(bucket awss3.IBucket, options *S3TargetOptions) error {
	return nil
}

