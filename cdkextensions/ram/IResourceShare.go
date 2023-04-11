package ram

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/vibe-io/cdk-extensions-go/cdkextensions/ram/internal"
)

// Represents an AWS Resource Access Manager (RAM) resource share in AWS.
type IResourceShare interface {
	awscdk.IResource
	// Adds a new principal to the resource share.
	//
	// The principal will have access to all the resources associated with the
	// resource share.
	AddPrincipal(principal ISharedPrincipal)
	// Adds a new resource to the resource share.
	//
	// The resource will be accessible to all pricipals associated with the
	// resource share.
	AddResource(resource ISharable)
	// The Amazon Resource Name (ARN) of the RAM resource share.
	ResourceShareArn() *string
	// The ID of the RAM resource share.
	ResourceShareId() *string
}

// The jsii proxy for IResourceShare
type jsiiProxy_IResourceShare struct {
	internal.Type__awscdkIResource
}

func (i *jsiiProxy_IResourceShare) AddPrincipal(principal ISharedPrincipal) {
	if err := i.validateAddPrincipalParameters(principal); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"addPrincipal",
		[]interface{}{principal},
	)
}

func (i *jsiiProxy_IResourceShare) AddResource(resource ISharable) {
	if err := i.validateAddResourceParameters(resource); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"addResource",
		[]interface{}{resource},
	)
}

func (j *jsiiProxy_IResourceShare) ResourceShareArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceShareArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IResourceShare) ResourceShareId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceShareId",
		&returns,
	)
	return returns
}

