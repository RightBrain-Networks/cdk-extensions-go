//go:build no_runtime_type_checking

package glue

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_AssetCode) validateBindParameters(scope constructs.Construct, grantable awsiam.IGrantable) error {
	return nil
}

func validateAssetCode_FromAssetParameters(path *string, options *awscdk.AssetOptions) error {
	return nil
}

func validateAssetCode_FromBucketParameters(bucket awss3.IBucket, key *string) error {
	return nil
}

func validateNewAssetCodeParameters(path *string, options *awscdk.AssetOptions) error {
	return nil
}

