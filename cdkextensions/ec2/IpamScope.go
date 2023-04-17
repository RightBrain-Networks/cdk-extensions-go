package ec2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/vibe-io/cdk-extensions-go/cdkextensions/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/constructs-go/constructs/v10"
)

type IpamScope interface {
}

// The jsii proxy struct for IpamScope
type jsiiProxy_IpamScope struct {
	_ byte // padding
}

func NewIpamScope() IpamScope {
	_init_.Initialize()

	j := jsiiProxy_IpamScope{}

	_jsii_.Create(
		"cdk-extensions.ec2.IpamScope",
		nil, // no parameters
		&j,
	)

	return &j
}

func NewIpamScope_Override(i IpamScope) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk-extensions.ec2.IpamScope",
		nil, // no parameters
		i,
	)
}

// Imports an existing IPAM scope by specifying its Amazon Resource Name (ARN).
//
// Returns: An object representing the imported IPAM scope.
func IpamScope_FromIpamScopeArn(scope constructs.IConstruct, id *string, ipamScopeArn *string) IIpamScope {
	_init_.Initialize()

	if err := validateIpamScope_FromIpamScopeArnParameters(scope, id, ipamScopeArn); err != nil {
		panic(err)
	}
	var returns IIpamScope

	_jsii_.StaticInvoke(
		"cdk-extensions.ec2.IpamScope",
		"fromIpamScopeArn",
		[]interface{}{scope, id, ipamScopeArn},
		&returns,
	)

	return returns
}

// Imports an existing IAPM scope by explicitly specifying its attributes.
//
// Returns: An object representing the imported IPAM scope.
func IpamScope_FromIpamScopeAttributes(scope constructs.IConstruct, id *string, attrs *IpamScopeAttributes) IIpamScope {
	_init_.Initialize()

	if err := validateIpamScope_FromIpamScopeAttributesParameters(scope, id, attrs); err != nil {
		panic(err)
	}
	var returns IIpamScope

	_jsii_.StaticInvoke(
		"cdk-extensions.ec2.IpamScope",
		"fromIpamScopeAttributes",
		[]interface{}{scope, id, attrs},
		&returns,
	)

	return returns
}

// Imports an existing IPAM scope by explicitly specifying its AWS generated ID.
//
// Returns: An object representing the imported IPAM scope.
func IpamScope_FromIpamScopeId(scope constructs.IConstruct, id *string, ipamScopeId *string) IIpamScope {
	_init_.Initialize()

	if err := validateIpamScope_FromIpamScopeIdParameters(scope, id, ipamScopeId); err != nil {
		panic(err)
	}
	var returns IIpamScope

	_jsii_.StaticInvoke(
		"cdk-extensions.ec2.IpamScope",
		"fromIpamScopeId",
		[]interface{}{scope, id, ipamScopeId},
		&returns,
	)

	return returns
}

func IpamScope_ARN_FORMAT() awscdk.ArnFormat {
	_init_.Initialize()
	var returns awscdk.ArnFormat
	_jsii_.StaticGet(
		"cdk-extensions.ec2.IpamScope",
		"ARN_FORMAT",
		&returns,
	)
	return returns
}

