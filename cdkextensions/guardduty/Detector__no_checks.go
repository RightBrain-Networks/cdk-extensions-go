//go:build no_runtime_type_checking

package guardduty

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_Detector) validateAddFeatureParameters(feature IFeature) error {
	return nil
}

func (d *jsiiProxy_Detector) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (d *jsiiProxy_Detector) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (d *jsiiProxy_Detector) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func validateDetector_FromDetectorArnParameters(scope constructs.IConstruct, id *string, detectorArn *string) error {
	return nil
}

func validateDetector_FromDetectorAttributesParameters(scope constructs.IConstruct, id *string, attrs *DetectorAttributes) error {
	return nil
}

func validateDetector_FromDetectorIdParameters(scope constructs.IConstruct, id *string, detectorId *string) error {
	return nil
}

func validateDetector_IsConstructParameters(x interface{}) error {
	return nil
}

func validateDetector_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateDetector_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewDetectorParameters(scope constructs.IConstruct, id *string, props *DetectorProps) error {
	return nil
}

