//go:build !no_runtime_type_checking

package core

import (
	"fmt"

	"github.com/aws/constructs-go/constructs/v10"
)

func (s *jsiiProxy_SecretReference) validateValueForScopeParameters(scope constructs.IConstruct) error {
	if scope == nil {
		return fmt.Errorf("parameter scope is required, but nil was provided")
	}

	return nil
}

func validateSecretReference_StringParameters(scope constructs.IConstruct, key *string, value *string) error {
	if scope == nil {
		return fmt.Errorf("parameter scope is required, but nil was provided")
	}

	if key == nil {
		return fmt.Errorf("parameter key is required, but nil was provided")
	}

	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}

	return nil
}

