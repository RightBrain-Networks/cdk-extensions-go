//go:build !no_runtime_type_checking

package ec2

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/constructs-go/constructs/v10"
)

func validateIpamScope_FromIpamScopeArnParameters(scope constructs.IConstruct, id *string, ipamScopeArn *string) error {
	if scope == nil {
		return fmt.Errorf("parameter scope is required, but nil was provided")
	}

	if id == nil {
		return fmt.Errorf("parameter id is required, but nil was provided")
	}

	if ipamScopeArn == nil {
		return fmt.Errorf("parameter ipamScopeArn is required, but nil was provided")
	}

	return nil
}

func validateIpamScope_FromIpamScopeAttributesParameters(scope constructs.IConstruct, id *string, attrs *IpamScopeAttributes) error {
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

func validateIpamScope_FromIpamScopeIdParameters(scope constructs.IConstruct, id *string, ipamScopeId *string) error {
	if scope == nil {
		return fmt.Errorf("parameter scope is required, but nil was provided")
	}

	if id == nil {
		return fmt.Errorf("parameter id is required, but nil was provided")
	}

	if ipamScopeId == nil {
		return fmt.Errorf("parameter ipamScopeId is required, but nil was provided")
	}

	return nil
}

