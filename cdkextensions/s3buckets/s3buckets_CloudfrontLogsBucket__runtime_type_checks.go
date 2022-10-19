//go:build !no_runtime_type_checking
// +build !no_runtime_type_checking

package s3buckets

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
	"github.com/aws/constructs-go/constructs/v10"
)

func (c *jsiiProxy_CloudfrontLogsBucket) validateAddEventNotificationParameters(_event awss3.EventType, _dest awss3.IBucketNotificationDestination, _filters *[]*awss3.NotificationKeyFilter) error {
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

func (c *jsiiProxy_CloudfrontLogsBucket) validateAddObjectCreatedNotificationParameters(_dest awss3.IBucketNotificationDestination, _filters *[]*awss3.NotificationKeyFilter) error {
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

func (c *jsiiProxy_CloudfrontLogsBucket) validateAddObjectRemovedNotificationParameters(_dest awss3.IBucketNotificationDestination, _filters *[]*awss3.NotificationKeyFilter) error {
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

func (c *jsiiProxy_CloudfrontLogsBucket) validateAddToResourcePolicyParameters(permission awsiam.PolicyStatement) error {
	if permission == nil {
		return fmt.Errorf("parameter permission is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_CloudfrontLogsBucket) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	if policy == "" {
		return fmt.Errorf("parameter policy is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_CloudfrontLogsBucket) validateArnForObjectsParameters(_keyPattern *string) error {
	if _keyPattern == nil {
		return fmt.Errorf("parameter _keyPattern is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_CloudfrontLogsBucket) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
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

func (c *jsiiProxy_CloudfrontLogsBucket) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	if nameAttr == nil {
		return fmt.Errorf("parameter nameAttr is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_CloudfrontLogsBucket) validateGrantDeleteParameters(_identity awsiam.IGrantable) error {
	if _identity == nil {
		return fmt.Errorf("parameter _identity is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_CloudfrontLogsBucket) validateGrantPutParameters(_identity awsiam.IGrantable) error {
	if _identity == nil {
		return fmt.Errorf("parameter _identity is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_CloudfrontLogsBucket) validateGrantPutAclParameters(_identity awsiam.IGrantable) error {
	if _identity == nil {
		return fmt.Errorf("parameter _identity is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_CloudfrontLogsBucket) validateGrantReadParameters(_identity awsiam.IGrantable) error {
	if _identity == nil {
		return fmt.Errorf("parameter _identity is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_CloudfrontLogsBucket) validateGrantReadWriteParameters(_identity awsiam.IGrantable) error {
	if _identity == nil {
		return fmt.Errorf("parameter _identity is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_CloudfrontLogsBucket) validateGrantWriteParameters(_identity awsiam.IGrantable) error {
	if _identity == nil {
		return fmt.Errorf("parameter _identity is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_CloudfrontLogsBucket) validateOnCloudTrailEventParameters(_id *string, _options *awss3.OnCloudTrailBucketEventOptions) error {
	if _id == nil {
		return fmt.Errorf("parameter _id is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(_options, func() string { return "parameter _options" }); err != nil {
		return err
	}

	return nil
}

func (c *jsiiProxy_CloudfrontLogsBucket) validateOnCloudTrailPutObjectParameters(_id *string, _options *awss3.OnCloudTrailBucketEventOptions) error {
	if _id == nil {
		return fmt.Errorf("parameter _id is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(_options, func() string { return "parameter _options" }); err != nil {
		return err
	}

	return nil
}

func (c *jsiiProxy_CloudfrontLogsBucket) validateOnCloudTrailWriteObjectParameters(_id *string, _options *awss3.OnCloudTrailBucketEventOptions) error {
	if _id == nil {
		return fmt.Errorf("parameter _id is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(_options, func() string { return "parameter _options" }); err != nil {
		return err
	}

	return nil
}

func (c *jsiiProxy_CloudfrontLogsBucket) validateTransferAccelerationUrlForObjectParameters(_options *awss3.TransferAccelerationUrlOptions) error {
	if err := _jsii_.ValidateStruct(_options, func() string { return "parameter _options" }); err != nil {
		return err
	}

	return nil
}

func (c *jsiiProxy_CloudfrontLogsBucket) validateVirtualHostedUrlForObjectParameters(_options *awss3.VirtualHostedStyleUrlOptions) error {
	if err := _jsii_.ValidateStruct(_options, func() string { return "parameter _options" }); err != nil {
		return err
	}

	return nil
}

func validateCloudfrontLogsBucket_IsConstructParameters(x interface{}) error {
	if x == nil {
		return fmt.Errorf("parameter x is required, but nil was provided")
	}

	return nil
}

func validateCloudfrontLogsBucket_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	if construct == nil {
		return fmt.Errorf("parameter construct is required, but nil was provided")
	}

	return nil
}

func validateCloudfrontLogsBucket_IsResourceParameters(construct constructs.IConstruct) error {
	if construct == nil {
		return fmt.Errorf("parameter construct is required, but nil was provided")
	}

	return nil
}

func validateNewCloudfrontLogsBucketParameters(scope constructs.Construct, id *string, props *CloudfrontLogsBucketProps) error {
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

