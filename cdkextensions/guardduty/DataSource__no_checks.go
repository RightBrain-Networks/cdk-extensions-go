//go:build no_runtime_type_checking

package guardduty

// Building without runtime type checking enabled, so all the below just return nil

func validateDataSource_KubernetesParameters(options *KubernetesOptions) error {
	return nil
}

func validateDataSource_MalwareProtectionParameters(options *MalwareProtectionOptions) error {
	return nil
}

func validateDataSource_S3LogsParameters(options *S3LogsOptions) error {
	return nil
}

