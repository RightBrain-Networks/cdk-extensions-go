//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package s3buckets

// Building without runtime type checking enabled, so all the below just return nil

func (w *jsiiProxy_WafLogsBucket) validateAddEventNotificationParameters(_event awss3.EventType, _dest awss3.IBucketNotificationDestination, _filters *[]*awss3.NotificationKeyFilter) error {
	return nil
}

func (w *jsiiProxy_WafLogsBucket) validateAddObjectCreatedNotificationParameters(_dest awss3.IBucketNotificationDestination, _filters *[]*awss3.NotificationKeyFilter) error {
	return nil
}

func (w *jsiiProxy_WafLogsBucket) validateAddObjectRemovedNotificationParameters(_dest awss3.IBucketNotificationDestination, _filters *[]*awss3.NotificationKeyFilter) error {
	return nil
}

func (w *jsiiProxy_WafLogsBucket) validateAddToResourcePolicyParameters(permission awsiam.PolicyStatement) error {
	return nil
}

func (w *jsiiProxy_WafLogsBucket) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (w *jsiiProxy_WafLogsBucket) validateArnForObjectsParameters(_keyPattern *string) error {
	return nil
}

func (w *jsiiProxy_WafLogsBucket) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (w *jsiiProxy_WafLogsBucket) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func (w *jsiiProxy_WafLogsBucket) validateGrantDeleteParameters(_identity awsiam.IGrantable) error {
	return nil
}

func (w *jsiiProxy_WafLogsBucket) validateGrantPutParameters(_identity awsiam.IGrantable) error {
	return nil
}

func (w *jsiiProxy_WafLogsBucket) validateGrantPutAclParameters(_identity awsiam.IGrantable) error {
	return nil
}

func (w *jsiiProxy_WafLogsBucket) validateGrantReadParameters(_identity awsiam.IGrantable) error {
	return nil
}

func (w *jsiiProxy_WafLogsBucket) validateGrantReadWriteParameters(_identity awsiam.IGrantable) error {
	return nil
}

func (w *jsiiProxy_WafLogsBucket) validateGrantWriteParameters(_identity awsiam.IGrantable) error {
	return nil
}

func (w *jsiiProxy_WafLogsBucket) validateOnCloudTrailEventParameters(_id *string, _options *awss3.OnCloudTrailBucketEventOptions) error {
	return nil
}

func (w *jsiiProxy_WafLogsBucket) validateOnCloudTrailPutObjectParameters(_id *string, _options *awss3.OnCloudTrailBucketEventOptions) error {
	return nil
}

func (w *jsiiProxy_WafLogsBucket) validateOnCloudTrailWriteObjectParameters(_id *string, _options *awss3.OnCloudTrailBucketEventOptions) error {
	return nil
}

func (w *jsiiProxy_WafLogsBucket) validateTransferAccelerationUrlForObjectParameters(_options *awss3.TransferAccelerationUrlOptions) error {
	return nil
}

func (w *jsiiProxy_WafLogsBucket) validateVirtualHostedUrlForObjectParameters(_options *awss3.VirtualHostedStyleUrlOptions) error {
	return nil
}

func validateWafLogsBucket_IsConstructParameters(x interface{}) error {
	return nil
}

func validateWafLogsBucket_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateWafLogsBucket_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewWafLogsBucketParameters(scope constructs.Construct, id *string, props *WafLogsBucketProps) error {
	return nil
}

