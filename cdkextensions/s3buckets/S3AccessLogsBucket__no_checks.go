//go:build no_runtime_type_checking

package s3buckets

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_S3AccessLogsBucket) validateAddEventNotificationParameters(_event awss3.EventType, _dest awss3.IBucketNotificationDestination, _filters *[]*awss3.NotificationKeyFilter) error {
	return nil
}

func (s *jsiiProxy_S3AccessLogsBucket) validateAddLoggingAspectParameters(scope constructs.IConstruct, options *LoggingAspectOptions) error {
	return nil
}

func (s *jsiiProxy_S3AccessLogsBucket) validateAddObjectCreatedNotificationParameters(_dest awss3.IBucketNotificationDestination, _filters *[]*awss3.NotificationKeyFilter) error {
	return nil
}

func (s *jsiiProxy_S3AccessLogsBucket) validateAddObjectRemovedNotificationParameters(_dest awss3.IBucketNotificationDestination, _filters *[]*awss3.NotificationKeyFilter) error {
	return nil
}

func (s *jsiiProxy_S3AccessLogsBucket) validateAddToResourcePolicyParameters(permission awsiam.PolicyStatement) error {
	return nil
}

func (s *jsiiProxy_S3AccessLogsBucket) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (s *jsiiProxy_S3AccessLogsBucket) validateArnForObjectsParameters(_keyPattern *string) error {
	return nil
}

func (s *jsiiProxy_S3AccessLogsBucket) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (s *jsiiProxy_S3AccessLogsBucket) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func (s *jsiiProxy_S3AccessLogsBucket) validateGrantDeleteParameters(_identity awsiam.IGrantable) error {
	return nil
}

func (s *jsiiProxy_S3AccessLogsBucket) validateGrantPutParameters(_identity awsiam.IGrantable) error {
	return nil
}

func (s *jsiiProxy_S3AccessLogsBucket) validateGrantPutAclParameters(_identity awsiam.IGrantable) error {
	return nil
}

func (s *jsiiProxy_S3AccessLogsBucket) validateGrantReadParameters(_identity awsiam.IGrantable) error {
	return nil
}

func (s *jsiiProxy_S3AccessLogsBucket) validateGrantReadWriteParameters(_identity awsiam.IGrantable) error {
	return nil
}

func (s *jsiiProxy_S3AccessLogsBucket) validateGrantWriteParameters(_identity awsiam.IGrantable) error {
	return nil
}

func (s *jsiiProxy_S3AccessLogsBucket) validateOnCloudTrailEventParameters(_id *string, _options *awss3.OnCloudTrailBucketEventOptions) error {
	return nil
}

func (s *jsiiProxy_S3AccessLogsBucket) validateOnCloudTrailPutObjectParameters(_id *string, _options *awss3.OnCloudTrailBucketEventOptions) error {
	return nil
}

func (s *jsiiProxy_S3AccessLogsBucket) validateOnCloudTrailWriteObjectParameters(_id *string, _options *awss3.OnCloudTrailBucketEventOptions) error {
	return nil
}

func (s *jsiiProxy_S3AccessLogsBucket) validateTransferAccelerationUrlForObjectParameters(_options *awss3.TransferAccelerationUrlOptions) error {
	return nil
}

func (s *jsiiProxy_S3AccessLogsBucket) validateVirtualHostedUrlForObjectParameters(_options *awss3.VirtualHostedStyleUrlOptions) error {
	return nil
}

func validateS3AccessLogsBucket_IsConstructParameters(x interface{}) error {
	return nil
}

func validateS3AccessLogsBucket_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateS3AccessLogsBucket_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewS3AccessLogsBucketParameters(scope constructs.Construct, id *string, props *S3AccessLogsBucketProps) error {
	return nil
}

