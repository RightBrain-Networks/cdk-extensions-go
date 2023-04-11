//go:build no_runtime_type_checking

package s3buckets

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_AlbLogsBucket) validateAddEventNotificationParameters(_event awss3.EventType, _dest awss3.IBucketNotificationDestination, _filters *[]*awss3.NotificationKeyFilter) error {
	return nil
}

func (a *jsiiProxy_AlbLogsBucket) validateAddObjectCreatedNotificationParameters(_dest awss3.IBucketNotificationDestination, _filters *[]*awss3.NotificationKeyFilter) error {
	return nil
}

func (a *jsiiProxy_AlbLogsBucket) validateAddObjectRemovedNotificationParameters(_dest awss3.IBucketNotificationDestination, _filters *[]*awss3.NotificationKeyFilter) error {
	return nil
}

func (a *jsiiProxy_AlbLogsBucket) validateAddToResourcePolicyParameters(permission awsiam.PolicyStatement) error {
	return nil
}

func (a *jsiiProxy_AlbLogsBucket) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (a *jsiiProxy_AlbLogsBucket) validateArnForObjectsParameters(_keyPattern *string) error {
	return nil
}

func (a *jsiiProxy_AlbLogsBucket) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (a *jsiiProxy_AlbLogsBucket) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func (a *jsiiProxy_AlbLogsBucket) validateGrantDeleteParameters(_identity awsiam.IGrantable) error {
	return nil
}

func (a *jsiiProxy_AlbLogsBucket) validateGrantPutParameters(_identity awsiam.IGrantable) error {
	return nil
}

func (a *jsiiProxy_AlbLogsBucket) validateGrantPutAclParameters(_identity awsiam.IGrantable) error {
	return nil
}

func (a *jsiiProxy_AlbLogsBucket) validateGrantReadParameters(_identity awsiam.IGrantable) error {
	return nil
}

func (a *jsiiProxy_AlbLogsBucket) validateGrantReadWriteParameters(_identity awsiam.IGrantable) error {
	return nil
}

func (a *jsiiProxy_AlbLogsBucket) validateGrantWriteParameters(_identity awsiam.IGrantable) error {
	return nil
}

func (a *jsiiProxy_AlbLogsBucket) validateOnCloudTrailEventParameters(_id *string, _options *awss3.OnCloudTrailBucketEventOptions) error {
	return nil
}

func (a *jsiiProxy_AlbLogsBucket) validateOnCloudTrailPutObjectParameters(_id *string, _options *awss3.OnCloudTrailBucketEventOptions) error {
	return nil
}

func (a *jsiiProxy_AlbLogsBucket) validateOnCloudTrailWriteObjectParameters(_id *string, _options *awss3.OnCloudTrailBucketEventOptions) error {
	return nil
}

func (a *jsiiProxy_AlbLogsBucket) validateTransferAccelerationUrlForObjectParameters(_options *awss3.TransferAccelerationUrlOptions) error {
	return nil
}

func (a *jsiiProxy_AlbLogsBucket) validateVirtualHostedUrlForObjectParameters(_options *awss3.VirtualHostedStyleUrlOptions) error {
	return nil
}

func validateAlbLogsBucket_IsConstructParameters(x interface{}) error {
	return nil
}

func validateAlbLogsBucket_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateAlbLogsBucket_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewAlbLogsBucketParameters(scope constructs.Construct, id *string, props *AlbLogsBucketProps) error {
	return nil
}

