//go:build !no_runtime_type_checking

package guardduty

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func validateDataSource_KubernetesParameters(options *KubernetesOptions) error {
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func validateDataSource_MalwareProtectionParameters(options *MalwareProtectionOptions) error {
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func validateDataSource_S3LogsParameters(options *S3LogsOptions) error {
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

