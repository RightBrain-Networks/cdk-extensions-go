//go:build !no_runtime_type_checking
// +build !no_runtime_type_checking

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

