//go:build no_runtime_type_checking

package k8saws

// Building without runtime type checking enabled, so all the below just return nil

func (f *jsiiProxy_FargateLogger) validateAddFargateProfileParameters(profile awseks.FargateProfile) error {
	return nil
}

func (f *jsiiProxy_FargateLogger) validateAddFilterParameters(filter IFluentBitFilterPlugin) error {
	return nil
}

func (f *jsiiProxy_FargateLogger) validateAddOutputParameters(output IFluentBitOutputPlugin) error {
	return nil
}

func (f *jsiiProxy_FargateLogger) validateAddParserParameters(parser IFluentBitParserPlugin) error {
	return nil
}

func (f *jsiiProxy_FargateLogger) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (f *jsiiProxy_FargateLogger) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (f *jsiiProxy_FargateLogger) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func validateFargateLogger_IsConstructParameters(x interface{}) error {
	return nil
}

func validateFargateLogger_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateFargateLogger_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewFargateLoggerParameters(scope constructs.Construct, id *string, props *FargateLoggerProps) error {
	return nil
}

