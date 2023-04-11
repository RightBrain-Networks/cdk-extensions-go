//go:build !no_runtime_type_checking

package ramresources

import (
	"fmt"

	"github.com/aws/aws-cdk-go/awscdk/v2/awscodebuild"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/vibe-io/cdk-extensions-go/cdkextensions/ec2"
)

func (s *jsiiProxy_SharedResource) validateShareParameters(_scope constructs.IConstruct) error {
	if _scope == nil {
		return fmt.Errorf("parameter _scope is required, but nil was provided")
	}

	return nil
}

func validateSharedResource_FromArnParameters(arn *string) error {
	if arn == nil {
		return fmt.Errorf("parameter arn is required, but nil was provided")
	}

	return nil
}

func validateSharedResource_FromIpamPoolParameters(pool ec2.IIpamPool) error {
	if pool == nil {
		return fmt.Errorf("parameter pool is required, but nil was provided")
	}

	return nil
}

func validateSharedResource_FromProjectParameters(project awscodebuild.IProject) error {
	if project == nil {
		return fmt.Errorf("parameter project is required, but nil was provided")
	}

	return nil
}

func validateSharedResource_FromSubnetParameters(subnet awsec2.ISubnet) error {
	if subnet == nil {
		return fmt.Errorf("parameter subnet is required, but nil was provided")
	}

	return nil
}

func validateSharedResource_FromTransitGatewayParameters(transitGateway ec2.ITransitGateway) error {
	if transitGateway == nil {
		return fmt.Errorf("parameter transitGateway is required, but nil was provided")
	}

	return nil
}

