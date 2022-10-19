//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package gluetables

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_SesLogsTable) validateAddColumnParameters(column glue.Column) error {
	return nil
}

func (s *jsiiProxy_SesLogsTable) validateAddParameterParameters(key *string, value *string) error {
	return nil
}

func (s *jsiiProxy_SesLogsTable) validateAddPartitionKeyParameters(column glue.Column) error {
	return nil
}

func (s *jsiiProxy_SesLogsTable) validateAddSerdeParameterParameters(key *string, value *string) error {
	return nil
}

func (s *jsiiProxy_SesLogsTable) validateAddStorageParameterParameters(key *string, value *string) error {
	return nil
}

func (s *jsiiProxy_SesLogsTable) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (s *jsiiProxy_SesLogsTable) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (s *jsiiProxy_SesLogsTable) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func validateSesLogsTable_IsConstructParameters(x interface{}) error {
	return nil
}

func validateSesLogsTable_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateSesLogsTable_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewSesLogsTableParameters(scope constructs.Construct, id *string, props *SesLogsTableProps) error {
	return nil
}

