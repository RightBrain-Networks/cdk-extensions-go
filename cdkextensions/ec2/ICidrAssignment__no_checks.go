//go:build no_runtime_type_checking

package ec2

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_ICidrAssignment) validateGetCidrParameters(scope constructs.IConstruct, id *string, options *CidrAssignmentBindOptions) error {
	return nil
}

func (i *jsiiProxy_ICidrAssignment) validateGetCidrOrIpamConfigurationParameters(options *CidrAssignmentBindOptions) error {
	return nil
}

