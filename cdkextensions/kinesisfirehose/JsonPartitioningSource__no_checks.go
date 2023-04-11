//go:build no_runtime_type_checking

package kinesisfirehose

// Building without runtime type checking enabled, so all the below just return nil

func (j *jsiiProxy_JsonPartitioningSource) validateAddPartitionParameters(name *string, query *string) error {
	return nil
}

func (j *jsiiProxy_JsonPartitioningSource) validateBindParameters(scope constructs.IConstruct) error {
	return nil
}

func validateJsonPartitioningSource_FromJsonParameters(options *JsonPartitioningOptions) error {
	return nil
}

func validateJsonPartitioningSource_FromLambdaParameters(options *LambdaPartitioningOptions) error {
	return nil
}

func validateNewJsonPartitioningSourceParameters(options *JsonPartitioningOptions) error {
	return nil
}

