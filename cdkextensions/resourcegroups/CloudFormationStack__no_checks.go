//go:build no_runtime_type_checking

package resourcegroups

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CloudFormationStack) validateAddResourceTypeParameters(typeId *string) error {
	return nil
}

func (c *jsiiProxy_CloudFormationStack) validateBindParameters(_scope constructs.IConstruct) error {
	return nil
}

func validateNewCloudFormationStackParameters(reference IStackReference, props *CloudFormationStackProps) error {
	return nil
}

