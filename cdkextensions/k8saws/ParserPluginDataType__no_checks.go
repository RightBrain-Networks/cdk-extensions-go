//go:build no_runtime_type_checking

package k8saws

// Building without runtime type checking enabled, so all the below just return nil

func validateParserPluginDataType_OfParameters(name *string) error {
	return nil
}

