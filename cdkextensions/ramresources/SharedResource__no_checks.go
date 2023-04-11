//go:build no_runtime_type_checking

package ramresources

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_SharedResource) validateShareParameters(_scope constructs.IConstruct) error {
	return nil
}

func validateSharedResource_FromArnParameters(arn *string) error {
	return nil
}

func validateSharedResource_FromIpamPoolParameters(pool ec2.IIpamPool) error {
	return nil
}

func validateSharedResource_FromProjectParameters(project awscodebuild.IProject) error {
	return nil
}

func validateSharedResource_FromSubnetParameters(subnet awsec2.ISubnet) error {
	return nil
}

func validateSharedResource_FromTransitGatewayParameters(transitGateway ec2.ITransitGateway) error {
	return nil
}

