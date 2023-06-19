//go:build !no_runtime_type_checking

package alerting

import (
	"fmt"
)

func validateNewIssueHandlerOverrideParameters(handler IIssueHandler, overrides *map[string]interface{}) error {
	if handler == nil {
		return fmt.Errorf("parameter handler is required, but nil was provided")
	}

	if overrides == nil {
		return fmt.Errorf("parameter overrides is required, but nil was provided")
	}

	return nil
}

