//go:build !no_runtime_type_checking

package s3buckets

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
	"github.com/aws/constructs-go/constructs/v10"
)

func (r *jsiiProxy_RawBucket) validateAddEventNotificationParameters(_event awss3.EventType, _dest awss3.IBucketNotificationDestination, _filters *[]*awss3.NotificationKeyFilter) error {
	if _event == "" {
		return fmt.Errorf("parameter _event is required, but nil was provided")
	}

	if _dest == nil {
		return fmt.Errorf("parameter _dest is required, but nil was provided")
	}

	for idx_263495, v := range *_filters {
		if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter _filters[%#v]", idx_263495) }); err != nil {
			return err
		}
	}

	return nil
}

func (r *jsiiProxy_RawBucket) validateAddObjectCreatedNotificationParameters(_dest awss3.IBucketNotificationDestination, _filters *[]*awss3.NotificationKeyFilter) error {
	if _dest == nil {
		return fmt.Errorf("parameter _dest is required, but nil was provided")
	}

	for idx_263495, v := range *_filters {
		if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter _filters[%#v]", idx_263495) }); err != nil {
			return err
		}
	}

	return nil
}

func (r *jsiiProxy_RawBucket) validateAddObjectRemovedNotificationParameters(_dest awss3.IBucketNotificationDestination, _filters *[]*awss3.NotificationKeyFilter) error {
	if _dest == nil {
		return fmt.Errorf("parameter _dest is required, but nil was provided")
	}

	for idx_263495, v := range *_filters {
		if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter _filters[%#v]", idx_263495) }); err != nil {
			return err
		}
	}

	return nil
}

func (r *jsiiProxy_RawBucket) validateAddToResourcePolicyParameters(permission awsiam.PolicyStatement) error {
	if permission == nil {
		return fmt.Errorf("parameter permission is required, but nil was provided")
	}

	return nil
}

func (r *jsiiProxy_RawBucket) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	if policy == "" {
		return fmt.Errorf("parameter policy is required, but nil was provided")
	}

	return nil
}

func (r *jsiiProxy_RawBucket) validateArnForObjectsParameters(_keyPattern *string) error {
	if _keyPattern == nil {
		return fmt.Errorf("parameter _keyPattern is required, but nil was provided")
	}

	return nil
}

func (r *jsiiProxy_RawBucket) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	if arnAttr == nil {
		return fmt.Errorf("parameter arnAttr is required, but nil was provided")
	}

	if arnComponents == nil {
		return fmt.Errorf("parameter arnComponents is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(arnComponents, func() string { return "parameter arnComponents" }); err != nil {
		return err
	}

	return nil
}

func (r *jsiiProxy_RawBucket) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	if nameAttr == nil {
		return fmt.Errorf("parameter nameAttr is required, but nil was provided")
	}

	return nil
}

func (r *jsiiProxy_RawBucket) validateGrantDeleteParameters(_identity awsiam.IGrantable) error {
	if _identity == nil {
		return fmt.Errorf("parameter _identity is required, but nil was provided")
	}

	return nil
}

func (r *jsiiProxy_RawBucket) validateGrantPutParameters(_identity awsiam.IGrantable) error {
	if _identity == nil {
		return fmt.Errorf("parameter _identity is required, but nil was provided")
	}

	return nil
}

func (r *jsiiProxy_RawBucket) validateGrantPutAclParameters(_identity awsiam.IGrantable) error {
	if _identity == nil {
		return fmt.Errorf("parameter _identity is required, but nil was provided")
	}

	return nil
}

func (r *jsiiProxy_RawBucket) validateGrantReadParameters(_identity awsiam.IGrantable) error {
	if _identity == nil {
		return fmt.Errorf("parameter _identity is required, but nil was provided")
	}

	return nil
}

func (r *jsiiProxy_RawBucket) validateGrantReadWriteParameters(_identity awsiam.IGrantable) error {
	if _identity == nil {
		return fmt.Errorf("parameter _identity is required, but nil was provided")
	}

	return nil
}

func (r *jsiiProxy_RawBucket) validateGrantWriteParameters(_identity awsiam.IGrantable) error {
	if _identity == nil {
		return fmt.Errorf("parameter _identity is required, but nil was provided")
	}

	return nil
}

func (r *jsiiProxy_RawBucket) validateOnCloudTrailEventParameters(_id *string, _options *awss3.OnCloudTrailBucketEventOptions) error {
	if _id == nil {
		return fmt.Errorf("parameter _id is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(_options, func() string { return "parameter _options" }); err != nil {
		return err
	}

	return nil
}

func (r *jsiiProxy_RawBucket) validateOnCloudTrailPutObjectParameters(_id *string, _options *awss3.OnCloudTrailBucketEventOptions) error {
	if _id == nil {
		return fmt.Errorf("parameter _id is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(_options, func() string { return "parameter _options" }); err != nil {
		return err
	}

	return nil
}

func (r *jsiiProxy_RawBucket) validateOnCloudTrailWriteObjectParameters(_id *string, _options *awss3.OnCloudTrailBucketEventOptions) error {
	if _id == nil {
		return fmt.Errorf("parameter _id is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(_options, func() string { return "parameter _options" }); err != nil {
		return err
	}

	return nil
}

func (r *jsiiProxy_RawBucket) validateTransferAccelerationUrlForObjectParameters(_options *awss3.TransferAccelerationUrlOptions) error {
	if err := _jsii_.ValidateStruct(_options, func() string { return "parameter _options" }); err != nil {
		return err
	}

	return nil
}

func (r *jsiiProxy_RawBucket) validateVirtualHostedUrlForObjectParameters(_options *awss3.VirtualHostedStyleUrlOptions) error {
	if err := _jsii_.ValidateStruct(_options, func() string { return "parameter _options" }); err != nil {
		return err
	}

	return nil
}

func validateRawBucket_IsConstructParameters(x interface{}) error {
	if x == nil {
		return fmt.Errorf("parameter x is required, but nil was provided")
	}

	return nil
}

func validateRawBucket_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	if construct == nil {
		return fmt.Errorf("parameter construct is required, but nil was provided")
	}

	return nil
}

func validateRawBucket_IsResourceParameters(construct constructs.IConstruct) error {
	if construct == nil {
		return fmt.Errorf("parameter construct is required, but nil was provided")
	}

	return nil
}

func validateNewRawBucketParameters(scope constructs.Construct, id *string, props *RawBucketProps) error {
	if scope == nil {
		return fmt.Errorf("parameter scope is required, but nil was provided")
	}

	if id == nil {
		return fmt.Errorf("parameter id is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

