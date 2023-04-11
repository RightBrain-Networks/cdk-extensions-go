//go:build !no_runtime_type_checking

package aps

import (
	"fmt"

	"github.com/aws/constructs-go/constructs/v10"
)

func (a *jsiiProxy_AlertManagerInhibitRule) validateAddEqualLabelParameters(label *string) error {
	if label == nil {
		return fmt.Errorf("parameter label is required, but nil was provided")
	}

	return nil
}

func (a *jsiiProxy_AlertManagerInhibitRule) validateAddSourceMatcherParameters(matcher AlertManagerMatcher) error {
	if matcher == nil {
		return fmt.Errorf("parameter matcher is required, but nil was provided")
	}

	return nil
}

func (a *jsiiProxy_AlertManagerInhibitRule) validateAddTargetMatcherParameters(matcher AlertManagerMatcher) error {
	if matcher == nil {
		return fmt.Errorf("parameter matcher is required, but nil was provided")
	}

	return nil
}

func (a *jsiiProxy_AlertManagerInhibitRule) validateBindParameters(_scope constructs.IConstruct) error {
	if _scope == nil {
		return fmt.Errorf("parameter _scope is required, but nil was provided")
	}

	return nil
}

func validateAlertManagerInhibitRule_IsConstructParameters(x interface{}) error {
	if x == nil {
		return fmt.Errorf("parameter x is required, but nil was provided")
	}

	return nil
}

func validateNewAlertManagerInhibitRuleParameters(scope AlertManagerConfiguration, id *string, options AlertManagerInhibitRuleProps) error {
	if scope == nil {
		return fmt.Errorf("parameter scope is required, but nil was provided")
	}

	if id == nil {
		return fmt.Errorf("parameter id is required, but nil was provided")
	}

	if options == nil {
		return fmt.Errorf("parameter options is required, but nil was provided")
	}

	return nil
}

