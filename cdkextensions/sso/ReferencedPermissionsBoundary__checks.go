//go:build !no_runtime_type_checking

package sso

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/constructs-go/constructs/v10"
)

func (r *jsiiProxy_ReferencedPermissionsBoundary) validateBindParameters(_scope constructs.IConstruct) error {
	if _scope == nil {
		return fmt.Errorf("parameter _scope is required, but nil was provided")
	}

	return nil
}

func validateReferencedPermissionsBoundary_FromManagedPolicyParameters(policy awsiam.IManagedPolicy) error {
	if policy == nil {
		return fmt.Errorf("parameter policy is required, but nil was provided")
	}

	return nil
}

func validateReferencedPermissionsBoundary_FromReferenceParameters(options *ReferenceOptions) error {
	if options == nil {
		return fmt.Errorf("parameter options is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func validateNewReferencedPermissionsBoundaryParameters(options *ReferenceOptions) error {
	if options == nil {
		return fmt.Errorf("parameter options is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

