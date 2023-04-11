package ec2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/vibe-io/cdk-extensions-go/cdkextensions/ec2/internal"
)

// Represents an IPAM resource discovery in AWS.
type IIpamResourceDiscovery interface {
	awscdk.IResource
	AddIpam(id *string, options *IpamProps) IIpam
	AssociateIpam(ipam IIpam) IIpamResourceDiscoveryAssociation
	// The resource discovery ARN.
	IpamResourceDiscoveryArn() *string
	// The resource discovery ID.
	IpamResourceDiscoveryId() *string
	// Defines if the resource discovery is the default.
	//
	// The default resource
	// discovery is the resource discovery automatically created when you create
	// an IPAM.
	IpamResourceDiscoveryIsDefault() awscdk.IResolvable
	// The owner ID.
	IpamResourceDiscoveryOwnerId() *string
	// The resource discovery Region.
	IpamResourceDiscoveryRegion() *string
	// The resource discovery's state.
	//
	// - create-in-progress - Resource discovery is being created.
	// - create-complete - Resource discovery creation is complete.
	// - create-failed - Resource discovery creation has failed.
	// - modify-in-progress - Resource discovery is being modified.
	// - modify-complete - Resource discovery modification is complete.
	// - modify-failed - Resource discovery modification has failed.
	// - delete-in-progress - Resource discovery is being deleted.
	// - delete-complete - Resource discovery deletion is complete.
	// - delete-failed - Resource discovery deletion has failed.
	// - isolate-in-progress - AWS account that created the resource discovery
	// has been removed and the resource discovery is being isolated.
	// - isolate-complete - Resource discovery isolation is complete.
	// - restore-in-progress - AWS account that created the resource discovery
	// and was isolated has been restored.
	IpamResourceDiscoveryState() *string
}

// The jsii proxy for IIpamResourceDiscovery
type jsiiProxy_IIpamResourceDiscovery struct {
	internal.Type__awscdkIResource
}

func (i *jsiiProxy_IIpamResourceDiscovery) AddIpam(id *string, options *IpamProps) IIpam {
	if err := i.validateAddIpamParameters(id, options); err != nil {
		panic(err)
	}
	var returns IIpam

	_jsii_.Invoke(
		i,
		"addIpam",
		[]interface{}{id, options},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IIpamResourceDiscovery) AssociateIpam(ipam IIpam) IIpamResourceDiscoveryAssociation {
	if err := i.validateAssociateIpamParameters(ipam); err != nil {
		panic(err)
	}
	var returns IIpamResourceDiscoveryAssociation

	_jsii_.Invoke(
		i,
		"associateIpam",
		[]interface{}{ipam},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_IIpamResourceDiscovery) IpamResourceDiscoveryArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipamResourceDiscoveryArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IIpamResourceDiscovery) IpamResourceDiscoveryId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipamResourceDiscoveryId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IIpamResourceDiscovery) IpamResourceDiscoveryIsDefault() awscdk.IResolvable {
	var returns awscdk.IResolvable
	_jsii_.Get(
		j,
		"ipamResourceDiscoveryIsDefault",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IIpamResourceDiscovery) IpamResourceDiscoveryOwnerId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipamResourceDiscoveryOwnerId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IIpamResourceDiscovery) IpamResourceDiscoveryRegion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipamResourceDiscoveryRegion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IIpamResourceDiscovery) IpamResourceDiscoveryState() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipamResourceDiscoveryState",
		&returns,
	)
	return returns
}

