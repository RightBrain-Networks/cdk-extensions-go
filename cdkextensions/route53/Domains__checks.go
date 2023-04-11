//go:build !no_runtime_type_checking

package route53

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsroute53"
	"github.com/aws/constructs-go/constructs/v10"
)

func (d *jsiiProxy_Domains) validateAddParameters(hostedZone awsroute53.IHostedZone, isPublic *bool, options *DomainOptions) error {
	if hostedZone == nil {
		return fmt.Errorf("parameter hostedZone is required, but nil was provided")
	}

	if isPublic == nil {
		return fmt.Errorf("parameter isPublic is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func validateDomains_OfParameters(scope constructs.IConstruct) error {
	if scope == nil {
		return fmt.Errorf("parameter scope is required, but nil was provided")
	}

	return nil
}

