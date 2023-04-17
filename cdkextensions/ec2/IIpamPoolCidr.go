package ec2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/vibe-io/cdk-extensions-go/cdkextensions/ec2/internal"
)

type IIpamPoolCidr interface {
	awscdk.IResource
	IpamPoolCidrId() *string
	IpamPoolCidrState() *string
}

// The jsii proxy for IIpamPoolCidr
type jsiiProxy_IIpamPoolCidr struct {
	internal.Type__awscdkIResource
}

func (j *jsiiProxy_IIpamPoolCidr) IpamPoolCidrId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipamPoolCidrId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IIpamPoolCidr) IpamPoolCidrState() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipamPoolCidrState",
		&returns,
	)
	return returns
}

