//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package kinesisfirehose

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DynamicPartitioning) validateBindParameters(_scope constructs.IConstruct) error {
	return nil
}

func validateDynamicPartitioning_FromJsonParameters(options *JsonPartitioningOptions) error {
	return nil
}

func validateDynamicPartitioning_FromLambdaParameters(options *LambdaPartitioningOptions) error {
	return nil
}

func validateNewDynamicPartitioningParameters(options *CommonPartitioningOptions) error {
	return nil
}

