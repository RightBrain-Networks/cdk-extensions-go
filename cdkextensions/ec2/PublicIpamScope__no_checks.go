//go:build no_runtime_type_checking

package ec2

// Building without runtime type checking enabled, so all the below just return nil

func validatePublicIpamScope_FromIpamScopeArnParameters(scope constructs.IConstruct, id *string, ipamScopeArn *string) error {
	return nil
}

func validatePublicIpamScope_FromIpamScopeAttributesParameters(scope constructs.IConstruct, id *string, attrs *IpamScopeAttributes) error {
	return nil
}

func validatePublicIpamScope_FromIpamScopeIdParameters(scope constructs.IConstruct, id *string, ipamScopeId *string) error {
	return nil
}

