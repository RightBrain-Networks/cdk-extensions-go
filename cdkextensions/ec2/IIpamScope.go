package ec2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/vibe-io/cdk-extensions-go/cdkextensions/ec2/internal"
)

// Represents an IPAM scope in AWS.
type IIpamScope interface {
	awscdk.IResource
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
	internal.Type__awscdkIResource
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

