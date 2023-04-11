//go:build no_runtime_type_checking

package k8saws

// Building without runtime type checking enabled, so all the below just return nil

func (r *jsiiProxy_Route53Dns) validateAddDomainFilterParameters(domain *string) error {
	return nil
}

func (r *jsiiProxy_Route53Dns) validateAddZoneTagParameters(tag *ExternalDnsZoneTag) error {
	return nil
}

func (r *jsiiProxy_Route53Dns) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (r *jsiiProxy_Route53Dns) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (r *jsiiProxy_Route53Dns) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func validateRoute53Dns_IsConstructParameters(x interface{}) error {
	return nil
}

func validateRoute53Dns_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateRoute53Dns_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewRoute53DnsParameters(scope constructs.Construct, id *string, props *Route53DnsProps) error {
	return nil
}

