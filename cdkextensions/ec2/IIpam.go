package ec2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Represents an IPAM in AWS.
type IIpam interface {
	// Adds an IPAM scope to the IPAM.
	//
	// In IPAM, a scope is the highest-level container within IPAM. Scopes enable
	// you to reuse IP addresses across multiple unconnected networks without
	// causing IP address overlap or conflict.
	//
	// Returns: The scope that was added to the IPAM.
	// See: [How IPAM works](https://docs.aws.amazon.com/vpc/latest/ipam/how-it-works-ipam.html)
	//
	AddScope(id *string, options *IpamScopeOptions) IIpamScope
	// Associates an existing IPAM resource discovery with the IPAM.
	//
	// IPAM aggregates the resource CIDRs discovered by the associated resource
	// discovery.
	//
	// Returns: The association resource that handles the association between the
	// IPAM and the resource discovery.
	AssociateResourceDiscovery(resourceDiscovery IIpamResourceDiscovery) IIpamResourceDiscoveryAssociation
	// The ARN of the IPAM.
	IpamArn() *string
	// The ID of the IPAM.
	IpamId() *string
	// The ID of the IPAM's default private scope.
	IpamPrivateDefaultScopeId() *string
	// The ID of the IPAM's default public scope.
	IpamPublicDefaultScopeId() *string
	// The number of scopes in the IPAM.
	//
	// The scope quota is 5.
	IpamScopeCount() *float64
}

// The jsii proxy for IIpam
type jsiiProxy_IIpam struct {
	_ byte // padding
}

func (i *jsiiProxy_IIpam) AddScope(id *string, options *IpamScopeOptions) IIpamScope {
	if err := i.validateAddScopeParameters(id, options); err != nil {
		panic(err)
	}
	var returns IIpamScope

	_jsii_.Invoke(
		i,
		"addScope",
		[]interface{}{id, options},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IIpam) AssociateResourceDiscovery(resourceDiscovery IIpamResourceDiscovery) IIpamResourceDiscoveryAssociation {
	if err := i.validateAssociateResourceDiscoveryParameters(resourceDiscovery); err != nil {
		panic(err)
	}
	var returns IIpamResourceDiscoveryAssociation

	_jsii_.Invoke(
		i,
		"associateResourceDiscovery",
		[]interface{}{resourceDiscovery},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_IIpam) IpamArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipamArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IIpam) IpamId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipamId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IIpam) IpamPrivateDefaultScopeId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipamPrivateDefaultScopeId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IIpam) IpamPublicDefaultScopeId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipamPublicDefaultScopeId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IIpam) IpamScopeCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ipamScopeCount",
		&returns,
	)
	return returns
}

