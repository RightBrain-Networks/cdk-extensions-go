//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package s3buckets

// Building without runtime type checking enabled, so all the below just return nil

func (r *jsiiProxy_RawBucket) validateAddEventNotificationParameters(_event awss3.EventType, _dest awss3.IBucketNotificationDestination, _filters *[]*awss3.NotificationKeyFilter) error {
	return nil
}

func (r *jsiiProxy_RawBucket) validateAddObjectCreatedNotificationParameters(_dest awss3.IBucketNotificationDestination, _filters *[]*awss3.NotificationKeyFilter) error {
	return nil
}

func (r *jsiiProxy_RawBucket) validateAddObjectRemovedNotificationParameters(_dest awss3.IBucketNotificationDestination, _filters *[]*awss3.NotificationKeyFilter) error {
	return nil
}

func (r *jsiiProxy_RawBucket) validateAddToResourcePolicyParameters(permission awsiam.PolicyStatement) error {
	return nil
}

func (r *jsiiProxy_RawBucket) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (r *jsiiProxy_RawBucket) validateArnForObjectsParameters(_keyPattern *string) error {
	return nil
}

func (r *jsiiProxy_RawBucket) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (r *jsiiProxy_RawBucket) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func (r *jsiiProxy_RawBucket) validateGrantDeleteParameters(_identity awsiam.IGrantable) error {
	return nil
}

func (r *jsiiProxy_RawBucket) validateGrantPutParameters(_identity awsiam.IGrantable) error {
	return nil
}

func (r *jsiiProxy_RawBucket) validateGrantPutAclParameters(_identity awsiam.IGrantable) error {
	return nil
}

func (r *jsiiProxy_RawBucket) validateGrantReadParameters(_identity awsiam.IGrantable) error {
	return nil
}

func (r *jsiiProxy_RawBucket) validateGrantReadWriteParameters(_identity awsiam.IGrantable) error {
	return nil
}

func (r *jsiiProxy_RawBucket) validateGrantWriteParameters(_identity awsiam.IGrantable) error {
	return nil
}

func (r *jsiiProxy_RawBucket) validateOnCloudTrailEventParameters(_id *string, _options *awss3.OnCloudTrailBucketEventOptions) error {
	return nil
}

func (r *jsiiProxy_RawBucket) validateOnCloudTrailPutObjectParameters(_id *string, _options *awss3.OnCloudTrailBucketEventOptions) error {
	return nil
}

func (r *jsiiProxy_RawBucket) validateOnCloudTrailWriteObjectParameters(_id *string, _options *awss3.OnCloudTrailBucketEventOptions) error {
	return nil
}

func (r *jsiiProxy_RawBucket) validateTransferAccelerationUrlForObjectParameters(_options *awss3.TransferAccelerationUrlOptions) error {
	return nil
}

func (r *jsiiProxy_RawBucket) validateVirtualHostedUrlForObjectParameters(_options *awss3.VirtualHostedStyleUrlOptions) error {
	return nil
}

func validateRawBucket_IsConstructParameters(x interface{}) error {
	return nil
}

func validateRawBucket_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateRawBucket_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewRawBucketParameters(scope constructs.Construct, id *string, props *RawBucketProps) error {
	return nil
}

