//go:build !no_runtime_type_checking

package ec2

import (
	"fmt"
)

func (f *jsiiProxy_FlowLogFormat) validateAddFieldParameters(field FlowLogField) error {
	if field == nil {
		return fmt.Errorf("parameter field is required, but nil was provided")
	}

	return nil
}

func validateFlowLogFormat_FromTemplateParameters(template *string) error {
	if template == nil {
		return fmt.Errorf("parameter template is required, but nil was provided")
	}

	return nil
}

