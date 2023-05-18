package ec2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/vibe-io/cdk-extensions-go/cdkextensions/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/vibe-io/cdk-extensions-go/cdkextensions/ec2/internal"
)

type TransitGatewayNatProvider interface {
	awsec2.NatProvider
	// Return list of gateways spawned by the provider.
	ConfiguredGateways() *[]*awsec2.GatewayConfig
	TransitGateway() ITransitGateway
	TransitGatewayAttachment() TransitGatewayAttachment
	// Called by the VPC to configure NAT.
	//
	// Don't call this directly, the VPC will call it automatically.
	ConfigureNat(options *awsec2.ConfigureNatOptions)
	// Configures subnet with the gateway.
	//
	// Don't call this directly, the VPC will call it automatically.
	ConfigureSubnet(subnet awsec2.PrivateSubnet)
}

// The jsii proxy struct for TransitGatewayNatProvider
type jsiiProxy_TransitGatewayNatProvider struct {
	internal.Type__awsec2NatProvider
}

func (j *jsiiProxy_TransitGatewayNatProvider) ConfiguredGateways() *[]*awsec2.GatewayConfig {
	var returns *[]*awsec2.GatewayConfig
	_jsii_.Get(
		j,
		"configuredGateways",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransitGatewayNatProvider) TransitGateway() ITransitGateway {
	var returns ITransitGateway
	_jsii_.Get(
		j,
		"transitGateway",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransitGatewayNatProvider) TransitGatewayAttachment() TransitGatewayAttachment {
	var returns TransitGatewayAttachment
	_jsii_.Get(
		j,
		"transitGatewayAttachment",
		&returns,
	)
	return returns
}


func NewTransitGatewayNatProvider(options *TransitGatewayNatProviderOptions) TransitGatewayNatProvider {
	_init_.Initialize()

	if err := validateNewTransitGatewayNatProviderParameters(options); err != nil {
		panic(err)
	}
	j := jsiiProxy_TransitGatewayNatProvider{}

	_jsii_.Create(
		"cdk-extensions.ec2.TransitGatewayNatProvider",
		[]interface{}{options},
		&j,
	)

	return &j
}

func NewTransitGatewayNatProvider_Override(t TransitGatewayNatProvider, options *TransitGatewayNatProviderOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk-extensions.ec2.TransitGatewayNatProvider",
		[]interface{}{options},
		t,
	)
}

// Use NAT Gateways to provide NAT services for your VPC.
//
// NAT gateways are managed by AWS.
// See: https://docs.aws.amazon.com/vpc/latest/userguide/vpc-nat-gateway.html
//
func TransitGatewayNatProvider_Gateway(props *awsec2.NatGatewayProps) awsec2.NatProvider {
	_init_.Initialize()

	if err := validateTransitGatewayNatProvider_GatewayParameters(props); err != nil {
		panic(err)
	}
	var returns awsec2.NatProvider

	_jsii_.StaticInvoke(
		"cdk-extensions.ec2.TransitGatewayNatProvider",
		"gateway",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Use NAT instances to provide NAT services for your VPC.
//
// NAT instances are managed by you, but in return allow more configuration.
//
// Be aware that instances created using this provider will not be
// automatically replaced if they are stopped for any reason. You should implement
// your own NatProvider based on AutoScaling groups if you need that.
// See: https://docs.aws.amazon.com/vpc/latest/userguide/VPC_NAT_Instance.html
//
func TransitGatewayNatProvider_Instance(props *awsec2.NatInstanceProps) awsec2.NatInstanceProvider {
	_init_.Initialize()

	if err := validateTransitGatewayNatProvider_InstanceParameters(props); err != nil {
		panic(err)
	}
	var returns awsec2.NatInstanceProvider

	_jsii_.StaticInvoke(
		"cdk-extensions.ec2.TransitGatewayNatProvider",
		"instance",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TransitGatewayNatProvider) ConfigureNat(options *awsec2.ConfigureNatOptions) {
	if err := t.validateConfigureNatParameters(options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"configureNat",
		[]interface{}{options},
	)
}

func (t *jsiiProxy_TransitGatewayNatProvider) ConfigureSubnet(subnet awsec2.PrivateSubnet) {
	if err := t.validateConfigureSubnetParameters(subnet); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"configureSubnet",
		[]interface{}{subnet},
	)
}

