//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package kinesisfirehose

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_S3Destination) validateBindParameters(scope constructs.IConstruct) error {
	return nil
}

func (s *jsiiProxy_S3Destination) validateBuildConfigurationParameters(scope constructs.IConstruct) error {
	return nil
}

func (s *jsiiProxy_S3Destination) validateRenderBackupConfigurationParameters(scope constructs.IConstruct) error {
	return nil
}

func validateNewS3DestinationParameters(bucket awss3.IBucket, options *S3DestinationOptions) error {
	return nil
}

