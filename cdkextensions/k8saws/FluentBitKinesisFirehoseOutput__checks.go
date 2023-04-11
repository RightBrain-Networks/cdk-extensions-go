//go:build !no_runtime_type_checking

package k8saws

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/constructs-go/constructs/v10"
)

func (f *jsiiProxy_FluentBitKinesisFirehoseOutput) validateBindParameters(scope constructs.IConstruct) error {
	if scope == nil {
		return fmt.Errorf("parameter scope is required, but nil was provided")
	}

	return nil
}

func (f *jsiiProxy_FluentBitKinesisFirehoseOutput) validateRenderConfigFileParameters(config *map[string]interface{}) error {
	if config == nil {
		return fmt.Errorf("parameter config is required, but nil was provided")
	}

	return nil
}

func validateNewFluentBitKinesisFirehoseOutputParameters(options *FluentBitKinesisFirehoseOutputOptions) error {
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

