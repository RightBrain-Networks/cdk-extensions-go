//go:build !no_runtime_type_checking

package guardduty

import (
	"fmt"
)

func validateFindingPublishingFrequency_OfParameters(label *string) error {
	if label == nil {
		return fmt.Errorf("parameter label is required, but nil was provided")
	}

	return nil
}

