package ec2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/vibe-io/cdk-extensions-go/cdkextensions/ec2/internal"
)

type IIpamPool interface {
	awscdk.IResource
	AddChildPool(id *string, options *AddChildPoolOptions) IIpamPool
	AddCidrToPool(id *string, options *AddCidrToPoolOptions) *AddCidrToPoolResult
	AllocateCidrFromPool(id *string, options *AllocateCidrFromPoolOptions) IIpamAllocation
	IpamPoolArn() *string
	IpamPoolDepth() *float64
	IpamPoolId() *string
	IpamPoolIpamArn() *string
	IpamPoolScopeArn() *string
	IpamPoolScopeType() *string
	IpamPoolState() *string
	IpamPoolStateMessage() *string
}

// The jsii proxy for IIpamPool
type jsiiProxy_IIpamPool struct {
	internal.Type__awscdkIResource
}

func (i *jsiiProxy_IIpamPool) AddChildPool(id *string, options *AddChildPoolOptions) IIpamPool {
	if err := i.validateAddChildPoolParameters(id, options); err != nil {
		panic(err)
	}
	var returns IIpamPool

	_jsii_.Invoke(
		i,
		"addChildPool",
		[]interface{}{id, options},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IIpamPool) AddCidrToPool(id *string, options *AddCidrToPoolOptions) *AddCidrToPoolResult {
	if err := i.validateAddCidrToPoolParameters(id, options); err != nil {
		panic(err)
	}
	var returns *AddCidrToPoolResult

	_jsii_.Invoke(
		i,
		"addCidrToPool",
		[]interface{}{id, options},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IIpamPool) AllocateCidrFromPool(id *string, options *AllocateCidrFromPoolOptions) IIpamAllocation {
	if err := i.validateAllocateCidrFromPoolParameters(id, options); err != nil {
		panic(err)
	}
	var returns IIpamAllocation

	_jsii_.Invoke(
		i,
		"allocateCidrFromPool",
		[]interface{}{id, options},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_IIpamPool) IpamPoolArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipamPoolArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IIpamPool) IpamPoolDepth() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ipamPoolDepth",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IIpamPool) IpamPoolId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipamPoolId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IIpamPool) IpamPoolIpamArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipamPoolIpamArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IIpamPool) IpamPoolScopeArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipamPoolScopeArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IIpamPool) IpamPoolScopeType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipamPoolScopeType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IIpamPool) IpamPoolState() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipamPoolState",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IIpamPool) IpamPoolStateMessage() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipamPoolStateMessage",
		&returns,
	)
	return returns
}

