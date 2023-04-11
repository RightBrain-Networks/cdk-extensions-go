//go:build !no_runtime_type_checking

package sso

import (
	"fmt"

	"github.com/aws/constructs-go/constructs/v10"
)

func validateGroup_FromGroupIdParameters(scope constructs.IConstruct, id *string, groupId *string) error {
	if scope == nil {
		return fmt.Errorf("parameter scope is required, but nil was provided")
	}

	if id == nil {
		return fmt.Errorf("parameter id is required, but nil was provided")
	}

	if groupId == nil {
		return fmt.Errorf("parameter groupId is required, but nil was provided")
	}

	return nil
}

