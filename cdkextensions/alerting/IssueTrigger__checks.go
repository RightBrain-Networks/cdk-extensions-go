//go:build !no_runtime_type_checking

package alerting

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsstepfunctions"
	"github.com/aws/constructs-go/constructs/v10"
)

func (i *jsiiProxy_IssueTrigger) validateAddOverrideParameters(handlerOverrides IssueHandlerOverride) error {
	if handlerOverrides == nil {
		return fmt.Errorf("parameter handlerOverrides is required, but nil was provided")
	}

	return nil
}

func (i *jsiiProxy_IssueTrigger) validateBindParameters(stateMachine awsstepfunctions.StateMachine) error {
	if stateMachine == nil {
		return fmt.Errorf("parameter stateMachine is required, but nil was provided")
	}

	return nil
}

func validateIssueTrigger_IsConstructParameters(x interface{}) error {
	if x == nil {
		return fmt.Errorf("parameter x is required, but nil was provided")
	}

	return nil
}

func validateNewIssueTriggerParameters(scope constructs.IConstruct, id *string, props *IssueTriggerProps) error {
	if scope == nil {
		return fmt.Errorf("parameter scope is required, but nil was provided")
	}

	if id == nil {
		return fmt.Errorf("parameter id is required, but nil was provided")
	}

	if props == nil {
		return fmt.Errorf("parameter props is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

