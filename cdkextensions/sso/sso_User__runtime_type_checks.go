//go:build !no_runtime_type_checking
// +build !no_runtime_type_checking

package sso

import (
	"fmt"

	"github.com/aws/constructs-go/constructs/v10"
)

func validateUser_FromUserIdParameters(scope constructs.IConstruct, id *string, userId *string) error {
	if scope == nil {
		return fmt.Errorf("parameter scope is required, but nil was provided")
	}

	if id == nil {
		return fmt.Errorf("parameter id is required, but nil was provided")
	}

	if userId == nil {
		return fmt.Errorf("parameter userId is required, but nil was provided")
	}

	return nil
}

