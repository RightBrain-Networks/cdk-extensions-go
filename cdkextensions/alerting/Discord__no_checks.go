//go:build no_runtime_type_checking

package alerting

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_Discord) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (d *jsiiProxy_Discord) validateBuildEventOverridesParameters(options *DiscordOverrideOptions) error {
	return nil
}

func (d *jsiiProxy_Discord) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (d *jsiiProxy_Discord) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func validateDiscord_IsConstructParameters(x interface{}) error {
	return nil
}

func validateDiscord_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateDiscord_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewDiscordParameters(scope constructs.IConstruct, id *string, props *DiscordProps) error {
	return nil
}

