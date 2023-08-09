//go:build !no_runtime_type_checking

package securityhub

import (
	"fmt"

	"github.com/aws/constructs-go/constructs/v10"
)

func (r *jsiiProxy_RuleSet) validateBindParameters(scope constructs.IConstruct) error {
	if scope == nil {
		return fmt.Errorf("parameter scope is required, but nil was provided")
	}

	return nil
}

