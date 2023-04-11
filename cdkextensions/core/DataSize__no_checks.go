//go:build no_runtime_type_checking

package core

// Building without runtime type checking enabled, so all the below just return nil

func validateDataSize_BytesParameters(bytes *float64) error {
	return nil
}

func validateDataSize_GibibytesParameters(gibibytes *float64) error {
	return nil
}

func validateDataSize_GigabytesParameters(gigabytes *float64) error {
	return nil
}

func validateDataSize_KibibytesParameters(kibibytes *float64) error {
	return nil
}

func validateDataSize_KilobytesParameters(kilobytes *float64) error {
	return nil
}

func validateDataSize_MebibytesParameters(mebibytes *float64) error {
	return nil
}

func validateDataSize_MegabytesParameters(megabytes *float64) error {
	return nil
}

func validateDataSize_PebibytesParameters(pebibytes *float64) error {
	return nil
}

func validateDataSize_PetabytesParameters(petabytes *float64) error {
	return nil
}

func validateDataSize_TebibytesParameters(tebibytes *float64) error {
	return nil
}

func validateDataSize_TerabytesParameters(terabytes *float64) error {
	return nil
}

