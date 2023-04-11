//go:build !no_runtime_type_checking

package k8saws

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func validateExternalDnsRegistry_TxtParameters(options *TxtRegistryOptions) error {
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

