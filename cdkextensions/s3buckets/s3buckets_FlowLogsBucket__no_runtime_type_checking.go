//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package s3buckets

// Building without runtime type checking enabled, so all the below just return nil

func (f *jsiiProxy_FlowLogsBucket) validateAddEventNotificationParameters(_event awss3.EventType, _dest awss3.IBucketNotificationDestination, _filters *[]*awss3.NotificationKeyFilter) error {
	return nil
}

func (f *jsiiProxy_FlowLogsBucket) validateAddObjectCreatedNotificationParameters(_dest awss3.IBucketNotificationDestination, _filters *[]*awss3.NotificationKeyFilter) error {
	return nil
}

func (f *jsiiProxy_FlowLogsBucket) validateAddObjectRemovedNotificationParameters(_dest awss3.IBucketNotificationDestination, _filters *[]*awss3.NotificationKeyFilter) error {
	return nil
}

func (f *jsiiProxy_FlowLogsBucket) validateAddToResourcePolicyParameters(permission awsiam.PolicyStatement) error {
	return nil
}

func (f *jsiiProxy_FlowLogsBucket) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (f *jsiiProxy_FlowLogsBucket) validateArnForObjectsParameters(_keyPattern *string) error {
	return nil
}

func (f *jsiiProxy_FlowLogsBucket) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (f *jsiiProxy_FlowLogsBucket) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func (f *jsiiProxy_FlowLogsBucket) validateGrantDeleteParameters(_identity awsiam.IGrantable) error {
	return nil
}

func (f *jsiiProxy_FlowLogsBucket) validateGrantPutParameters(_identity awsiam.IGrantable) error {
	return nil
}

func (f *jsiiProxy_FlowLogsBucket) validateGrantPutAclParameters(_identity awsiam.IGrantable) error {
	return nil
}

func (f *jsiiProxy_FlowLogsBucket) validateGrantReadParameters(_identity awsiam.IGrantable) error {
	return nil
}

func (f *jsiiProxy_FlowLogsBucket) validateGrantReadWriteParameters(_identity awsiam.IGrantable) error {
	return nil
}

func (f *jsiiProxy_FlowLogsBucket) validateGrantWriteParameters(_identity awsiam.IGrantable) error {
	return nil
}

func (f *jsiiProxy_FlowLogsBucket) validateOnCloudTrailEventParameters(_id *string, _options *awss3.OnCloudTrailBucketEventOptions) error {
	return nil
}

func (f *jsiiProxy_FlowLogsBucket) validateOnCloudTrailPutObjectParameters(_id *string, _options *awss3.OnCloudTrailBucketEventOptions) error {
	return nil
}

func (f *jsiiProxy_FlowLogsBucket) validateOnCloudTrailWriteObjectParameters(_id *string, _options *awss3.OnCloudTrailBucketEventOptions) error {
	return nil
}

func (f *jsiiProxy_FlowLogsBucket) validateTransferAccelerationUrlForObjectParameters(_options *awss3.TransferAccelerationUrlOptions) error {
	return nil
}

func (f *jsiiProxy_FlowLogsBucket) validateVirtualHostedUrlForObjectParameters(_options *awss3.VirtualHostedStyleUrlOptions) error {
	return nil
}

func validateFlowLogsBucket_IsConstructParameters(x interface{}) error {
	return nil
}

func validateFlowLogsBucket_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateFlowLogsBucket_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewFlowLogsBucketParameters(scope constructs.Construct, id *string, props *FlowLogsBucketProps) error {
	return nil
}

