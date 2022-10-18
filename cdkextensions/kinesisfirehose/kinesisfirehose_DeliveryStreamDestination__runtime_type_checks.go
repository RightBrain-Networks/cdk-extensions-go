//go:build !no_runtime_type_checking
// +build !no_runtime_type_checking

package kinesisfirehose

import (
	"fmt"

	"github.com/aws/constructs-go/constructs/v10"
)

func (d *jsiiProxy_DeliveryStreamDestination) validateBindParameters(scope constructs.IConstruct) error {
	if scope == nil {
		return fmt.Errorf("parameter scope is required, but nil was provided")
	}

	return nil
}

