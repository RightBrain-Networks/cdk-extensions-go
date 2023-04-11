//go:build !no_runtime_type_checking

package k8saws

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsssm"
	"github.com/aws/constructs-go/constructs/v10"
)

func (s *jsiiProxy_SsmParameterReference) validateAddFieldMappingParameters(field *SecretFieldReference) error {
	if field == nil {
		return fmt.Errorf("parameter field is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(field, func() string { return "parameter field" }); err != nil {
		return err
	}

	return nil
}

func (s *jsiiProxy_SsmParameterReference) validateBindParameters(_scope constructs.IConstruct) error {
	if _scope == nil {
		return fmt.Errorf("parameter _scope is required, but nil was provided")
	}

	return nil
}

func validateNewSsmParameterReferenceParameters(parameter awsssm.IParameter, options *SsmParameterReferenceOptions) error {
	if parameter == nil {
		return fmt.Errorf("parameter parameter is required, but nil was provided")
	}

	if options == nil {
		return fmt.Errorf("parameter options is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

