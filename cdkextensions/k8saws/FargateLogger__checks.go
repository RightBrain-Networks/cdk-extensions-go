//go:build !no_runtime_type_checking

package k8saws

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awseks"
	"github.com/aws/constructs-go/constructs/v10"
)

func (f *jsiiProxy_FargateLogger) validateAddFargateProfileParameters(profile awseks.FargateProfile) error {
	if profile == nil {
		return fmt.Errorf("parameter profile is required, but nil was provided")
	}

	return nil
}

func (f *jsiiProxy_FargateLogger) validateAddFilterParameters(filter IFluentBitFilterPlugin) error {
	if filter == nil {
		return fmt.Errorf("parameter filter is required, but nil was provided")
	}

	return nil
}

func (f *jsiiProxy_FargateLogger) validateAddOutputParameters(output IFluentBitOutputPlugin) error {
	if output == nil {
		return fmt.Errorf("parameter output is required, but nil was provided")
	}

	return nil
}

func (f *jsiiProxy_FargateLogger) validateAddParserParameters(parser IFluentBitParserPlugin) error {
	if parser == nil {
		return fmt.Errorf("parameter parser is required, but nil was provided")
	}

	return nil
}

func (f *jsiiProxy_FargateLogger) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	if policy == "" {
		return fmt.Errorf("parameter policy is required, but nil was provided")
	}

	return nil
}

func (f *jsiiProxy_FargateLogger) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
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

func (f *jsiiProxy_FargateLogger) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	if nameAttr == nil {
		return fmt.Errorf("parameter nameAttr is required, but nil was provided")
	}

	return nil
}

func validateFargateLogger_IsConstructParameters(x interface{}) error {
	if x == nil {
		return fmt.Errorf("parameter x is required, but nil was provided")
	}

	return nil
}

func validateFargateLogger_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	if construct == nil {
		return fmt.Errorf("parameter construct is required, but nil was provided")
	}

	return nil
}

func validateFargateLogger_IsResourceParameters(construct constructs.IConstruct) error {
	if construct == nil {
		return fmt.Errorf("parameter construct is required, but nil was provided")
	}

	return nil
}

func validateNewFargateLoggerParameters(scope constructs.Construct, id *string, props *FargateLoggerProps) error {
	if scope == nil {
		return fmt.Errorf("parameter scope is required, but nil was provided")
	}

	if id == nil {
		return fmt.Errorf("parameter id is required, but nil was provided")
	}

	if props == nil {
		return fmt.Errorf("parameter props is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

