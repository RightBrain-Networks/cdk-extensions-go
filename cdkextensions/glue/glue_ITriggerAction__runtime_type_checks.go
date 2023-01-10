//go:build !no_runtime_type_checking
// +build !no_runtime_type_checking

package glue

import (
	"fmt"

	"github.com/aws/constructs-go/constructs/v10"
)

func (i *jsiiProxy_ITriggerAction) validateBindParameters(scope constructs.IConstruct) error {
	if scope == nil {
		return fmt.Errorf("parameter scope is required, but nil was provided")
	}

	return nil
}

