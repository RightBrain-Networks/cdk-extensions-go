//go:build no_runtime_type_checking

package gluetables

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_AlbLogsTable) validateAddColumnParameters(column glue.Column) error {
	return nil
}

func (a *jsiiProxy_AlbLogsTable) validateAddParameterParameters(key *string, value *string) error {
	return nil
}

func (a *jsiiProxy_AlbLogsTable) validateAddPartitionKeyParameters(column glue.Column) error {
	return nil
}

func (a *jsiiProxy_AlbLogsTable) validateAddSerdeParameterParameters(key *string, value *string) error {
	return nil
}

func (a *jsiiProxy_AlbLogsTable) validateAddStorageParameterParameters(key *string, value *string) error {
	return nil
}

func (a *jsiiProxy_AlbLogsTable) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (a *jsiiProxy_AlbLogsTable) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (a *jsiiProxy_AlbLogsTable) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func validateAlbLogsTable_IsConstructParameters(x interface{}) error {
	return nil
}

func validateAlbLogsTable_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateAlbLogsTable_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewAlbLogsTableParameters(scope constructs.Construct, id *string, props *AlbLogsTableProps) error {
	return nil
}

