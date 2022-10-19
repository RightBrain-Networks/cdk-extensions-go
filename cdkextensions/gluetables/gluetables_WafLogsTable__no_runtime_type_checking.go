//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package gluetables

// Building without runtime type checking enabled, so all the below just return nil

func (w *jsiiProxy_WafLogsTable) validateAddColumnParameters(column glue.Column) error {
	return nil
}

func (w *jsiiProxy_WafLogsTable) validateAddParameterParameters(key *string, value *string) error {
	return nil
}

func (w *jsiiProxy_WafLogsTable) validateAddPartitionKeyParameters(column glue.Column) error {
	return nil
}

func (w *jsiiProxy_WafLogsTable) validateAddSerdeParameterParameters(key *string, value *string) error {
	return nil
}

func (w *jsiiProxy_WafLogsTable) validateAddStorageParameterParameters(key *string, value *string) error {
	return nil
}

func (w *jsiiProxy_WafLogsTable) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (w *jsiiProxy_WafLogsTable) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (w *jsiiProxy_WafLogsTable) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func validateWafLogsTable_IsConstructParameters(x interface{}) error {
	return nil
}

func validateWafLogsTable_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateWafLogsTable_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewWafLogsTableParameters(scope constructs.Construct, id *string, props *WafLogsTableProps) error {
	return nil
}

