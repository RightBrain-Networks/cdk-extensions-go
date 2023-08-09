//go:build no_runtime_type_checking

package config

// Building without runtime type checking enabled, so all the below just return nil

func validateRemediationTarget_AutomationDocumentParameters(props *AutomationDocumentRemediationProps) error {
	return nil
}

