//go:build !no_runtime_type_checking
// +build !no_runtime_type_checking

package k8saws

import (
	"fmt"

	"github.com/vibe-io/cdk-extensions-go/cdkextensions/core"
)

func validateElasticsearchOutputBufferSize_BytesParameters(size core.DataSize) error {
	if size == nil {
		return fmt.Errorf("parameter size is required, but nil was provided")
	}

	return nil
}

func validateElasticsearchOutputBufferSize_OfParameters(value *string) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}

	return nil
}

