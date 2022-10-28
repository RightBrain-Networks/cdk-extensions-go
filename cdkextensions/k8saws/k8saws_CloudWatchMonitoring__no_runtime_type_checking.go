//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package k8saws

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CloudWatchMonitoring) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (c *jsiiProxy_CloudWatchMonitoring) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (c *jsiiProxy_CloudWatchMonitoring) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func validateCloudWatchMonitoring_IsConstructParameters(x interface{}) error {
	return nil
}

func validateCloudWatchMonitoring_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCloudWatchMonitoring_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewCloudWatchMonitoringParameters(scope constructs.Construct, id *string, props *CloudWatchMonitoringProps) error {
	return nil
}

