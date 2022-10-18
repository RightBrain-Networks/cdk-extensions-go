//go:build !no_runtime_type_checking
// +build !no_runtime_type_checking

package kinesisfirehose

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/constructs-go/constructs/v10"
)

func (o *jsiiProxy_OpenxJsonInputSerDe) validateAddColumnKeyMappingParameters(columnName *string, jsonKey *string) error {
	if columnName == nil {
		return fmt.Errorf("parameter columnName is required, but nil was provided")
	}

	if jsonKey == nil {
		return fmt.Errorf("parameter jsonKey is required, but nil was provided")
	}

	return nil
}

func (o *jsiiProxy_OpenxJsonInputSerDe) validateBindParameters(_scope constructs.IConstruct) error {
	if _scope == nil {
		return fmt.Errorf("parameter _scope is required, but nil was provided")
	}

	return nil
}

func validateOpenxJsonInputSerDe_HiveJsonParameters(options *HiveJsonInputSerDeOptions) error {
	if options == nil {
		return fmt.Errorf("parameter options is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func validateOpenxJsonInputSerDe_OpenxJsonParameters(options *OpenxJsonInputSerDeOptions) error {
	if options == nil {
		return fmt.Errorf("parameter options is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func validateNewOpenxJsonInputSerDeParameters(options *OpenxJsonInputSerDeOptions) error {
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

