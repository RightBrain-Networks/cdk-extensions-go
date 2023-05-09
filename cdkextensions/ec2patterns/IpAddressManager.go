package ec2patterns

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/vibe-io/cdk-extensions-go/cdkextensions/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/vibe-io/cdk-extensions-go/cdkextensions/ec2"
	"github.com/vibe-io/cdk-extensions-go/cdkextensions/ec2patterns/internal"
	"github.com/vibe-io/cdk-extensions-go/cdkextensions/ram"
)

type IpAddressManager interface {
	awscdk.Resource
	AllowExternalPricipals() *bool
	ClientVpnAllocationMask() *float64
	// The environment this resource belongs to.
	//
	// For resources that are created and managed by the CDK
	// (generally, those created by creating new class instances like Role, Bucket, etc.),
	// this is always the same as the environment of the stack they belong to;
	// however, for imported resources
	// (those obtained from static methods like fromRoleArn, fromBucketName, etc.),
	// that might be different than the stack they were imported into.
	Env() *awscdk.ResourceEnvironment
	Ipam() ec2.Ipam
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
	ResourceShare() ram.ResourceShare
	SharingEnabled() *bool
	// The stack in which this resource is defined.
	Stack() awscdk.Stack
	VpcAllocationMask() *float64
	AddRegion(region *string)
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
	GeneratePhysicalName() *string
	GetClientVpnConfiguration(scope constructs.IConstruct, id *string, options *GetClientVpnConfigurationOptions) ec2.IIpv4CidrAssignment
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
	GetVpcConfiguration(scope constructs.IConstruct, id *string, options *GetVpcConfigurationOptions) ec2.IIpv4CidrAssignment
	PrivateVpcPoolForEnvironment(account *string, region *string) ec2.IIpamPool
	PrivateVpcPoolForRegion(region *string) ec2.IIpamPool
	PrivateVpnPoolForEnvironment(account *string, region *string) ec2.IIpamPool
	PrivateVpnPoolForRegion(region *string) ec2.IIpamPool
	RegisterAccount(account *string, pool ec2.IIpamPool)
	RegisterCidr(scope constructs.IConstruct, id *string, cidr *string)
	// Returns a string representation of this construct.
	ToString() *string
}

// The jsii proxy struct for IpAddressManager
type jsiiProxy_IpAddressManager struct {
	internal.Type__awscdkResource
}

func (j *jsiiProxy_IpAddressManager) AllowExternalPricipals() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"allowExternalPricipals",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IpAddressManager) ClientVpnAllocationMask() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"clientVpnAllocationMask",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IpAddressManager) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IpAddressManager) Ipam() ec2.Ipam {
	var returns ec2.Ipam
	_jsii_.Get(
		j,
		"ipam",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IpAddressManager) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IpAddressManager) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IpAddressManager) ResourceShare() ram.ResourceShare {
	var returns ram.ResourceShare
	_jsii_.Get(
		j,
		"resourceShare",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IpAddressManager) SharingEnabled() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"sharingEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IpAddressManager) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IpAddressManager) VpcAllocationMask() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"vpcAllocationMask",
		&returns,
	)
	return returns
}


func NewIpAddressManager(scope constructs.IConstruct, id *string, props *IpAddressManagerProps) IpAddressManager {
	_init_.Initialize()

	if err := validateNewIpAddressManagerParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_IpAddressManager{}

	_jsii_.Create(
		"cdk-extensions.ec2_patterns.IpAddressManager",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewIpAddressManager_Override(i IpAddressManager, scope constructs.IConstruct, id *string, props *IpAddressManagerProps) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk-extensions.ec2_patterns.IpAddressManager",
		[]interface{}{scope, id, props},
		i,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead.
func IpAddressManager_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateIpAddressManager_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"cdk-extensions.ec2_patterns.IpAddressManager",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Returns true if the construct was created by CDK, and false otherwise.
func IpAddressManager_IsOwnedResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	if err := validateIpAddressManager_IsOwnedResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"cdk-extensions.ec2_patterns.IpAddressManager",
		"isOwnedResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
func IpAddressManager_IsResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	if err := validateIpAddressManager_IsResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"cdk-extensions.ec2_patterns.IpAddressManager",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

func IpAddressManager_DEFAULT_CLIENT_VPN_ALLOCATION_MASK() *float64 {
	_init_.Initialize()
	var returns *float64
	_jsii_.StaticGet(
		"cdk-extensions.ec2_patterns.IpAddressManager",
		"DEFAULT_CLIENT_VPN_ALLOCATION_MASK",
		&returns,
	)
	return returns
}

func IpAddressManager_DEFAULT_VPC_ALLOCATION_MASK() *float64 {
	_init_.Initialize()
	var returns *float64
	_jsii_.StaticGet(
		"cdk-extensions.ec2_patterns.IpAddressManager",
		"DEFAULT_VPC_ALLOCATION_MASK",
		&returns,
	)
	return returns
}

func IpAddressManager_DEFAULT_VPC_POOL_CIDRS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"cdk-extensions.ec2_patterns.IpAddressManager",
		"DEFAULT_VPC_POOL_CIDRS",
		&returns,
	)
	return returns
}

func IpAddressManager_DEFAULT_VPN_POOL_CIDRS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"cdk-extensions.ec2_patterns.IpAddressManager",
		"DEFAULT_VPN_POOL_CIDRS",
		&returns,
	)
	return returns
}

func (i *jsiiProxy_IpAddressManager) AddRegion(region *string) {
	if err := i.validateAddRegionParameters(region); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"addRegion",
		[]interface{}{region},
	)
}

func (i *jsiiProxy_IpAddressManager) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	if err := i.validateApplyRemovalPolicyParameters(policy); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (i *jsiiProxy_IpAddressManager) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"generatePhysicalName",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IpAddressManager) GetClientVpnConfiguration(scope constructs.IConstruct, id *string, options *GetClientVpnConfigurationOptions) ec2.IIpv4CidrAssignment {
	if err := i.validateGetClientVpnConfigurationParameters(scope, id, options); err != nil {
		panic(err)
	}
	var returns ec2.IIpv4CidrAssignment

	_jsii_.Invoke(
		i,
		"getClientVpnConfiguration",
		[]interface{}{scope, id, options},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IpAddressManager) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	if err := i.validateGetResourceArnAttributeParameters(arnAttr, arnComponents); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		i,
		"getResourceArnAttribute",
		[]interface{}{arnAttr, arnComponents},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IpAddressManager) GetResourceNameAttribute(nameAttr *string) *string {
	if err := i.validateGetResourceNameAttributeParameters(nameAttr); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		i,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IpAddressManager) GetVpcConfiguration(scope constructs.IConstruct, id *string, options *GetVpcConfigurationOptions) ec2.IIpv4CidrAssignment {
	if err := i.validateGetVpcConfigurationParameters(scope, id, options); err != nil {
		panic(err)
	}
	var returns ec2.IIpv4CidrAssignment

	_jsii_.Invoke(
		i,
		"getVpcConfiguration",
		[]interface{}{scope, id, options},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IpAddressManager) PrivateVpcPoolForEnvironment(account *string, region *string) ec2.IIpamPool {
	if err := i.validatePrivateVpcPoolForEnvironmentParameters(account, region); err != nil {
		panic(err)
	}
	var returns ec2.IIpamPool

	_jsii_.Invoke(
		i,
		"privateVpcPoolForEnvironment",
		[]interface{}{account, region},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IpAddressManager) PrivateVpcPoolForRegion(region *string) ec2.IIpamPool {
	if err := i.validatePrivateVpcPoolForRegionParameters(region); err != nil {
		panic(err)
	}
	var returns ec2.IIpamPool

	_jsii_.Invoke(
		i,
		"privateVpcPoolForRegion",
		[]interface{}{region},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IpAddressManager) PrivateVpnPoolForEnvironment(account *string, region *string) ec2.IIpamPool {
	if err := i.validatePrivateVpnPoolForEnvironmentParameters(account, region); err != nil {
		panic(err)
	}
	var returns ec2.IIpamPool

	_jsii_.Invoke(
		i,
		"privateVpnPoolForEnvironment",
		[]interface{}{account, region},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IpAddressManager) PrivateVpnPoolForRegion(region *string) ec2.IIpamPool {
	if err := i.validatePrivateVpnPoolForRegionParameters(region); err != nil {
		panic(err)
	}
	var returns ec2.IIpamPool

	_jsii_.Invoke(
		i,
		"privateVpnPoolForRegion",
		[]interface{}{region},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IpAddressManager) RegisterAccount(account *string, pool ec2.IIpamPool) {
	if err := i.validateRegisterAccountParameters(account, pool); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"registerAccount",
		[]interface{}{account, pool},
	)
}

func (i *jsiiProxy_IpAddressManager) RegisterCidr(scope constructs.IConstruct, id *string, cidr *string) {
	if err := i.validateRegisterCidrParameters(scope, id, cidr); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"registerCidr",
		[]interface{}{scope, id, cidr},
	)
}

func (i *jsiiProxy_IpAddressManager) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

