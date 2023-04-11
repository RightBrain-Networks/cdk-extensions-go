//go:build !no_runtime_type_checking

package ram

import (
	"fmt"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/constructs-go/constructs/v10"
)

func (s *jsiiProxy_SharedPrincipal) validateBindParameters(_scope constructs.IConstruct) error {
	if _scope == nil {
		return fmt.Errorf("parameter _scope is required, but nil was provided")
	}

	return nil
}

func validateSharedPrincipal_FromAccountIdParameters(account *string) error {
	if account == nil {
		return fmt.Errorf("parameter account is required, but nil was provided")
	}

	return nil
}

func validateSharedPrincipal_FromConstructParameters(construct constructs.IConstruct) error {
	if construct == nil {
		return fmt.Errorf("parameter construct is required, but nil was provided")
	}

	return nil
}

func validateSharedPrincipal_FromOrganizationalUnitArnParameters(arn *string) error {
	if arn == nil {
		return fmt.Errorf("parameter arn is required, but nil was provided")
	}

	return nil
}

func validateSharedPrincipal_FromOrganizationArnParameters(arn *string) error {
	if arn == nil {
		return fmt.Errorf("parameter arn is required, but nil was provided")
	}

	return nil
}

func validateSharedPrincipal_FromRoleParameters(role awsiam.IRole) error {
	if role == nil {
		return fmt.Errorf("parameter role is required, but nil was provided")
	}

	return nil
}

func validateSharedPrincipal_FromStageParameters(stage awscdk.Stage) error {
	if stage == nil {
		return fmt.Errorf("parameter stage is required, but nil was provided")
	}

	return nil
}

func validateSharedPrincipal_FromUserParameters(user awsiam.IUser) error {
	if user == nil {
		return fmt.Errorf("parameter user is required, but nil was provided")
	}

	return nil
}

