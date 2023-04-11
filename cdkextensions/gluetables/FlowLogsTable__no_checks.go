//go:build no_runtime_type_checking

package gluetables

// Building without runtime type checking enabled, so all the below just return nil

func (f *jsiiProxy_FlowLogsTable) validateAddColumnParameters(column glue.Column) error {
	return nil
}

func (f *jsiiProxy_FlowLogsTable) validateAddParameterParameters(key *string, value *string) error {
	return nil
}

func (f *jsiiProxy_FlowLogsTable) validateAddPartitionKeyParameters(column glue.Column) error {
	return nil
}

func (f *jsiiProxy_FlowLogsTable) validateAddSerdeParameterParameters(key *string, value *string) error {
	return nil
}

func (f *jsiiProxy_FlowLogsTable) validateAddStorageParameterParameters(key *string, value *string) error {
	return nil
}

func (f *jsiiProxy_FlowLogsTable) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (f *jsiiProxy_FlowLogsTable) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (f *jsiiProxy_FlowLogsTable) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func validateFlowLogsTable_IsConstructParameters(x interface{}) error {
	return nil
}

func validateFlowLogsTable_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateFlowLogsTable_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewFlowLogsTableParameters(scope constructs.Construct, id *string, props *FlowLogsTableProps) error {
	return nil
}

