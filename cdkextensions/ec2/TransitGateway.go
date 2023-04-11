package ec2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/vibe-io/cdk-extensions-go/cdkextensions/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/vibe-io/cdk-extensions-go/cdkextensions/ec2/internal"
	"github.com/vibe-io/cdk-extensions-go/cdkextensions/ram"
)

type TransitGateway interface {
	awscdk.Resource
	ITransitGateway
	// A private Autonomous System Number (ASN) for the Amazon side of a BGP session.
	//
	// The range is 64512 to 65534 for 16-bit ASNs.
	// See: [TransitGateway.AmazonSideAsn](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-transitgateway.html#cfn-ec2-transitgateway-amazonsideasn)
	//
	AmazonSideAsn() *float64
	// Enable or disable automatic acceptance of attachment requests.
	//
	// When enabled any new transit gateway attachments that are created in other
	// accounts via a resource share will be accepted automatically. Otherwise,
	// manual intervention will be required to approve all new attachments.
	//
	// This is disabled by default to maintain the highest levels of security,
	// however enabling should be strongly considered as without this full
	// automation of infrastructure will not be possible for cross account
	// setups.
	// See: [Accept a shared attachment](https://docs.aws.amazon.com/vpc/latest/tgw/tgw-transit-gateways.html#tgw-accept-shared-attachment)
	//
	AutoAcceptSharedAttachments() *bool
	// The default route table that got created along with the Transit Gateway.
	//
	// This information is not exposed by CloudFormation. As such, this resource
	// will only be available if the default reoute table ID is passed in.
	DefaultRouteTable() ITransitGatewayRouteTable
	// Enable or disable automatic association with the default association route table.
	//
	// When enabled, all new attachments that are accepted will be automatically
	// associated with the default association route table. By default this is
	// the route table that is created automatically when the transit gateway is
	// created.
	// See: [TransitGateway.DefaultRouteTableAssociation](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-transitgateway.html#cfn-ec2-transitgateway-defaultroutetableassociation)
	//
	DefaultRouteTableAssociation() *bool
	// The ID of the default route table that was created with the transit gateway.
	//
	// This route table is critical to some transit gateway architectures and is
	// not exposed by CloudFormation.
	//
	// Passing in the ID of the default route table will make an object available
	// that represents it and can be used for further operations.
	DefaultRouteTableId() *string
	// Enable or disable automatic propagation of routes to the default propagation route table.
	//
	// When a new attachment is accepted, the routes associated with that
	// attachment will automatically be added to the default propagation route
	// table. By default this is the route table that is created automatically
	// when the transit gateway is created.
	// See: [Route propagation](https://docs.aws.amazon.com/vpc/latest/tgw/how-transit-gateways-work.html#tgw-route-propagation-overview)
	//
	DefaultRouteTablePropagation() *bool
	// The description of the transit gateway.
	// See: [TransitGateway.Description](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-transitgateway.html#cfn-ec2-transitgateway-description)
	//
	Description() *string
	// Enable or disable DNS support.
	//
	// When DNS support is enabled on a transit gateway, VPC DNS resolution in
	// attached VPC's will automatically resolve public IP addresses from other
	// VPC's to their provate IP address equivalent.
	// See: [Create a transit gateway](https://docs.aws.amazon.com/vpc/latest/tgw/tgw-transit-gateways.html#create-tgw)
	//
	DnsSupport() *bool
	// The environment this resource belongs to.
	//
	// For resources that are created and managed by the CDK
	// (generally, those created by creating new class instances like Role, Bucket, etc.),
	// this is always the same as the environment of the stack they belong to;
	// however, for imported resources
	// (those obtained from static methods like fromRoleArn, fromBucketName, etc.),
	// that might be different than the stack they were imported into.
	Env() *awscdk.ResourceEnvironment
	// Indicates whether multicast is enabled on the transit gateway.
	// See: [Multicast reference architectures](https://d1.awsstatic.com/architecture-diagrams/ArchitectureDiagrams/transitgateway_multicast_ra.pdf?did=wp_card&trk=wp_card)
	//
	MulticastSupport() *bool
	// The name of the transit gateway.
	//
	// Used to tag the transit gateway with a name that will be displayed in the
	// AWS VPC console.
	// See: [TransitGateway.Tags](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-transitgateway.html#cfn-ec2-transitgateway-tags)
	//
	Name() *string
	// The tree node.
	Node() constructs.Node
	// Returns a string-encoded token that resolves to the physical name that should be passed to the CloudFormation resource.
	//
	// This value will resolve to one of the following:
	// - a concrete value (e.g. `"my-awesome-bucket"`)
	// - `undefined`, when a name should be generated by CloudFormation
	// - a concrete name generated automatically during synthesis, in
	//   cross-environment scenarios.
	PhysicalName() *string
	// The underlying TransitGateway CloudFormation resource.
	// See: [AWS::EC2::TransitGateway](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-transitgateway.html)
	//
	Resource() awsec2.CfnTransitGateway
	// The RAM resource share that is used for sharing the transit gateway with other accounts.
	ResourceShare() ram.ResourceShare
	// The stack in which this resource is defined.
	Stack() awscdk.Stack
	// The ARN of this Transit Gateway.
	TransitGatewayArn() *string
	// The ID of this Transit Gateway.
	TransitGatewayId() *string
	// Enable or disable Equal Cost Multipath Protocol support.
	// See: [Achieve ECMP with multiple VPN tunnels](https://aws.amazon.com/premiumsupport/knowledge-center/transit-gateway-ecmp-multiple-tunnels/)
	//
	VpnEcmpSupport() *bool
	AddCidrBlock(cidr *string)
	// Creates a new Transit Gateway Route Table for this Transit Gateway.
	//
	// Returns: The newly created Transit Gateway Route Table.
	AddRouteTable(options *TransitGatewayRouteTableOptions) TransitGatewayRouteTable
	// Creates a new VPN connection that terminates on the AWS side at this Transit Gateway.
	//
	// Returns: The VPN connection that was created.
	AddVpn(id *string, options *VpnAttachmentOptions) VpnConnection
	// Apply the given removal policy to this resource.
	//
	// The Removal Policy controls what happens to this resource when it stops
	// being managed by CloudFormation, either because you've removed it from the
	// CDK application or because you've made a change that requires the resource
	// to be replaced.
	//
	// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
	// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy)
	// Creates a new transit gateway peering attachment for this transit gateway.
	//
	// Returns: The newly created TransitGatewayPeeringAttachment.
	AttachPeer(peer ITransitGateway, options *TransitGatewayPeeringAttachmentOptions) TransitGatewayPeeringAttachment
	// Creates a new VPC transit gateway attachment for this transit gateway.
	//
	// Returns: The newly created TransitGatewayAttachment.
	AttachVpc(vpc awsec2.IVpc, options *VpcAttachmentOptions) TransitGatewayAttachment
	EnableSharing(options *SharingOptions) ram.ResourceShare
	GeneratePhysicalName() *string
	// Returns an environment-sensitive token that should be used for the resource's "ARN" attribute (e.g. `bucket.bucketArn`).
	//
	// Normally, this token will resolve to `arnAttr`, but if the resource is
	// referenced across environments, `arnComponents` will be used to synthesize
	// a concrete ARN with the resource's physical name. Make sure to reference
	// `this.physicalName` in `arnComponents`.
	GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string
	// Returns an environment-sensitive token that should be used for the resource's "name" attribute (e.g. `bucket.bucketName`).
	//
	// Normally, this token will resolve to `nameAttr`, but if the resource is
	// referenced across environments, it will be resolved to `this.physicalName`,
	// which will be a concrete name.
	GetResourceNameAttribute(nameAttr *string) *string
	// Returns a string representation of this construct.
	ToString() *string
}

// The jsii proxy struct for TransitGateway
type jsiiProxy_TransitGateway struct {
	internal.Type__awscdkResource
	jsiiProxy_ITransitGateway
}

func (j *jsiiProxy_TransitGateway) AmazonSideAsn() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"amazonSideAsn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransitGateway) AutoAcceptSharedAttachments() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"autoAcceptSharedAttachments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransitGateway) DefaultRouteTable() ITransitGatewayRouteTable {
	var returns ITransitGatewayRouteTable
	_jsii_.Get(
		j,
		"defaultRouteTable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransitGateway) DefaultRouteTableAssociation() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"defaultRouteTableAssociation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransitGateway) DefaultRouteTableId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultRouteTableId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransitGateway) DefaultRouteTablePropagation() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"defaultRouteTablePropagation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransitGateway) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransitGateway) DnsSupport() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"dnsSupport",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransitGateway) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransitGateway) MulticastSupport() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"multicastSupport",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransitGateway) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransitGateway) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransitGateway) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransitGateway) Resource() awsec2.CfnTransitGateway {
	var returns awsec2.CfnTransitGateway
	_jsii_.Get(
		j,
		"resource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransitGateway) ResourceShare() ram.ResourceShare {
	var returns ram.ResourceShare
	_jsii_.Get(
		j,
		"resourceShare",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransitGateway) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransitGateway) TransitGatewayArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"transitGatewayArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransitGateway) TransitGatewayId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"transitGatewayId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransitGateway) VpnEcmpSupport() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"vpnEcmpSupport",
		&returns,
	)
	return returns
}


// Creates a new instance of the Database class.
func NewTransitGateway(scope constructs.Construct, id *string, props *TransitGatewayProps) TransitGateway {
	_init_.Initialize()

	if err := validateNewTransitGatewayParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_TransitGateway{}

	_jsii_.Create(
		"cdk-extensions.ec2.TransitGateway",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Creates a new instance of the Database class.
func NewTransitGateway_Override(t TransitGateway, scope constructs.Construct, id *string, props *TransitGatewayProps) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk-extensions.ec2.TransitGateway",
		[]interface{}{scope, id, props},
		t,
	)
}

func TransitGateway_FromTransitGatewayId(scope constructs.IConstruct, id *string, transitGatewayId *string) ITransitGateway {
	_init_.Initialize()

	if err := validateTransitGateway_FromTransitGatewayIdParameters(scope, id, transitGatewayId); err != nil {
		panic(err)
	}
	var returns ITransitGateway

	_jsii_.StaticInvoke(
		"cdk-extensions.ec2.TransitGateway",
		"fromTransitGatewayId",
		[]interface{}{scope, id, transitGatewayId},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead.
func TransitGateway_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateTransitGateway_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"cdk-extensions.ec2.TransitGateway",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Returns true if the construct was created by CDK, and false otherwise.
func TransitGateway_IsOwnedResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	if err := validateTransitGateway_IsOwnedResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"cdk-extensions.ec2.TransitGateway",
		"isOwnedResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
func TransitGateway_IsResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	if err := validateTransitGateway_IsResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"cdk-extensions.ec2.TransitGateway",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TransitGateway) AddCidrBlock(cidr *string) {
	if err := t.validateAddCidrBlockParameters(cidr); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"addCidrBlock",
		[]interface{}{cidr},
	)
}

func (t *jsiiProxy_TransitGateway) AddRouteTable(options *TransitGatewayRouteTableOptions) TransitGatewayRouteTable {
	if err := t.validateAddRouteTableParameters(options); err != nil {
		panic(err)
	}
	var returns TransitGatewayRouteTable

	_jsii_.Invoke(
		t,
		"addRouteTable",
		[]interface{}{options},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TransitGateway) AddVpn(id *string, options *VpnAttachmentOptions) VpnConnection {
	if err := t.validateAddVpnParameters(id, options); err != nil {
		panic(err)
	}
	var returns VpnConnection

	_jsii_.Invoke(
		t,
		"addVpn",
		[]interface{}{id, options},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TransitGateway) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	if err := t.validateApplyRemovalPolicyParameters(policy); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (t *jsiiProxy_TransitGateway) AttachPeer(peer ITransitGateway, options *TransitGatewayPeeringAttachmentOptions) TransitGatewayPeeringAttachment {
	if err := t.validateAttachPeerParameters(peer, options); err != nil {
		panic(err)
	}
	var returns TransitGatewayPeeringAttachment

	_jsii_.Invoke(
		t,
		"attachPeer",
		[]interface{}{peer, options},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TransitGateway) AttachVpc(vpc awsec2.IVpc, options *VpcAttachmentOptions) TransitGatewayAttachment {
	if err := t.validateAttachVpcParameters(vpc, options); err != nil {
		panic(err)
	}
	var returns TransitGatewayAttachment

	_jsii_.Invoke(
		t,
		"attachVpc",
		[]interface{}{vpc, options},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TransitGateway) EnableSharing(options *SharingOptions) ram.ResourceShare {
	if err := t.validateEnableSharingParameters(options); err != nil {
		panic(err)
	}
	var returns ram.ResourceShare

	_jsii_.Invoke(
		t,
		"enableSharing",
		[]interface{}{options},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TransitGateway) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"generatePhysicalName",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TransitGateway) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	if err := t.validateGetResourceArnAttributeParameters(arnAttr, arnComponents); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		t,
		"getResourceArnAttribute",
		[]interface{}{arnAttr, arnComponents},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TransitGateway) GetResourceNameAttribute(nameAttr *string) *string {
	if err := t.validateGetResourceNameAttributeParameters(nameAttr); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		t,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TransitGateway) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

