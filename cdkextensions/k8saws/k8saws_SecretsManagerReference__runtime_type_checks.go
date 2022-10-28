//go:build !no_runtime_type_checking
// +build !no_runtime_type_checking

package k8saws

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awssecretsmanager"
	"github.com/aws/constructs-go/constructs/v10"
)

func (s *jsiiProxy_SecretsManagerReference) validateAddFieldMappingParameters(field *SecretFieldReference) error {
	if field == nil {
		return fmt.Errorf("parameter field is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(field, func() string { return "parameter field" }); err != nil {
		return err
	}

	return nil
}

func (s *jsiiProxy_SecretsManagerReference) validateBindParameters(_scope constructs.IConstruct) error {
	if _scope == nil {
		return fmt.Errorf("parameter _scope is required, but nil was provided")
	}

	return nil
}

func validateNewSecretsManagerReferenceParameters(secret awssecretsmanager.ISecret, options *SecretsManagerReferenceOptions) error {
	if secret == nil {
		return fmt.Errorf("parameter secret is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

