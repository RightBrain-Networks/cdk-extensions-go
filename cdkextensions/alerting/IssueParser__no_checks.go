//go:build no_runtime_type_checking

package alerting

// Building without runtime type checking enabled, so all the below just return nil

func validateIssueParser_ConfigComplianceChangeParameters(scope constructs.IConstruct, id *string, props *ConfigComplianceChangeProps) error {
	return nil
}

func validateIssueParser_EcrScanFindingParameters(scope constructs.IConstruct, id *string, props *EcrScanFindingProps) error {
	return nil
}

func validateIssueParser_GuardDutyFindingParameters(scope constructs.IConstruct, id *string, props *GuardDutyFindingProps) error {
	return nil
}

func validateIssueParser_InspectorFindingParameters(scope constructs.IConstruct, id *string, props *InspectorFindingProps) error {
	return nil
}

func validateIssueParser_OpenSearchEventParameters(scope constructs.IConstruct, id *string, props *OpenSearchEventProps) error {
	return nil
}

func validateIssueParser_SecurityHubFindingParameters(scope constructs.IConstruct, id *string, props *SecurityHubFindingProps) error {
	return nil
}

