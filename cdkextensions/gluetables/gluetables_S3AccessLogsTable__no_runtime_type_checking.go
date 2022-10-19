//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package gluetables

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_S3AccessLogsTable) validateAddColumnParameters(column glue.Column) error {
	return nil
}

func (s *jsiiProxy_S3AccessLogsTable) validateAddParameterParameters(key *string, value *string) error {
	return nil
}

func (s *jsiiProxy_S3AccessLogsTable) validateAddPartitionKeyParameters(column glue.Column) error {
	return nil
}

func (s *jsiiProxy_S3AccessLogsTable) validateAddSerdeParameterParameters(key *string, value *string) error {
	return nil
}

func (s *jsiiProxy_S3AccessLogsTable) validateAddStorageParameterParameters(key *string, value *string) error {
	return nil
}

func (s *jsiiProxy_S3AccessLogsTable) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (s *jsiiProxy_S3AccessLogsTable) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (s *jsiiProxy_S3AccessLogsTable) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func validateS3AccessLogsTable_IsConstructParameters(x interface{}) error {
	return nil
}

func validateS3AccessLogsTable_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateS3AccessLogsTable_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewS3AccessLogsTableParameters(scope constructs.Construct, id *string, props *S3AccessLogsTableProps) error {
	return nil
}

