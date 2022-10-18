//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package kinesisfirehose

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CloudWatchLoggingConfiguration) validateBindParameters(scope constructs.IConstruct) error {
	return nil
}

func validateNewCloudWatchLoggingConfigurationParameters(options *CloudWatchLoggingConfigurationOptions) error {
	return nil
}

