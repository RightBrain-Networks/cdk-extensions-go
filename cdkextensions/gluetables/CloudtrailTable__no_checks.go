//go:build no_runtime_type_checking

package gluetables

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CloudtrailTable) validateAddColumnParameters(column glue.Column) error {
	return nil
}

func (c *jsiiProxy_CloudtrailTable) validateAddParameterParameters(key *string, value *string) error {
	return nil
}

func (c *jsiiProxy_CloudtrailTable) validateAddPartitionKeyParameters(column glue.Column) error {
	return nil
}

func (c *jsiiProxy_CloudtrailTable) validateAddSerdeParameterParameters(key *string, value *string) error {
	return nil
}

func (c *jsiiProxy_CloudtrailTable) validateAddStorageParameterParameters(key *string, value *string) error {
	return nil
}

func (c *jsiiProxy_CloudtrailTable) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (c *jsiiProxy_CloudtrailTable) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (c *jsiiProxy_CloudtrailTable) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func validateCloudtrailTable_IsConstructParameters(x interface{}) error {
	return nil
}

func validateCloudtrailTable_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCloudtrailTable_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewCloudtrailTableParameters(scope constructs.Construct, id *string, props *CloudtrailTableProps) error {
	return nil
}

