//go:build !no_runtime_type_checking
// +build !no_runtime_type_checking

package k8saws

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/constructs-go/constructs/v10"
)

func (f *jsiiProxy_FluentBitRecordModifierFilter) validateAddAllowParameters(tag *string) error {
	if tag == nil {
		return fmt.Errorf("parameter tag is required, but nil was provided")
	}

	return nil
}

func (f *jsiiProxy_FluentBitRecordModifierFilter) validateAddRecordParameters(record *AppendedRecord) error {
	if record == nil {
		return fmt.Errorf("parameter record is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(record, func() string { return "parameter record" }); err != nil {
		return err
	}

	return nil
}

func (f *jsiiProxy_FluentBitRecordModifierFilter) validateAddRemoveParameters(tag *string) error {
	if tag == nil {
		return fmt.Errorf("parameter tag is required, but nil was provided")
	}

	return nil
}

func (f *jsiiProxy_FluentBitRecordModifierFilter) validateBindParameters(_scope constructs.IConstruct) error {
	if _scope == nil {
		return fmt.Errorf("parameter _scope is required, but nil was provided")
	}

	return nil
}

func (f *jsiiProxy_FluentBitRecordModifierFilter) validateRenderConfigFileParameters(config *map[string]interface{}) error {
	if config == nil {
		return fmt.Errorf("parameter config is required, but nil was provided")
	}

	return nil
}

func validateNewFluentBitRecordModifierFilterParameters(options *FluentBitRecordModifierFilterOptions) error {
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

