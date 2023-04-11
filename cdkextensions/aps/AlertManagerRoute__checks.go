//go:build !no_runtime_type_checking

package aps

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/constructs-go/constructs/v10"
)

func (a *jsiiProxy_AlertManagerRoute) validateAddActiveTimeIntervalParameters(interval TimeInterval) error {
	if interval == nil {
		return fmt.Errorf("parameter interval is required, but nil was provided")
	}

	return nil
}

func (a *jsiiProxy_AlertManagerRoute) validateAddChildParameters(id *string, options *AlertManagerRouteProps) error {
	if id == nil {
		return fmt.Errorf("parameter id is required, but nil was provided")
	}

	if options == nil {
		return fmt.Errorf("parameter options is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func (a *jsiiProxy_AlertManagerRoute) validateAddGroupByLabelParameters(label *string) error {
	if label == nil {
		return fmt.Errorf("parameter label is required, but nil was provided")
	}

	return nil
}

func (a *jsiiProxy_AlertManagerRoute) validateAddMatcherParameters(matcher AlertManagerMatcher) error {
	if matcher == nil {
		return fmt.Errorf("parameter matcher is required, but nil was provided")
	}

	return nil
}

func (a *jsiiProxy_AlertManagerRoute) validateAddMuteTimeIntervalParameters(interval TimeInterval) error {
	if interval == nil {
		return fmt.Errorf("parameter interval is required, but nil was provided")
	}

	return nil
}

func (a *jsiiProxy_AlertManagerRoute) validateBindParameters(scope constructs.IConstruct) error {
	if scope == nil {
		return fmt.Errorf("parameter scope is required, but nil was provided")
	}

	return nil
}

func validateAlertManagerRoute_IsConstructParameters(x interface{}) error {
	if x == nil {
		return fmt.Errorf("parameter x is required, but nil was provided")
	}

	return nil
}

func validateNewAlertManagerRouteParameters(scope interface{}, id *string, options *AlertManagerRouteProps) error {
	if scope == nil {
		return fmt.Errorf("parameter scope is required, but nil was provided")
	}
	switch scope.(type) {
	case AlertManagerConfiguration:
		// ok
	case AlertManagerRoute:
		// ok
	default:
		if !_jsii_.IsAnonymousProxy(scope) {
			return fmt.Errorf("parameter scope must be one of the allowed types: AlertManagerConfiguration, AlertManagerRoute; received %#v (a %T)", scope, scope)
		}
	}

	if id == nil {
		return fmt.Errorf("parameter id is required, but nil was provided")
	}

	if options == nil {
		return fmt.Errorf("parameter options is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

