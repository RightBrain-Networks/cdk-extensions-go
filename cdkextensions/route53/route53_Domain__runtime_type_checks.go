//go:build !no_runtime_type_checking
// +build !no_runtime_type_checking

package route53

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsroute53"
	"github.com/aws/constructs-go/constructs/v10"
)

func (d *jsiiProxy_Domain) validateVisitParameters(node constructs.IConstruct) error {
	if node == nil {
		return fmt.Errorf("parameter node is required, but nil was provided")
	}

	return nil
}

func validateNewDomainParameters(zone awsroute53.IHostedZone, isPublic *bool, options *DomainOptions) error {
	if zone == nil {
		return fmt.Errorf("parameter zone is required, but nil was provided")
	}

	if isPublic == nil {
		return fmt.Errorf("parameter isPublic is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

