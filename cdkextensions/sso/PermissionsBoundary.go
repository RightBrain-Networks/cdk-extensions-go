package sso

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/vibe-io/cdk-extensions-go/cdkextensions/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssso"
	"github.com/aws/constructs-go/constructs/v10"
)

type PermissionsBoundary interface {
	Bind(scope constructs.IConstruct) *awssso.CfnPermissionSet_PermissionsBoundaryProperty
}

// The jsii proxy struct for PermissionsBoundary
type jsiiProxy_PermissionsBoundary struct {
	_ byte // padding
}

func NewPermissionsBoundary_Override(p PermissionsBoundary) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk-extensions.sso.PermissionsBoundary",
		nil, // no parameters
		p,
	)
}

func PermissionsBoundary_FromManagedPolicy(policy awsiam.IManagedPolicy) ManagedPolicyPermissionsBoundary {
	_init_.Initialize()

	if err := validatePermissionsBoundary_FromManagedPolicyParameters(policy); err != nil {
		panic(err)
	}
	var returns ManagedPolicyPermissionsBoundary

	_jsii_.StaticInvoke(
		"cdk-extensions.sso.PermissionsBoundary",
		"fromManagedPolicy",
		[]interface{}{policy},
		&returns,
	)

	return returns
}

func PermissionsBoundary_FromReference(options *ReferenceOptions) ReferencedPermissionsBoundary {
	_init_.Initialize()

	if err := validatePermissionsBoundary_FromReferenceParameters(options); err != nil {
		panic(err)
	}
	var returns ReferencedPermissionsBoundary

	_jsii_.StaticInvoke(
		"cdk-extensions.sso.PermissionsBoundary",
		"fromReference",
		[]interface{}{options},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PermissionsBoundary) Bind(scope constructs.IConstruct) *awssso.CfnPermissionSet_PermissionsBoundaryProperty {
	if err := p.validateBindParameters(scope); err != nil {
		panic(err)
	}
	var returns *awssso.CfnPermissionSet_PermissionsBoundaryProperty

	_jsii_.Invoke(
		p,
		"bind",
		[]interface{}{scope},
		&returns,
	)

	return returns
}

