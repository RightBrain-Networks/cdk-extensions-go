//go:build !no_runtime_type_checking

package k8saws

import (
	"fmt"

	"github.com/aws/constructs-go/constructs/v10"
)

func (n *jsiiProxy_NoopRegistry) validateBindParameters(_scope constructs.IConstruct) error {
	if _scope == nil {
		return fmt.Errorf("parameter _scope is required, but nil was provided")
	}

	return nil
}

