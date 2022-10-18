//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package kinesisfirehose

// Building without runtime type checking enabled, so all the below just return nil

func (l *jsiiProxy_LambdaPartitioningSource) validateBindParameters(scope constructs.IConstruct) error {
	return nil
}

func validateLambdaPartitioningSource_FromJsonParameters(options *JsonPartitioningOptions) error {
	return nil
}

func validateLambdaPartitioningSource_FromLambdaParameters(options *LambdaPartitioningOptions) error {
	return nil
}

func validateNewLambdaPartitioningSourceParameters(options *LambdaPartitioningOptions) error {
	return nil
}

