//go:build no_runtime_type_checking

package guardduty

// Building without runtime type checking enabled, so all the below just return nil

func validateFeature_EbsMalwareProtectionParameters(options *FeatureOptions) error {
	return nil
}

func validateFeature_EksAuditLogsParameters(options *FeatureOptions) error {
	return nil
}

func validateFeature_EksRuntimeMonitoringParameters(options *EksRuntimeMonitoringOptions) error {
	return nil
}

func validateFeature_LambdaNetworkLogsParameters(options *FeatureOptions) error {
	return nil
}

func validateFeature_RdsLoginEventsParameters(options *FeatureOptions) error {
	return nil
}

func validateFeature_S3DataEventsParameters(options *FeatureOptions) error {
	return nil
}

