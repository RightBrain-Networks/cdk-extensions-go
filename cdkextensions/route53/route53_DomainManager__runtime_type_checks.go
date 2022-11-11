//go:build !no_runtime_type_checking
// +build !no_runtime_type_checking

package route53

import (
	"fmt"

	"github.com/aws/constructs-go/constructs/v10"
)

func validateDomainManager_IsDnsResolvableParameters(node constructs.IConstruct) error {
	if node == nil {
		return fmt.Errorf("parameter node is required, but nil was provided")
	}

	return nil
}

