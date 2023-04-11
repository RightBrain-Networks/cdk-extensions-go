//go:build no_runtime_type_checking

package glue

// Building without runtime type checking enabled, so all the below just return nil

func (j *jsiiProxy_JdbcConnection) validateAddMatchCriteriaParameters(value *string) error {
	return nil
}

func (j *jsiiProxy_JdbcConnection) validateAddPropertyParameters(key *string, value *string) error {
	return nil
}

func (j *jsiiProxy_JdbcConnection) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (j *jsiiProxy_JdbcConnection) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (j *jsiiProxy_JdbcConnection) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func validateJdbcConnection_IsConstructParameters(x interface{}) error {
	return nil
}

func validateJdbcConnection_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateJdbcConnection_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewJdbcConnectionParameters(scope constructs.Construct, id *string, props *JdbcConnectionProps) error {
	return nil
}

