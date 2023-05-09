package ec2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/vibe-io/cdk-extensions-go/cdkextensions/ec2/internal"
)

type IVpcCidrBlock interface {
	awscdk.IResource
	VpcCidrBlockAssociationId() *string
	VpcCidrBlockCidr() *string
}

// The jsii proxy for IVpcCidrBlock
type jsiiProxy_IVpcCidrBlock struct {
	internal.Type__awscdkIResource
}

func (j *jsiiProxy_IVpcCidrBlock) VpcCidrBlockAssociationId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vpcCidrBlockAssociationId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IVpcCidrBlock) VpcCidrBlockCidr() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vpcCidrBlockCidr",
		&returns,
	)
	return returns
}

