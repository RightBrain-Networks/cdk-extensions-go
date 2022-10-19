//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package s3buckets

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CloudtrailBucket) validateAddEventNotificationParameters(_event awss3.EventType, _dest awss3.IBucketNotificationDestination, _filters *[]*awss3.NotificationKeyFilter) error {
	return nil
}

func (c *jsiiProxy_CloudtrailBucket) validateAddObjectCreatedNotificationParameters(_dest awss3.IBucketNotificationDestination, _filters *[]*awss3.NotificationKeyFilter) error {
	return nil
}

func (c *jsiiProxy_CloudtrailBucket) validateAddObjectRemovedNotificationParameters(_dest awss3.IBucketNotificationDestination, _filters *[]*awss3.NotificationKeyFilter) error {
	return nil
}

func (c *jsiiProxy_CloudtrailBucket) validateAddToResourcePolicyParameters(permission awsiam.PolicyStatement) error {
	return nil
}

func (c *jsiiProxy_CloudtrailBucket) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (c *jsiiProxy_CloudtrailBucket) validateArnForObjectsParameters(_keyPattern *string) error {
	return nil
}

func (c *jsiiProxy_CloudtrailBucket) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (c *jsiiProxy_CloudtrailBucket) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func (c *jsiiProxy_CloudtrailBucket) validateGrantDeleteParameters(_identity awsiam.IGrantable) error {
	return nil
}

func (c *jsiiProxy_CloudtrailBucket) validateGrantPutParameters(_identity awsiam.IGrantable) error {
	return nil
}

func (c *jsiiProxy_CloudtrailBucket) validateGrantPutAclParameters(_identity awsiam.IGrantable) error {
	return nil
}

func (c *jsiiProxy_CloudtrailBucket) validateGrantReadParameters(_identity awsiam.IGrantable) error {
	return nil
}

func (c *jsiiProxy_CloudtrailBucket) validateGrantReadWriteParameters(_identity awsiam.IGrantable) error {
	return nil
}

func (c *jsiiProxy_CloudtrailBucket) validateGrantWriteParameters(_identity awsiam.IGrantable) error {
	return nil
}

func (c *jsiiProxy_CloudtrailBucket) validateOnCloudTrailEventParameters(_id *string, _options *awss3.OnCloudTrailBucketEventOptions) error {
	return nil
}

func (c *jsiiProxy_CloudtrailBucket) validateOnCloudTrailPutObjectParameters(_id *string, _options *awss3.OnCloudTrailBucketEventOptions) error {
	return nil
}

func (c *jsiiProxy_CloudtrailBucket) validateOnCloudTrailWriteObjectParameters(_id *string, _options *awss3.OnCloudTrailBucketEventOptions) error {
	return nil
}

func (c *jsiiProxy_CloudtrailBucket) validateTransferAccelerationUrlForObjectParameters(_options *awss3.TransferAccelerationUrlOptions) error {
	return nil
}

func (c *jsiiProxy_CloudtrailBucket) validateVirtualHostedUrlForObjectParameters(_options *awss3.VirtualHostedStyleUrlOptions) error {
	return nil
}

func validateCloudtrailBucket_IsConstructParameters(x interface{}) error {
	return nil
}

func validateCloudtrailBucket_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCloudtrailBucket_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewCloudtrailBucketParameters(scope constructs.Construct, id *string, props *CloudtrailBucketProps) error {
	return nil
}

