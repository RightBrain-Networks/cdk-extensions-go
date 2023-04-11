//go:build !no_runtime_type_checking

package ec2

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/constructs-go/constructs/v10"
)

func (i *jsiiProxy_IpamResourceDiscovery) validateAddIpamParameters(id *string, options *IpamProps) error {
	if id == nil {
		return fmt.Errorf("parameter id is required, but nil was provided")
	}

	if options == nil {
		return fmt.Errorf("parameter options is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func (i *jsiiProxy_IpamResourceDiscovery) validateAddRegionParameters(region *string) error {
	if region == nil {
		return fmt.Errorf("parameter region is required, but nil was provided")
	}

	return nil
}

func (i *jsiiProxy_IpamResourceDiscovery) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	if policy == "" {
		return fmt.Errorf("parameter policy is required, but nil was provided")
	}

	return nil
}

func (i *jsiiProxy_IpamResourceDiscovery) validateAssociateIpamParameters(ipam IIpam) error {
	if ipam == nil {
		return fmt.Errorf("parameter ipam is required, but nil was provided")
	}

	return nil
}

func (i *jsiiProxy_IpamResourceDiscovery) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
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

func (i *jsiiProxy_IpamResourceDiscovery) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	if nameAttr == nil {
		return fmt.Errorf("parameter nameAttr is required, but nil was provided")
	}

	return nil
}

func validateIpamResourceDiscovery_FromIpamResourceDiscoveryArnParameters(scope constructs.IConstruct, id *string, ipamResourceDiscoveryArn *string) error {
	if scope == nil {
		return fmt.Errorf("parameter scope is required, but nil was provided")
	}

	if id == nil {
		return fmt.Errorf("parameter id is required, but nil was provided")
	}

	if ipamResourceDiscoveryArn == nil {
		return fmt.Errorf("parameter ipamResourceDiscoveryArn is required, but nil was provided")
	}

	return nil
}

func validateIpamResourceDiscovery_FromIpamResourceDiscoveryAttributesParameters(scope constructs.IConstruct, id *string, attrs *IpamResourceDiscoveryAttributes) error {
	if scope == nil {
		return fmt.Errorf("parameter scope is required, but nil was provided")
	}

	if id == nil {
		return fmt.Errorf("parameter id is required, but nil was provided")
	}

	if attrs == nil {
		return fmt.Errorf("parameter attrs is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(attrs, func() string { return "parameter attrs" }); err != nil {
		return err
	}

	return nil
}

func validateIpamResourceDiscovery_FromIpamResourceDiscoveryIdParameters(scope constructs.IConstruct, id *string, ipamResourceDiscoveryId *string) error {
	if scope == nil {
		return fmt.Errorf("parameter scope is required, but nil was provided")
	}

	if id == nil {
		return fmt.Errorf("parameter id is required, but nil was provided")
	}

	if ipamResourceDiscoveryId == nil {
		return fmt.Errorf("parameter ipamResourceDiscoveryId is required, but nil was provided")
	}

	return nil
}

func validateIpamResourceDiscovery_IsConstructParameters(x interface{}) error {
	if x == nil {
		return fmt.Errorf("parameter x is required, but nil was provided")
	}

	return nil
}

func validateIpamResourceDiscovery_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	if construct == nil {
		return fmt.Errorf("parameter construct is required, but nil was provided")
	}

	return nil
}

func validateIpamResourceDiscovery_IsResourceParameters(construct constructs.IConstruct) error {
	if construct == nil {
		return fmt.Errorf("parameter construct is required, but nil was provided")
	}

	return nil
}

func validateNewIpamResourceDiscoveryParameters(scope constructs.IConstruct, id *string, props *IpamResourceDiscoveryProps) error {
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

