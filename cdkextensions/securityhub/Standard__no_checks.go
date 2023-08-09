//go:build no_runtime_type_checking

package securityhub

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_Standard) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (s *jsiiProxy_Standard) validateDisableControlParameters(control *string, options *DisableControlOptions) error {
	return nil
}

func (s *jsiiProxy_Standard) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (s *jsiiProxy_Standard) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func validateStandard_FromStandardArnParameters(scope constructs.IConstruct, id *string, arn *string) error {
	return nil
}

func validateStandard_IsConstructParameters(x interface{}) error {
	return nil
}

func validateStandard_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateStandard_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewStandardParameters(scope constructs.IConstruct, id *string, props *StandardProps) error {
	return nil
}

