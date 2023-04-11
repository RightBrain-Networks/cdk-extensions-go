package ram

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/vibe-io/cdk-extensions-go/cdkextensions/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/constructs-go/constructs/v10"
)

type SharedPrincipal interface {
	ISharedPrincipal
	Bind(_scope constructs.IConstruct) *string
}

// The jsii proxy struct for SharedPrincipal
type jsiiProxy_SharedPrincipal struct {
	jsiiProxy_ISharedPrincipal
}

func SharedPrincipal_FromAccountId(account *string) SharedPrincipal {
	_init_.Initialize()

	if err := validateSharedPrincipal_FromAccountIdParameters(account); err != nil {
		panic(err)
	}
	var returns SharedPrincipal

	_jsii_.StaticInvoke(
		"cdk-extensions.ram.SharedPrincipal",
		"fromAccountId",
		[]interface{}{account},
		&returns,
	)

	return returns
}

func SharedPrincipal_FromConstruct(construct constructs.IConstruct) SharedPrincipal {
	_init_.Initialize()

	if err := validateSharedPrincipal_FromConstructParameters(construct); err != nil {
		panic(err)
	}
	var returns SharedPrincipal

	_jsii_.StaticInvoke(
		"cdk-extensions.ram.SharedPrincipal",
		"fromConstruct",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

func SharedPrincipal_FromOrganizationalUnitArn(arn *string) SharedPrincipal {
	_init_.Initialize()

	if err := validateSharedPrincipal_FromOrganizationalUnitArnParameters(arn); err != nil {
		panic(err)
	}
	var returns SharedPrincipal

	_jsii_.StaticInvoke(
		"cdk-extensions.ram.SharedPrincipal",
		"fromOrganizationalUnitArn",
		[]interface{}{arn},
		&returns,
	)

	return returns
}

func SharedPrincipal_FromOrganizationArn(arn *string) SharedPrincipal {
	_init_.Initialize()

	if err := validateSharedPrincipal_FromOrganizationArnParameters(arn); err != nil {
		panic(err)
	}
	var returns SharedPrincipal

	_jsii_.StaticInvoke(
		"cdk-extensions.ram.SharedPrincipal",
		"fromOrganizationArn",
		[]interface{}{arn},
		&returns,
	)

	return returns
}

func SharedPrincipal_FromRole(role awsiam.IRole) SharedPrincipal {
	_init_.Initialize()

	if err := validateSharedPrincipal_FromRoleParameters(role); err != nil {
		panic(err)
	}
	var returns SharedPrincipal

	_jsii_.StaticInvoke(
		"cdk-extensions.ram.SharedPrincipal",
		"fromRole",
		[]interface{}{role},
		&returns,
	)

	return returns
}

func SharedPrincipal_FromStage(stage awscdk.Stage) SharedPrincipal {
	_init_.Initialize()

	if err := validateSharedPrincipal_FromStageParameters(stage); err != nil {
		panic(err)
	}
	var returns SharedPrincipal

	_jsii_.StaticInvoke(
		"cdk-extensions.ram.SharedPrincipal",
		"fromStage",
		[]interface{}{stage},
		&returns,
	)

	return returns
}

func SharedPrincipal_FromUser(user awsiam.IUser) SharedPrincipal {
	_init_.Initialize()

	if err := validateSharedPrincipal_FromUserParameters(user); err != nil {
		panic(err)
	}
	var returns SharedPrincipal

	_jsii_.StaticInvoke(
		"cdk-extensions.ram.SharedPrincipal",
		"fromUser",
		[]interface{}{user},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SharedPrincipal) Bind(_scope constructs.IConstruct) *string {
	if err := s.validateBindParameters(_scope); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		s,
		"bind",
		[]interface{}{_scope},
		&returns,
	)

	return returns
}

