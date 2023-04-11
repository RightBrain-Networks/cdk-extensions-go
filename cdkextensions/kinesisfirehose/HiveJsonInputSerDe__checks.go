//go:build !no_runtime_type_checking

package kinesisfirehose

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/constructs-go/constructs/v10"
)

func (h *jsiiProxy_HiveJsonInputSerDe) validateAddTimestampFormatParameters(format *string) error {
	if format == nil {
		return fmt.Errorf("parameter format is required, but nil was provided")
	}

	return nil
}

func (h *jsiiProxy_HiveJsonInputSerDe) validateBindParameters(_scope constructs.IConstruct) error {
	if _scope == nil {
		return fmt.Errorf("parameter _scope is required, but nil was provided")
	}

	return nil
}

func validateHiveJsonInputSerDe_HiveJsonParameters(options *HiveJsonInputSerDeOptions) error {
	if options == nil {
		return fmt.Errorf("parameter options is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func validateHiveJsonInputSerDe_OpenxJsonParameters(options *OpenxJsonInputSerDeOptions) error {
	if options == nil {
		return fmt.Errorf("parameter options is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func validateNewHiveJsonInputSerDeParameters(options *HiveJsonInputSerDeOptions) error {
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

