//go:build !no_runtime_type_checking
// +build !no_runtime_type_checking

package core

import (
	"fmt"
)

func validateDataSize_BytesParameters(bytes *float64) error {
	if bytes == nil {
		return fmt.Errorf("parameter bytes is required, but nil was provided")
	}

	return nil
}

func validateDataSize_GibibytesParameters(gibibytes *float64) error {
	if gibibytes == nil {
		return fmt.Errorf("parameter gibibytes is required, but nil was provided")
	}

	return nil
}

func validateDataSize_GigabytesParameters(gigabytes *float64) error {
	if gigabytes == nil {
		return fmt.Errorf("parameter gigabytes is required, but nil was provided")
	}

	return nil
}

func validateDataSize_KibibytesParameters(kibibytes *float64) error {
	if kibibytes == nil {
		return fmt.Errorf("parameter kibibytes is required, but nil was provided")
	}

	return nil
}

func validateDataSize_KilobytesParameters(kilobytes *float64) error {
	if kilobytes == nil {
		return fmt.Errorf("parameter kilobytes is required, but nil was provided")
	}

	return nil
}

func validateDataSize_MebibytesParameters(mebibytes *float64) error {
	if mebibytes == nil {
		return fmt.Errorf("parameter mebibytes is required, but nil was provided")
	}

	return nil
}

func validateDataSize_MegabytesParameters(megabytes *float64) error {
	if megabytes == nil {
		return fmt.Errorf("parameter megabytes is required, but nil was provided")
	}

	return nil
}

func validateDataSize_PebibytesParameters(pebibytes *float64) error {
	if pebibytes == nil {
		return fmt.Errorf("parameter pebibytes is required, but nil was provided")
	}

	return nil
}

func validateDataSize_PetabytesParameters(petabytes *float64) error {
	if petabytes == nil {
		return fmt.Errorf("parameter petabytes is required, but nil was provided")
	}

	return nil
}

func validateDataSize_TebibytesParameters(tebibytes *float64) error {
	if tebibytes == nil {
		return fmt.Errorf("parameter tebibytes is required, but nil was provided")
	}

	return nil
}

func validateDataSize_TerabytesParameters(terabytes *float64) error {
	if terabytes == nil {
		return fmt.Errorf("parameter terabytes is required, but nil was provided")
	}

	return nil
}

