//go:build no_runtime_type_checking

package kinesisfirehose

// Building without runtime type checking enabled, so all the below just return nil

func validateSubRecordType_OfParameters(name *string) error {
	return nil
}

