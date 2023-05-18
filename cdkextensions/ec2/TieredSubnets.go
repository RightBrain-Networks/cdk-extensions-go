package ec2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/vibe-io/cdk-extensions-go/cdkextensions/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/vibe-io/cdk-extensions-go/cdkextensions/ec2/internal"
)

type TieredSubnets interface {
	awsec2.IIpAddresses
	IpamOptions() *awsec2.VpcIpamOptions
	IpamPool() IIpamPool
	Netmask() *float64
	TierMask() *float64
	// Called by the VPC to retrieve Subnet options from the Ipam.
	//
	// Don't call this directly, the VPC will call it automatically.
	AllocateSubnetsCidr(input *awsec2.AllocateCidrRequest) *awsec2.SubnetIpamOptions
	// Called by the VPC to retrieve VPC options from the Ipam.
	//
	// Don't call this directly, the VPC will call it automatically.
	AllocateVpcCidr() *awsec2.VpcIpamOptions
}

// The jsii proxy struct for TieredSubnets
type jsiiProxy_TieredSubnets struct {
	internal.Type__awsec2IIpAddresses
}

func (j *jsiiProxy_TieredSubnets) IpamOptions() *awsec2.VpcIpamOptions {
	var returns *awsec2.VpcIpamOptions
	_jsii_.Get(
		j,
		"ipamOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TieredSubnets) IpamPool() IIpamPool {
	var returns IIpamPool
	_jsii_.Get(
		j,
		"ipamPool",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TieredSubnets) Netmask() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"netmask",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TieredSubnets) TierMask() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"tierMask",
		&returns,
	)
	return returns
}


func NewTieredSubnets(options *TieredSubnetsOptions) TieredSubnets {
	_init_.Initialize()

	if err := validateNewTieredSubnetsParameters(options); err != nil {
		panic(err)
	}
	j := jsiiProxy_TieredSubnets{}

	_jsii_.Create(
		"cdk-extensions.ec2.TieredSubnets",
		[]interface{}{options},
		&j,
	)

	return &j
}

func NewTieredSubnets_Override(t TieredSubnets, options *TieredSubnetsOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk-extensions.ec2.TieredSubnets",
		[]interface{}{options},
		t,
	)
}

func (t *jsiiProxy_TieredSubnets) AllocateSubnetsCidr(input *awsec2.AllocateCidrRequest) *awsec2.SubnetIpamOptions {
	if err := t.validateAllocateSubnetsCidrParameters(input); err != nil {
		panic(err)
	}
	var returns *awsec2.SubnetIpamOptions

	_jsii_.Invoke(
		t,
		"allocateSubnetsCidr",
		[]interface{}{input},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TieredSubnets) AllocateVpcCidr() *awsec2.VpcIpamOptions {
	var returns *awsec2.VpcIpamOptions

	_jsii_.Invoke(
		t,
		"allocateVpcCidr",
		nil, // no parameters
		&returns,
	)

	return returns
}

