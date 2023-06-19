//go:build no_runtime_type_checking

package alerting

// Building without runtime type checking enabled, so all the below just return nil

func validateIssueHander_JiraTicketParameters(scope constructs.IConstruct, id *string, props *JiraTicketProps) error {
	return nil
}

