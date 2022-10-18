//go:build !no_runtime_type_checking
// +build !no_runtime_type_checking

package glue

import (
	"fmt"
)

func (i *jsiiProxy_ITriggerPredicate) validateBindParameters(trigger Trigger) error {
	if trigger == nil {
		return fmt.Errorf("parameter trigger is required, but nil was provided")
	}

	return nil
}

