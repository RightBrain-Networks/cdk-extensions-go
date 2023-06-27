//go:build !no_runtime_type_checking

package guardduty

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func validateFeature_EbsMalwareProtectionParameters(options *FeatureOptions) error {
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func validateFeature_EksAuditLogsParameters(options *FeatureOptions) error {
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func validateFeature_EksRuntimeMonitoringParameters(options *EksRuntimeMonitoringOptions) error {
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func validateFeature_LambdaNetworkLogsParameters(options *FeatureOptions) error {
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func validateFeature_RdsLoginEventsParameters(options *FeatureOptions) error {
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func validateFeature_S3DataEventsParameters(options *FeatureOptions) error {
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

