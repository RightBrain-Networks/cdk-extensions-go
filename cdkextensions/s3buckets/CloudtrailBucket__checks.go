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

func (c *jsiiProxy_CloudtrailBucket) validateAddEventNotificationParameters(_event awss3.EventType, _dest awss3.IBucketNotificationDestination, _filters *[]*awss3.NotificationKeyFilter) error {
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

func (c *jsiiProxy_CloudtrailBucket) validateAddObjectCreatedNotificationParameters(_dest awss3.IBucketNotificationDestination, _filters *[]*awss3.NotificationKeyFilter) error {
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

func (c *jsiiProxy_CloudtrailBucket) validateAddObjectRemovedNotificationParameters(_dest awss3.IBucketNotificationDestination, _filters *[]*awss3.NotificationKeyFilter) error {
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

func (c *jsiiProxy_CloudtrailBucket) validateAddToResourcePolicyParameters(permission awsiam.PolicyStatement) error {
	if permission == nil {
		return fmt.Errorf("parameter permission is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_CloudtrailBucket) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	if policy == "" {
		return fmt.Errorf("parameter policy is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_CloudtrailBucket) validateArnForObjectsParameters(_keyPattern *string) error {
	if _keyPattern == nil {
		return fmt.Errorf("parameter _keyPattern is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_CloudtrailBucket) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
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

func (c *jsiiProxy_CloudtrailBucket) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	if nameAttr == nil {
		return fmt.Errorf("parameter nameAttr is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_CloudtrailBucket) validateGrantDeleteParameters(_identity awsiam.IGrantable) error {
	if _identity == nil {
		return fmt.Errorf("parameter _identity is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_CloudtrailBucket) validateGrantPutParameters(_identity awsiam.IGrantable) error {
	if _identity == nil {
		return fmt.Errorf("parameter _identity is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_CloudtrailBucket) validateGrantPutAclParameters(_identity awsiam.IGrantable) error {
	if _identity == nil {
		return fmt.Errorf("parameter _identity is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_CloudtrailBucket) validateGrantReadParameters(_identity awsiam.IGrantable) error {
	if _identity == nil {
		return fmt.Errorf("parameter _identity is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_CloudtrailBucket) validateGrantReadWriteParameters(_identity awsiam.IGrantable) error {
	if _identity == nil {
		return fmt.Errorf("parameter _identity is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_CloudtrailBucket) validateGrantWriteParameters(_identity awsiam.IGrantable) error {
	if _identity == nil {
		return fmt.Errorf("parameter _identity is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_CloudtrailBucket) validateOnCloudTrailEventParameters(_id *string, _options *awss3.OnCloudTrailBucketEventOptions) error {
	if _id == nil {
		return fmt.Errorf("parameter _id is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(_options, func() string { return "parameter _options" }); err != nil {
		return err
	}

	return nil
}

func (c *jsiiProxy_CloudtrailBucket) validateOnCloudTrailPutObjectParameters(_id *string, _options *awss3.OnCloudTrailBucketEventOptions) error {
	if _id == nil {
		return fmt.Errorf("parameter _id is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(_options, func() string { return "parameter _options" }); err != nil {
		return err
	}

	return nil
}

func (c *jsiiProxy_CloudtrailBucket) validateOnCloudTrailWriteObjectParameters(_id *string, _options *awss3.OnCloudTrailBucketEventOptions) error {
	if _id == nil {
		return fmt.Errorf("parameter _id is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(_options, func() string { return "parameter _options" }); err != nil {
		return err
	}

	return nil
}

func (c *jsiiProxy_CloudtrailBucket) validateTransferAccelerationUrlForObjectParameters(_options *awss3.TransferAccelerationUrlOptions) error {
	if err := _jsii_.ValidateStruct(_options, func() string { return "parameter _options" }); err != nil {
		return err
	}

	return nil
}

func (c *jsiiProxy_CloudtrailBucket) validateVirtualHostedUrlForObjectParameters(_options *awss3.VirtualHostedStyleUrlOptions) error {
	if err := _jsii_.ValidateStruct(_options, func() string { return "parameter _options" }); err != nil {
		return err
	}

	return nil
}

func validateCloudtrailBucket_IsConstructParameters(x interface{}) error {
	if x == nil {
		return fmt.Errorf("parameter x is required, but nil was provided")
	}

	return nil
}

func validateCloudtrailBucket_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	if construct == nil {
		return fmt.Errorf("parameter construct is required, but nil was provided")
	}

	return nil
}

func validateCloudtrailBucket_IsResourceParameters(construct constructs.IConstruct) error {
	if construct == nil {
		return fmt.Errorf("parameter construct is required, but nil was provided")
	}

	return nil
}

func validateNewCloudtrailBucketParameters(scope constructs.Construct, id *string, props *CloudtrailBucketProps) error {
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

