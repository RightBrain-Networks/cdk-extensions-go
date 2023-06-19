//go:build !no_runtime_type_checking

package resourcegroups

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/constructs-go/constructs/v10"
)

func (c *jsiiProxy_CloudFormationStack) validateAddResourceTypeParameters(typeId *string) error {
	if typeId == nil {
		return fmt.Errorf("parameter typeId is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_CloudFormationStack) validateBindParameters(_scope constructs.IConstruct) error {
	if _scope == nil {
		return fmt.Errorf("parameter _scope is required, but nil was provided")
	}

	return nil
}

func validateNewCloudFormationStackParameters(reference IStackReference, props *CloudFormationStackProps) error {
	if reference == nil {
		return fmt.Errorf("parameter reference is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

