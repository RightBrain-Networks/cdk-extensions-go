//go:build !no_runtime_type_checking

package ec2

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
)

func (t *jsiiProxy_TieredSubnets) validateAllocateSubnetsCidrParameters(input *awsec2.AllocateCidrRequest) error {
	if input == nil {
		return fmt.Errorf("parameter input is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(input, func() string { return "parameter input" }); err != nil {
		return err
	}

	return nil
}

func validateNewTieredSubnetsParameters(options *TieredSubnetsOptions) error {
	if options == nil {
		return fmt.Errorf("parameter options is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

