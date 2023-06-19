//go:build !no_runtime_type_checking

package resourcegroups

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func validateGroupConfiguration_CloudFormationStackParameters(reference IStackReference, props *CloudFormationStackProps) error {
	if reference == nil {
		return fmt.Errorf("parameter reference is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

