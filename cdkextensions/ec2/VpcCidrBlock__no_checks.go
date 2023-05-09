//go:build no_runtime_type_checking

package ec2

// Building without runtime type checking enabled, so all the below just return nil

func (v *jsiiProxy_VpcCidrBlock) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (v *jsiiProxy_VpcCidrBlock) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (v *jsiiProxy_VpcCidrBlock) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func validateVpcCidrBlock_FromVpcCidrBlockAttributesParameters(scope constructs.IConstruct, id *string, attrs *VpcCidrBlockAttributes) error {
	return nil
}

func validateVpcCidrBlock_IsConstructParameters(x interface{}) error {
	return nil
}

func validateVpcCidrBlock_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateVpcCidrBlock_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewVpcCidrBlockParameters(scope constructs.IConstruct, id *string, props *VpcCidrBlockProps) error {
	return nil
}

