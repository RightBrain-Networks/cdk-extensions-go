package sso

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/vibe-io/cdk-extensions-go/cdkextensions/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssso"
	"github.com/aws/constructs-go/constructs/v10"
)

type ManagedPolicyPermissionsBoundary interface {
	PermissionsBoundary
	ManagedPolicy() awsiam.IManagedPolicy
	Bind(_scope constructs.IConstruct) *awssso.CfnPermissionSet_PermissionsBoundaryProperty
}

// The jsii proxy struct for ManagedPolicyPermissionsBoundary
type jsiiProxy_ManagedPolicyPermissionsBoundary struct {
	jsiiProxy_PermissionsBoundary
}

func (j *jsiiProxy_ManagedPolicyPermissionsBoundary) ManagedPolicy() awsiam.IManagedPolicy {
	var returns awsiam.IManagedPolicy
	_jsii_.Get(
		j,
		"managedPolicy",
		&returns,
	)
	return returns
}


func NewManagedPolicyPermissionsBoundary(policy awsiam.IManagedPolicy) ManagedPolicyPermissionsBoundary {
	_init_.Initialize()

	if err := validateNewManagedPolicyPermissionsBoundaryParameters(policy); err != nil {
		panic(err)
	}
	j := jsiiProxy_ManagedPolicyPermissionsBoundary{}

	_jsii_.Create(
		"cdk-extensions.sso.ManagedPolicyPermissionsBoundary",
		[]interface{}{policy},
		&j,
	)

	return &j
}

func NewManagedPolicyPermissionsBoundary_Override(m ManagedPolicyPermissionsBoundary, policy awsiam.IManagedPolicy) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk-extensions.sso.ManagedPolicyPermissionsBoundary",
		[]interface{}{policy},
		m,
	)
}

func ManagedPolicyPermissionsBoundary_FromManagedPolicy(policy awsiam.IManagedPolicy) ManagedPolicyPermissionsBoundary {
	_init_.Initialize()

	if err := validateManagedPolicyPermissionsBoundary_FromManagedPolicyParameters(policy); err != nil {
		panic(err)
	}
	var returns ManagedPolicyPermissionsBoundary

	_jsii_.StaticInvoke(
		"cdk-extensions.sso.ManagedPolicyPermissionsBoundary",
		"fromManagedPolicy",
		[]interface{}{policy},
		&returns,
	)

	return returns
}

func ManagedPolicyPermissionsBoundary_FromReference(options *ReferenceOptions) ReferencedPermissionsBoundary {
	_init_.Initialize()

	if err := validateManagedPolicyPermissionsBoundary_FromReferenceParameters(options); err != nil {
		panic(err)
	}
	var returns ReferencedPermissionsBoundary

	_jsii_.StaticInvoke(
		"cdk-extensions.sso.ManagedPolicyPermissionsBoundary",
		"fromReference",
		[]interface{}{options},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagedPolicyPermissionsBoundary) Bind(_scope constructs.IConstruct) *awssso.CfnPermissionSet_PermissionsBoundaryProperty {
	if err := m.validateBindParameters(_scope); err != nil {
		panic(err)
	}
	var returns *awssso.CfnPermissionSet_PermissionsBoundaryProperty

	_jsii_.Invoke(
		m,
		"bind",
		[]interface{}{_scope},
		&returns,
	)

	return returns
}

