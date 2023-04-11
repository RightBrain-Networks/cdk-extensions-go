package sso

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/vibe-io/cdk-extensions-go/cdkextensions/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssso"
	"github.com/aws/constructs-go/constructs/v10"
)

type ReferencedPermissionsBoundary interface {
	PermissionsBoundary
	ReferencedPolicy() ReferencedManagedPolicy
	Bind(_scope constructs.IConstruct) *awssso.CfnPermissionSet_PermissionsBoundaryProperty
}

// The jsii proxy struct for ReferencedPermissionsBoundary
type jsiiProxy_ReferencedPermissionsBoundary struct {
	jsiiProxy_PermissionsBoundary
}

func (j *jsiiProxy_ReferencedPermissionsBoundary) ReferencedPolicy() ReferencedManagedPolicy {
	var returns ReferencedManagedPolicy
	_jsii_.Get(
		j,
		"referencedPolicy",
		&returns,
	)
	return returns
}


func NewReferencedPermissionsBoundary(options *ReferenceOptions) ReferencedPermissionsBoundary {
	_init_.Initialize()

	if err := validateNewReferencedPermissionsBoundaryParameters(options); err != nil {
		panic(err)
	}
	j := jsiiProxy_ReferencedPermissionsBoundary{}

	_jsii_.Create(
		"cdk-extensions.sso.ReferencedPermissionsBoundary",
		[]interface{}{options},
		&j,
	)

	return &j
}

func NewReferencedPermissionsBoundary_Override(r ReferencedPermissionsBoundary, options *ReferenceOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk-extensions.sso.ReferencedPermissionsBoundary",
		[]interface{}{options},
		r,
	)
}

func ReferencedPermissionsBoundary_FromManagedPolicy(policy awsiam.IManagedPolicy) ManagedPolicyPermissionsBoundary {
	_init_.Initialize()

	if err := validateReferencedPermissionsBoundary_FromManagedPolicyParameters(policy); err != nil {
		panic(err)
	}
	var returns ManagedPolicyPermissionsBoundary

	_jsii_.StaticInvoke(
		"cdk-extensions.sso.ReferencedPermissionsBoundary",
		"fromManagedPolicy",
		[]interface{}{policy},
		&returns,
	)

	return returns
}

func ReferencedPermissionsBoundary_FromReference(options *ReferenceOptions) ReferencedPermissionsBoundary {
	_init_.Initialize()

	if err := validateReferencedPermissionsBoundary_FromReferenceParameters(options); err != nil {
		panic(err)
	}
	var returns ReferencedPermissionsBoundary

	_jsii_.StaticInvoke(
		"cdk-extensions.sso.ReferencedPermissionsBoundary",
		"fromReference",
		[]interface{}{options},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ReferencedPermissionsBoundary) Bind(_scope constructs.IConstruct) *awssso.CfnPermissionSet_PermissionsBoundaryProperty {
	if err := r.validateBindParameters(_scope); err != nil {
		panic(err)
	}
	var returns *awssso.CfnPermissionSet_PermissionsBoundaryProperty

	_jsii_.Invoke(
		r,
		"bind",
		[]interface{}{_scope},
		&returns,
	)

	return returns
}

