package ec2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/vibe-io/cdk-extensions-go/cdkextensions/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/constructs-go/constructs/v10"
)

type PublicIpamScope interface {
}

// The jsii proxy struct for PublicIpamScope
type jsiiProxy_PublicIpamScope struct {
	_ byte // padding
}

func NewPublicIpamScope() PublicIpamScope {
	_init_.Initialize()

	j := jsiiProxy_PublicIpamScope{}

	_jsii_.Create(
		"cdk-extensions.ec2.PublicIpamScope",
		nil, // no parameters
		&j,
	)

	return &j
}

func NewPublicIpamScope_Override(p PublicIpamScope) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk-extensions.ec2.PublicIpamScope",
		nil, // no parameters
		p,
	)
}

// Imports an existing IPAM scope by specifying its Amazon Resource Name (ARN).
//
// Returns: An object representing the imported IPAM scope.
func PublicIpamScope_FromIpamScopeArn(scope constructs.IConstruct, id *string, ipamScopeArn *string) IPublicIpamScope {
	_init_.Initialize()

	if err := validatePublicIpamScope_FromIpamScopeArnParameters(scope, id, ipamScopeArn); err != nil {
		panic(err)
	}
	var returns IPublicIpamScope

	_jsii_.StaticInvoke(
		"cdk-extensions.ec2.PublicIpamScope",
		"fromIpamScopeArn",
		[]interface{}{scope, id, ipamScopeArn},
		&returns,
	)

	return returns
}

// Imports an existing IAPM scope by explicitly specifying its attributes.
//
// Returns: An object representing the imported IPAM scope.
func PublicIpamScope_FromIpamScopeAttributes(scope constructs.IConstruct, id *string, attrs *IpamScopeAttributes) IPublicIpamScope {
	_init_.Initialize()

	if err := validatePublicIpamScope_FromIpamScopeAttributesParameters(scope, id, attrs); err != nil {
		panic(err)
	}
	var returns IPublicIpamScope

	_jsii_.StaticInvoke(
		"cdk-extensions.ec2.PublicIpamScope",
		"fromIpamScopeAttributes",
		[]interface{}{scope, id, attrs},
		&returns,
	)

	return returns
}

// Imports an existing IPAM scope by explicitly specifying its AWS generated ID.
//
// Returns: An object representing the imported IPAM scope.
func PublicIpamScope_FromIpamScopeId(scope constructs.IConstruct, id *string, ipamScopeId *string) IPublicIpamScope {
	_init_.Initialize()

	if err := validatePublicIpamScope_FromIpamScopeIdParameters(scope, id, ipamScopeId); err != nil {
		panic(err)
	}
	var returns IPublicIpamScope

	_jsii_.StaticInvoke(
		"cdk-extensions.ec2.PublicIpamScope",
		"fromIpamScopeId",
		[]interface{}{scope, id, ipamScopeId},
		&returns,
	)

	return returns
}

func PublicIpamScope_ARN_FORMAT() awscdk.ArnFormat {
	_init_.Initialize()
	var returns awscdk.ArnFormat
	_jsii_.StaticGet(
		"cdk-extensions.ec2.PublicIpamScope",
		"ARN_FORMAT",
		&returns,
	)
	return returns
}

