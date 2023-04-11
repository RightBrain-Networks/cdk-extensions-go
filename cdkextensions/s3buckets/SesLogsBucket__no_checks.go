//go:build no_runtime_type_checking

package s3buckets

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_SesLogsBucket) validateAddEventNotificationParameters(_event awss3.EventType, _dest awss3.IBucketNotificationDestination, _filters *[]*awss3.NotificationKeyFilter) error {
	return nil
}

func (s *jsiiProxy_SesLogsBucket) validateAddObjectCreatedNotificationParameters(_dest awss3.IBucketNotificationDestination, _filters *[]*awss3.NotificationKeyFilter) error {
	return nil
}

func (s *jsiiProxy_SesLogsBucket) validateAddObjectRemovedNotificationParameters(_dest awss3.IBucketNotificationDestination, _filters *[]*awss3.NotificationKeyFilter) error {
	return nil
}

func (s *jsiiProxy_SesLogsBucket) validateAddToResourcePolicyParameters(permission awsiam.PolicyStatement) error {
	return nil
}

func (s *jsiiProxy_SesLogsBucket) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (s *jsiiProxy_SesLogsBucket) validateArnForObjectsParameters(_keyPattern *string) error {
	return nil
}

func (s *jsiiProxy_SesLogsBucket) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (s *jsiiProxy_SesLogsBucket) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func (s *jsiiProxy_SesLogsBucket) validateGrantDeleteParameters(_identity awsiam.IGrantable) error {
	return nil
}

func (s *jsiiProxy_SesLogsBucket) validateGrantPutParameters(_identity awsiam.IGrantable) error {
	return nil
}

func (s *jsiiProxy_SesLogsBucket) validateGrantPutAclParameters(_identity awsiam.IGrantable) error {
	return nil
}

func (s *jsiiProxy_SesLogsBucket) validateGrantReadParameters(_identity awsiam.IGrantable) error {
	return nil
}

func (s *jsiiProxy_SesLogsBucket) validateGrantReadWriteParameters(_identity awsiam.IGrantable) error {
	return nil
}

func (s *jsiiProxy_SesLogsBucket) validateGrantWriteParameters(_identity awsiam.IGrantable) error {
	return nil
}

func (s *jsiiProxy_SesLogsBucket) validateOnCloudTrailEventParameters(_id *string, _options *awss3.OnCloudTrailBucketEventOptions) error {
	return nil
}

func (s *jsiiProxy_SesLogsBucket) validateOnCloudTrailPutObjectParameters(_id *string, _options *awss3.OnCloudTrailBucketEventOptions) error {
	return nil
}

func (s *jsiiProxy_SesLogsBucket) validateOnCloudTrailWriteObjectParameters(_id *string, _options *awss3.OnCloudTrailBucketEventOptions) error {
	return nil
}

func (s *jsiiProxy_SesLogsBucket) validateTransferAccelerationUrlForObjectParameters(_options *awss3.TransferAccelerationUrlOptions) error {
	return nil
}

func (s *jsiiProxy_SesLogsBucket) validateVirtualHostedUrlForObjectParameters(_options *awss3.VirtualHostedStyleUrlOptions) error {
	return nil
}

func validateSesLogsBucket_IsConstructParameters(x interface{}) error {
	return nil
}

func validateSesLogsBucket_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateSesLogsBucket_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewSesLogsBucketParameters(scope constructs.Construct, id *string, props *SesLogsBucketProps) error {
	return nil
}

