//go:build no_runtime_type_checking

package networkmanager

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_Site) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (s *jsiiProxy_Site) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (s *jsiiProxy_Site) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func validateSite_IsConstructParameters(x interface{}) error {
	return nil
}

func validateSite_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateSite_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewSiteParameters(scope constructs.IConstruct, id *string, props *SiteProps) error {
	return nil
}

