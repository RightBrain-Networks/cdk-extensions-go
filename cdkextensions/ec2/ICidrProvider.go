package ec2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
)

type ICidrProvider interface {
	IpamOptions() *awsec2.VpcIpamOptions
	IpamPool() IIpamPool
	Netmask() *float64
}

// The jsii proxy for ICidrProvider
type jsiiProxy_ICidrProvider struct {
	_ byte // padding
}

func (j *jsiiProxy_ICidrProvider) IpamOptions() *awsec2.VpcIpamOptions {
	var returns *awsec2.VpcIpamOptions
	_jsii_.Get(
		j,
		"ipamOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ICidrProvider) IpamPool() IIpamPool {
	var returns IIpamPool
	_jsii_.Get(
		j,
		"ipamPool",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ICidrProvider) Netmask() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"netmask",
		&returns,
	)
	return returns
}

