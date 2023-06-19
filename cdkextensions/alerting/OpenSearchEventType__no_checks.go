//go:build no_runtime_type_checking

package alerting

// Building without runtime type checking enabled, so all the below just return nil

func validateOpenSearchEventType_OfParameters(props *OpenSearchEventTypeProps) error {
	return nil
}

