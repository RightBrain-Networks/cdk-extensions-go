//go:build no_runtime_type_checking

package ssm

// Building without runtime type checking enabled, so all the below just return nil

func validateDocumentContent_FromObjectParameters(props *ObjectContentProps) error {
	return nil
}

func validateDocumentContent_FromStringParameters(props *StringContentProps) error {
	return nil
}

