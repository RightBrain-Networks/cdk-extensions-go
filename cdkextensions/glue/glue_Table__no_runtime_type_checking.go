//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package glue

// Building without runtime type checking enabled, so all the below just return nil

func (t *jsiiProxy_Table) validateAddColumnParameters(column Column) error {
	return nil
}

func (t *jsiiProxy_Table) validateAddParameterParameters(key *string, value *string) error {
	return nil
}

func (t *jsiiProxy_Table) validateAddPartitionKeyParameters(column Column) error {
	return nil
}

func (t *jsiiProxy_Table) validateAddSerdeParameterParameters(key *string, value *string) error {
	return nil
}

func (t *jsiiProxy_Table) validateAddStorageParameterParameters(key *string, value *string) error {
	return nil
}

func (t *jsiiProxy_Table) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (t *jsiiProxy_Table) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (t *jsiiProxy_Table) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func validateTable_IsConstructParameters(x interface{}) error {
	return nil
}

func validateTable_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateTable_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewTableParameters(scope constructs.Construct, id *string, props *TableProps) error {
	return nil
}

