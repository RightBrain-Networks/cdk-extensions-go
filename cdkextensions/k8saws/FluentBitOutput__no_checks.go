//go:build no_runtime_type_checking

package k8saws

// Building without runtime type checking enabled, so all the below just return nil

func validateFluentBitOutput_CloudwatchLogsParameters(match FluentBitMatch, logGroup awslogs.ILogGroup) error {
	return nil
}

func validateFluentBitOutput_KinesisParameters(match FluentBitMatch, stream awskinesis.IStream) error {
	return nil
}

func validateFluentBitOutput_KinesisFirehoseParameters(match FluentBitMatch, deliveryStream kinesisfirehose.IDeliveryStream) error {
	return nil
}

func validateFluentBitOutput_OpensearchParameters(match FluentBitMatch, domain awsopensearchservice.IDomain) error {
	return nil
}

