//go:build no_runtime_type_checking

package kinesisfirehose

// Building without runtime type checking enabled, so all the below just return nil

func (e *jsiiProxy_ExtendedS3Destination) validateAddProcessorParameters(processor DeliveryStreamProcessor) error {
	return nil
}

func (e *jsiiProxy_ExtendedS3Destination) validateBindParameters(scope constructs.IConstruct) error {
	return nil
}

func (e *jsiiProxy_ExtendedS3Destination) validateBuildConfigurationParameters(scope constructs.IConstruct) error {
	return nil
}

func (e *jsiiProxy_ExtendedS3Destination) validateRenderBackupConfigurationParameters(scope constructs.IConstruct) error {
	return nil
}

func (e *jsiiProxy_ExtendedS3Destination) validateRenderProcessorConfigurationParameters(scope constructs.IConstruct) error {
	return nil
}

func validateNewExtendedS3DestinationParameters(bucket awss3.IBucket, options *ExtendedS3DestinationOptions) error {
	return nil
}

