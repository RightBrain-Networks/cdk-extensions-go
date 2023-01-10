//go:build !no_runtime_type_checking
// +build !no_runtime_type_checking

package glue

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/constructs-go/constructs/v10"
)

func (w *jsiiProxy_WorkflowJobAction) validateAddArgumentParameters(key *string, value *string) error {
	if key == nil {
		return fmt.Errorf("parameter key is required, but nil was provided")
	}

	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}

	return nil
}

func (w *jsiiProxy_WorkflowJobAction) validateBindParameters(scope constructs.IConstruct) error {
	if scope == nil {
		return fmt.Errorf("parameter scope is required, but nil was provided")
	}

	return nil
}

func (w *jsiiProxy_WorkflowJobAction) validateBindOptionsParameters(_scope constructs.IConstruct) error {
	if _scope == nil {
		return fmt.Errorf("parameter _scope is required, but nil was provided")
	}

	return nil
}

func validateNewWorkflowJobActionParameters(job IJob, options *WorkflowJobActionOptions) error {
	if job == nil {
		return fmt.Errorf("parameter job is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

