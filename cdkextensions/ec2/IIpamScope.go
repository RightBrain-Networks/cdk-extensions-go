package ec2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Represents an IPAM scope in AWS.
type IIpamScope interface {
	// Adds an IPAM pool to the IPAM scope.
	//
	// A pool is a collection of contiguous IP address ranges (or CIDRs). IPAM
	// pools enable you to organize your IP addresses according to your routing
	// and security needs.
	//
	// Returns: The pool that was added to the scope.
	// See: [How IPAM works](https://docs.aws.amazon.com/vpc/latest/ipam/how-it-works-ipam.html)
	//
	AddPool(id *string, options *IpamPoolOptions) IIpamPool
	// The ARN of the scope.
	IpamScopeArn() *string
	// The ID of an IPAM scope.
	IpamScopeId() *string
	// The ARN of an IPAM.
	IpamScopeIpamArn() *string
	// Defines if the scope is the default scope or not.
	IpamScopeIsDefault() awscdk.IResolvable
	// The number of pools in a scope.
	IpamScopePoolCount() *float64
	// The type of the scope.
	IpamScopeType() *string
}

// The jsii proxy for IIpamScope
type jsiiProxy_IIpamScope struct {
	_ byte // padding
}

func (i *jsiiProxy_IIpamScope) AddPool(id *string, options *IpamPoolOptions) IIpamPool {
	if err := i.validateAddPoolParameters(id, options); err != nil {
		panic(err)
	}
	var returns IIpamPool

	_jsii_.Invoke(
		i,
		"addPool",
		[]interface{}{id, options},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_IIpamScope) IpamScopeArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipamScopeArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IIpamScope) IpamScopeId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipamScopeId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IIpamScope) IpamScopeIpamArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipamScopeIpamArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IIpamScope) IpamScopeIsDefault() awscdk.IResolvable {
	var returns awscdk.IResolvable
	_jsii_.Get(
		j,
		"ipamScopeIsDefault",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IIpamScope) IpamScopePoolCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ipamScopePoolCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IIpamScope) IpamScopeType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipamScopeType",
		&returns,
	)
	return returns
}

