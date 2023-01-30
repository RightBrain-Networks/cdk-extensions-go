//go:build !no_runtime_type_checking
// +build !no_runtime_type_checking

package aps

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/constructs-go/constructs/v10"
)

func (a *jsiiProxy_AlertManagerConfiguration) validateAddInhibitRuleParameters(id *string, options AlertManagerInhibitRuleProps) error {
	if id == nil {
		return fmt.Errorf("parameter id is required, but nil was provided")
	}

	if options == nil {
		return fmt.Errorf("parameter options is required, but nil was provided")
	}

	return nil
}

func (a *jsiiProxy_AlertManagerConfiguration) validateAddRecieverParameters(id *string, options *AlertManagerReceiverProps) error {
	if id == nil {
		return fmt.Errorf("parameter id is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func (a *jsiiProxy_AlertManagerConfiguration) validateAddTemplateParameters(template AlertManagerTemplate) error {
	if template == nil {
		return fmt.Errorf("parameter template is required, but nil was provided")
	}

	return nil
}

func (a *jsiiProxy_AlertManagerConfiguration) validateAddTimeIntervalParameters(id *string, options *TimeIntervalProps) error {
	if id == nil {
		return fmt.Errorf("parameter id is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func (a *jsiiProxy_AlertManagerConfiguration) validateBindParameters(scope constructs.IConstruct) error {
	if scope == nil {
		return fmt.Errorf("parameter scope is required, but nil was provided")
	}

	return nil
}

func (a *jsiiProxy_AlertManagerConfiguration) validateGenerateTemplateNameParameters(idx *float64) error {
	if idx == nil {
		return fmt.Errorf("parameter idx is required, but nil was provided")
	}

	return nil
}

func (a *jsiiProxy_AlertManagerConfiguration) validateRenderAlertManagerConfigParameters(scope constructs.IConstruct) error {
	if scope == nil {
		return fmt.Errorf("parameter scope is required, but nil was provided")
	}

	return nil
}

func (a *jsiiProxy_AlertManagerConfiguration) validateRenderTemplatesParameters(_scope constructs.IConstruct) error {
	if _scope == nil {
		return fmt.Errorf("parameter _scope is required, but nil was provided")
	}

	return nil
}

func validateAlertManagerConfiguration_FromFullConfigurationContentParameters(scope constructs.IConstruct, id *string, content *string) error {
	if scope == nil {
		return fmt.Errorf("parameter scope is required, but nil was provided")
	}

	if id == nil {
		return fmt.Errorf("parameter id is required, but nil was provided")
	}

	if content == nil {
		return fmt.Errorf("parameter content is required, but nil was provided")
	}

	return nil
}

func validateAlertManagerConfiguration_FromFullConfigurationFileParameters(scope constructs.IConstruct, id *string, path *string) error {
	if scope == nil {
		return fmt.Errorf("parameter scope is required, but nil was provided")
	}

	if id == nil {
		return fmt.Errorf("parameter id is required, but nil was provided")
	}

	if path == nil {
		return fmt.Errorf("parameter path is required, but nil was provided")
	}

	return nil
}

func validateAlertManagerConfiguration_FromSplitConfigurationContentParameters(scope constructs.IConstruct, id *string, content *string) error {
	if scope == nil {
		return fmt.Errorf("parameter scope is required, but nil was provided")
	}

	if id == nil {
		return fmt.Errorf("parameter id is required, but nil was provided")
	}

	if content == nil {
		return fmt.Errorf("parameter content is required, but nil was provided")
	}

	return nil
}

func validateAlertManagerConfiguration_FromSplitConfigurationFilesParameters(scope constructs.IConstruct, id *string, path *string) error {
	if scope == nil {
		return fmt.Errorf("parameter scope is required, but nil was provided")
	}

	if id == nil {
		return fmt.Errorf("parameter id is required, but nil was provided")
	}

	if path == nil {
		return fmt.Errorf("parameter path is required, but nil was provided")
	}

	return nil
}

func validateAlertManagerConfiguration_IsConstructParameters(x interface{}) error {
	if x == nil {
		return fmt.Errorf("parameter x is required, but nil was provided")
	}

	return nil
}

func validateNewAlertManagerConfigurationParameters(scope constructs.IConstruct, id *string, props *AlertManagerConfigurationProps) error {
	if scope == nil {
		return fmt.Errorf("parameter scope is required, but nil was provided")
	}

	if id == nil {
		return fmt.Errorf("parameter id is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

