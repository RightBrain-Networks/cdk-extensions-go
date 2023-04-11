//go:build no_runtime_type_checking

package gluetables

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CloudfrontLogsTable) validateAddColumnParameters(column glue.Column) error {
	return nil
}

func (c *jsiiProxy_CloudfrontLogsTable) validateAddParameterParameters(key *string, value *string) error {
	return nil
}

func (c *jsiiProxy_CloudfrontLogsTable) validateAddPartitionKeyParameters(column glue.Column) error {
	return nil
}

func (c *jsiiProxy_CloudfrontLogsTable) validateAddSerdeParameterParameters(key *string, value *string) error {
	return nil
}

func (c *jsiiProxy_CloudfrontLogsTable) validateAddStorageParameterParameters(key *string, value *string) error {
	return nil
}

func (c *jsiiProxy_CloudfrontLogsTable) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (c *jsiiProxy_CloudfrontLogsTable) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (c *jsiiProxy_CloudfrontLogsTable) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func validateCloudfrontLogsTable_IsConstructParameters(x interface{}) error {
	return nil
}

func validateCloudfrontLogsTable_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCloudfrontLogsTable_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewCloudfrontLogsTableParameters(scope constructs.Construct, id *string, props *CloudfrontLogsTableProps) error {
	return nil
}

