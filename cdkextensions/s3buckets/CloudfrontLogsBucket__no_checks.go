//go:build no_runtime_type_checking

package s3buckets

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CloudfrontLogsBucket) validateAddEventNotificationParameters(_event awss3.EventType, _dest awss3.IBucketNotificationDestination, _filters *[]*awss3.NotificationKeyFilter) error {
	return nil
}

func (c *jsiiProxy_CloudfrontLogsBucket) validateAddObjectCreatedNotificationParameters(_dest awss3.IBucketNotificationDestination, _filters *[]*awss3.NotificationKeyFilter) error {
	return nil
}

func (c *jsiiProxy_CloudfrontLogsBucket) validateAddObjectRemovedNotificationParameters(_dest awss3.IBucketNotificationDestination, _filters *[]*awss3.NotificationKeyFilter) error {
	return nil
}

func (c *jsiiProxy_CloudfrontLogsBucket) validateAddToResourcePolicyParameters(permission awsiam.PolicyStatement) error {
	return nil
}

func (c *jsiiProxy_CloudfrontLogsBucket) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (c *jsiiProxy_CloudfrontLogsBucket) validateArnForObjectsParameters(_keyPattern *string) error {
	return nil
}

func (c *jsiiProxy_CloudfrontLogsBucket) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (c *jsiiProxy_CloudfrontLogsBucket) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func (c *jsiiProxy_CloudfrontLogsBucket) validateGrantDeleteParameters(_identity awsiam.IGrantable) error {
	return nil
}

func (c *jsiiProxy_CloudfrontLogsBucket) validateGrantPutParameters(_identity awsiam.IGrantable) error {
	return nil
}

func (c *jsiiProxy_CloudfrontLogsBucket) validateGrantPutAclParameters(_identity awsiam.IGrantable) error {
	return nil
}

func (c *jsiiProxy_CloudfrontLogsBucket) validateGrantReadParameters(_identity awsiam.IGrantable) error {
	return nil
}

func (c *jsiiProxy_CloudfrontLogsBucket) validateGrantReadWriteParameters(_identity awsiam.IGrantable) error {
	return nil
}

func (c *jsiiProxy_CloudfrontLogsBucket) validateGrantWriteParameters(_identity awsiam.IGrantable) error {
	return nil
}

func (c *jsiiProxy_CloudfrontLogsBucket) validateOnCloudTrailEventParameters(_id *string, _options *awss3.OnCloudTrailBucketEventOptions) error {
	return nil
}

func (c *jsiiProxy_CloudfrontLogsBucket) validateOnCloudTrailPutObjectParameters(_id *string, _options *awss3.OnCloudTrailBucketEventOptions) error {
	return nil
}

func (c *jsiiProxy_CloudfrontLogsBucket) validateOnCloudTrailWriteObjectParameters(_id *string, _options *awss3.OnCloudTrailBucketEventOptions) error {
	return nil
}

func (c *jsiiProxy_CloudfrontLogsBucket) validateTransferAccelerationUrlForObjectParameters(_options *awss3.TransferAccelerationUrlOptions) error {
	return nil
}

func (c *jsiiProxy_CloudfrontLogsBucket) validateVirtualHostedUrlForObjectParameters(_options *awss3.VirtualHostedStyleUrlOptions) error {
	return nil
}

func validateCloudfrontLogsBucket_IsConstructParameters(x interface{}) error {
	return nil
}

func validateCloudfrontLogsBucket_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCloudfrontLogsBucket_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewCloudfrontLogsBucketParameters(scope constructs.Construct, id *string, props *CloudfrontLogsBucketProps) error {
	return nil
}

