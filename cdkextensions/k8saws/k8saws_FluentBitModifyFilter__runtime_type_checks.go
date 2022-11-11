//go:build !no_runtime_type_checking
// +build !no_runtime_type_checking

package k8saws

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/constructs-go/constructs/v10"
)

func (f *jsiiProxy_FluentBitModifyFilter) validateAddConditionParameters(condition ModifyCondition) error {
	if condition == nil {
		return fmt.Errorf("parameter condition is required, but nil was provided")
	}

	return nil
}

func (f *jsiiProxy_FluentBitModifyFilter) validateAddOperationParameters(operation ModifyOperation) error {
	if operation == nil {
		return fmt.Errorf("parameter operation is required, but nil was provided")
	}

	return nil
}

func (f *jsiiProxy_FluentBitModifyFilter) validateBindParameters(_scope constructs.IConstruct) error {
	if _scope == nil {
		return fmt.Errorf("parameter _scope is required, but nil was provided")
	}

	return nil
}

func (f *jsiiProxy_FluentBitModifyFilter) validateRenderConfigFileParameters(config *map[string]interface{}) error {
	if config == nil {
		return fmt.Errorf("parameter config is required, but nil was provided")
	}

	return nil
}

func validateNewFluentBitModifyFilterParameters(options *FluentBitModifyFilterOptions) error {
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

