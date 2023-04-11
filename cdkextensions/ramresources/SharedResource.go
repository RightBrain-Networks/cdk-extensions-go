package ramresources

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/vibe-io/cdk-extensions-go/cdkextensions/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2/awscodebuild"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/vibe-io/cdk-extensions-go/cdkextensions/ec2"
	"github.com/vibe-io/cdk-extensions-go/cdkextensions/ram"
	"github.com/vibe-io/cdk-extensions-go/cdkextensions/ramresources/internal"
)

type SharedResource interface {
	ram.ISharable
	// Configures resource sharing for the associated resource.
	Share(_scope constructs.IConstruct) *string
}

// The jsii proxy struct for SharedResource
type jsiiProxy_SharedResource struct {
	internal.Type__ramISharable
}

func SharedResource_FromArn(arn *string) SharedResource {
	_init_.Initialize()

	if err := validateSharedResource_FromArnParameters(arn); err != nil {
		panic(err)
	}
	var returns SharedResource

	_jsii_.StaticInvoke(
		"cdk-extensions.ram_resources.SharedResource",
		"fromArn",
		[]interface{}{arn},
		&returns,
	)

	return returns
}

func SharedResource_FromIpamPool(pool ec2.IIpamPool) SharedResource {
	_init_.Initialize()

	if err := validateSharedResource_FromIpamPoolParameters(pool); err != nil {
		panic(err)
	}
	var returns SharedResource

	_jsii_.StaticInvoke(
		"cdk-extensions.ram_resources.SharedResource",
		"fromIpamPool",
		[]interface{}{pool},
		&returns,
	)

	return returns
}

func SharedResource_FromProject(project awscodebuild.IProject) SharedResource {
	_init_.Initialize()

	if err := validateSharedResource_FromProjectParameters(project); err != nil {
		panic(err)
	}
	var returns SharedResource

	_jsii_.StaticInvoke(
		"cdk-extensions.ram_resources.SharedResource",
		"fromProject",
		[]interface{}{project},
		&returns,
	)

	return returns
}

func SharedResource_FromSubnet(subnet awsec2.ISubnet) SharedResource {
	_init_.Initialize()

	if err := validateSharedResource_FromSubnetParameters(subnet); err != nil {
		panic(err)
	}
	var returns SharedResource

	_jsii_.StaticInvoke(
		"cdk-extensions.ram_resources.SharedResource",
		"fromSubnet",
		[]interface{}{subnet},
		&returns,
	)

	return returns
}

func SharedResource_FromTransitGateway(transitGateway ec2.ITransitGateway) SharedResource {
	_init_.Initialize()

	if err := validateSharedResource_FromTransitGatewayParameters(transitGateway); err != nil {
		panic(err)
	}
	var returns SharedResource

	_jsii_.StaticInvoke(
		"cdk-extensions.ram_resources.SharedResource",
		"fromTransitGateway",
		[]interface{}{transitGateway},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SharedResource) Share(_scope constructs.IConstruct) *string {
	if err := s.validateShareParameters(_scope); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		s,
		"share",
		[]interface{}{_scope},
		&returns,
	)

	return returns
}

