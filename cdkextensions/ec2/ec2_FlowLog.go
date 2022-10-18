package ec2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/vibe-io/cdk-extensions-go/cdkextensions/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/vibe-io/cdk-extensions-go/cdkextensions/ec2/internal"
)

type FlowLog interface {
	awscdk.Resource
	awsec2.IFlowLog
	// The location where flow logs should be delivered.
	// See: [FlowLog LogDestinationType](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-flowlog.html#cfn-ec2-flowlog-logdestinationtype)
	//
	Destination() FlowLogDestination
	// The environment this resource belongs to.
	//
	// For resources that are created and managed by the CDK
	// (generally, those created by creating new class instances like Role, Bucket, etc.),
	// this is always the same as the environment of the stack they belong to;
	// however, for imported resources
	// (those obtained from static methods like fromRoleArn, fromBucketName, etc.),
	// that might be different than the stack they were imported into.
	Env() *awscdk.ResourceEnvironment
	// The Amazon Resource Name (ARN) of the flow log.
	FlowLogArn() *string
	// The ID of the flow log.
	FlowLogId() *string
	// The fields to include in the flow log record, in the order in which they should appear.
	//
	// For a list of available fields, see {@link FlowLogField}.
	// See: [FlowLog LogFormat](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-flowlog.html#cfn-ec2-flowlog-logformat)
	//
	Format() FlowLogFormat
	// The maximum interval of time during which a flow of packets is captured and aggregated into a flow log record.
	// See: [FlowLog MaxAggregationInterval](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-flowlog.html#cfn-ec2-flowlog-maxaggregationinterval)
	//
	MaxAggregationInterval() FlowLogAggregationInterval
	// The tree node.
	Node() constructs.Node
	// Returns a string-encoded token that resolves to the physical name that should be passed to the CloudFormation resource.
	//
	// This value will resolve to one of the following:
	// - a concrete value (e.g. `"my-awesome-bucket"`)
	// - `undefined`, when a name should be generated by CloudFormation
	// - a concrete name generated automatically during synthesis, in
	//    cross-environment scenarios.
	PhysicalName() *string
	// The underlying FlowLog CloudFormation resource.
	// See: [AWS::EC2::FlowLog](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-flowlog.html)
	//
	Resource() awsec2.CfnFlowLog
	// Details for the resource from which flow logs will be captured.
	// See: [FlowLog ResourceType](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-flowlog.html#cfn-ec2-flowlog-resourcetype)
	//
	ResourceType() awsec2.FlowLogResourceType
	// The stack in which this resource is defined.
	Stack() awscdk.Stack
	// The type of traffic to monitor (accepted traffic, rejected traffic, or all traffic).
	// See: [FlowLog TrafficType](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-flowlog.html#cfn-ec2-flowlog-traffictype)
	//
	TrafficType() awsec2.FlowLogTrafficType
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

// The jsii proxy struct for FlowLog
type jsiiProxy_FlowLog struct {
	internal.Type__awscdkResource
	internal.Type__awsec2IFlowLog
}

func (j *jsiiProxy_FlowLog) Destination() FlowLogDestination {
	var returns FlowLogDestination
	_jsii_.Get(
		j,
		"destination",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FlowLog) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FlowLog) FlowLogArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"flowLogArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FlowLog) FlowLogId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"flowLogId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FlowLog) Format() FlowLogFormat {
	var returns FlowLogFormat
	_jsii_.Get(
		j,
		"format",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FlowLog) MaxAggregationInterval() FlowLogAggregationInterval {
	var returns FlowLogAggregationInterval
	_jsii_.Get(
		j,
		"maxAggregationInterval",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FlowLog) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FlowLog) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FlowLog) Resource() awsec2.CfnFlowLog {
	var returns awsec2.CfnFlowLog
	_jsii_.Get(
		j,
		"resource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FlowLog) ResourceType() awsec2.FlowLogResourceType {
	var returns awsec2.FlowLogResourceType
	_jsii_.Get(
		j,
		"resourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FlowLog) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FlowLog) TrafficType() awsec2.FlowLogTrafficType {
	var returns awsec2.FlowLogTrafficType
	_jsii_.Get(
		j,
		"trafficType",
		&returns,
	)
	return returns
}


// Creates a new instance of the FlowLog class.
func NewFlowLog(scope constructs.IConstruct, id *string, props *FlowLogProps) FlowLog {
	_init_.Initialize()

	if err := validateNewFlowLogParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_FlowLog{}

	_jsii_.Create(
		"cdk-extensions.ec2.FlowLog",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Creates a new instance of the FlowLog class.
func NewFlowLog_Override(f FlowLog, scope constructs.IConstruct, id *string, props *FlowLogProps) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk-extensions.ec2.FlowLog",
		[]interface{}{scope, id, props},
		f,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead.
func FlowLog_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateFlowLog_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"cdk-extensions.ec2.FlowLog",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Returns true if the construct was created by CDK, and false otherwise.
func FlowLog_IsOwnedResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	if err := validateFlowLog_IsOwnedResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"cdk-extensions.ec2.FlowLog",
		"isOwnedResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
func FlowLog_IsResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	if err := validateFlowLog_IsResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"cdk-extensions.ec2.FlowLog",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FlowLog) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	if err := f.validateApplyRemovalPolicyParameters(policy); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (f *jsiiProxy_FlowLog) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"generatePhysicalName",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FlowLog) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	if err := f.validateGetResourceArnAttributeParameters(arnAttr, arnComponents); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		f,
		"getResourceArnAttribute",
		[]interface{}{arnAttr, arnComponents},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FlowLog) GetResourceNameAttribute(nameAttr *string) *string {
	if err := f.validateGetResourceNameAttributeParameters(nameAttr); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		f,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FlowLog) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

