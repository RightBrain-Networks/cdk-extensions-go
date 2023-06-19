//go:build no_runtime_type_checking

package alerting

// Building without runtime type checking enabled, so all the below just return nil

func validateNewIssueHandlerOverrideParameters(handler IIssueHandler, overrides *map[string]interface{}) error {
	return nil
}

