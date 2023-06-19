package resourcegroups

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/vibe-io/cdk-extensions-go/cdkextensions/jsii"

	"github.com/aws/constructs-go/constructs/v10"
)

// Configuration object for a Resource Group whose resources mirror those controlled by a CloudFormation stack.
type CloudFormationStack interface {
	IGroupConfiguration
	// The details of the CloudFormation stack that is referenced to create the Resource Group.
	Reference() IStackReference
	// Collection of resource types that are allowed to be in the Resource Group being configured.
	ResourceTypes() *[]*string
	// Add a resource type to the resource group.
	//
	// If no resource types are registered in the configuration then all resource
	// types are allowed.
	//
	// Returns: The Resource Group configuration object to which the type was
	// registered.
	AddResourceType(typeId *string) CloudFormationStack
	// Associates this configuration with a construct that is handling the creation of a resource group.
	//
	// Returns: The configuration to be used for the Resource Group.
	Bind(_scope constructs.IConstruct) *BoundGroupConfiguration
}

// The jsii proxy struct for CloudFormationStack
type jsiiProxy_CloudFormationStack struct {
	jsiiProxy_IGroupConfiguration
}

func (j *jsiiProxy_CloudFormationStack) Reference() IStackReference {
	var returns IStackReference
	_jsii_.Get(
		j,
		"reference",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudFormationStack) ResourceTypes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"resourceTypes",
		&returns,
	)
	return returns
}


// Creates a new instance of the CloudFormationStack class.
func NewCloudFormationStack(reference IStackReference, props *CloudFormationStackProps) CloudFormationStack {
	_init_.Initialize()

	if err := validateNewCloudFormationStackParameters(reference, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_CloudFormationStack{}

	_jsii_.Create(
		"cdk-extensions.resourcegroups.CloudFormationStack",
		[]interface{}{reference, props},
		&j,
	)

	return &j
}

// Creates a new instance of the CloudFormationStack class.
func NewCloudFormationStack_Override(c CloudFormationStack, reference IStackReference, props *CloudFormationStackProps) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk-extensions.resourcegroups.CloudFormationStack",
		[]interface{}{reference, props},
		c,
	)
}

func (c *jsiiProxy_CloudFormationStack) AddResourceType(typeId *string) CloudFormationStack {
	if err := c.validateAddResourceTypeParameters(typeId); err != nil {
		panic(err)
	}
	var returns CloudFormationStack

	_jsii_.Invoke(
		c,
		"addResourceType",
		[]interface{}{typeId},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudFormationStack) Bind(_scope constructs.IConstruct) *BoundGroupConfiguration {
	if err := c.validateBindParameters(_scope); err != nil {
		panic(err)
	}
	var returns *BoundGroupConfiguration

	_jsii_.Invoke(
		c,
		"bind",
		[]interface{}{_scope},
		&returns,
	)

	return returns
}

