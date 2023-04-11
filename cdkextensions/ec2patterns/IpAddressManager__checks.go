//go:build !no_runtime_type_checking

package ec2patterns

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/vibe-io/cdk-extensions-go/cdkextensions/ec2"
)

func (i *jsiiProxy_IpAddressManager) validateAddPrivatePoolParameters(name *string, options *AddPoolOptions) error {
	if name == nil {
		return fmt.Errorf("parameter name is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func (i *jsiiProxy_IpAddressManager) validateAddRegionParameters(region *string) error {
	if region == nil {
		return fmt.Errorf("parameter region is required, but nil was provided")
	}

	return nil
}

func (i *jsiiProxy_IpAddressManager) validateAddStagePoolParameters(scope constructs.IConstruct, parent ec2.IIpamPool) error {
	if scope == nil {
		return fmt.Errorf("parameter scope is required, but nil was provided")
	}

	if parent == nil {
		return fmt.Errorf("parameter parent is required, but nil was provided")
	}

	return nil
}

func (i *jsiiProxy_IpAddressManager) validateAllocatePrivateNetworkParameters(scope constructs.IConstruct, id *string, options *AllocatePrivateNetworkOptions) error {
	if scope == nil {
		return fmt.Errorf("parameter scope is required, but nil was provided")
	}

	if id == nil {
		return fmt.Errorf("parameter id is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func (i *jsiiProxy_IpAddressManager) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	if policy == "" {
		return fmt.Errorf("parameter policy is required, but nil was provided")
	}

	return nil
}

func (i *jsiiProxy_IpAddressManager) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
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

func (i *jsiiProxy_IpAddressManager) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	if nameAttr == nil {
		return fmt.Errorf("parameter nameAttr is required, but nil was provided")
	}

	return nil
}

func (i *jsiiProxy_IpAddressManager) validateRegisterAccountParameters(account *string, pool ec2.IIpamPool) error {
	if account == nil {
		return fmt.Errorf("parameter account is required, but nil was provided")
	}

	if pool == nil {
		return fmt.Errorf("parameter pool is required, but nil was provided")
	}

	return nil
}

func validateIpAddressManager_IsConstructParameters(x interface{}) error {
	if x == nil {
		return fmt.Errorf("parameter x is required, but nil was provided")
	}

	return nil
}

func validateIpAddressManager_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	if construct == nil {
		return fmt.Errorf("parameter construct is required, but nil was provided")
	}

	return nil
}

func validateIpAddressManager_IsResourceParameters(construct constructs.IConstruct) error {
	if construct == nil {
		return fmt.Errorf("parameter construct is required, but nil was provided")
	}

	return nil
}

func validateNewIpAddressManagerParameters(scope constructs.IConstruct, id *string, props *IpAddressManagerProps) error {
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

