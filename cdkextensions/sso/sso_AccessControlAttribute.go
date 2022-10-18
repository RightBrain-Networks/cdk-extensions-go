package sso

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/vibe-io/cdk-extensions-go/cdkextensions/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2/awssso"
	"github.com/aws/constructs-go/constructs/v10"
)

// Represents and ABAC attribute in IAM Identity Center.
//
// These are IAM Identity Center identity store attributes that you can
// configure for use in attributes-based access control (ABAC). You can create
// permissions policies that determine who can access your AWS resources based
// upon the configured attribute values. When you enable ABAC and specify
// `AccessControlAttributes`, IAM Identity Center passes the attribute values
// of the authenticated user into IAM for use in policy evaluation.
type AccessControlAttribute interface {
	// The name of the attribute associated with your identities in your identity source.
	//
	// This is used to map a specified attribute in your
	// identity source with an attribute in IAM Identity Center.
	Name() *string
	// A list of identity sources to use when mapping a specified attribute to IAM Identity Center.
	//
	// Note that the array is readonly and changes made
	// to it will not be reflected when generating ABAC attribute
	// configuration. To add a source to the attribute use the {@link addSource}
	// method.
	Sources() *[]*string
	// Adds an identity source to use when mapping the attribute to IAM Identity Center.
	//
	// Returns: The ABAC attribute the source was associated with.
	AddSource(source *string) AccessControlAttribute
	// Generates the raw CloudFormation configuration that this attribute represents within the context of a given scope.
	//
	// Returns: The raw CloudFormation configuration that this attribute
	// represents.
	Bind(scope constructs.IConstruct) *awssso.CfnInstanceAccessControlAttributeConfiguration_AccessControlAttributeProperty
}

// The jsii proxy struct for AccessControlAttribute
type jsiiProxy_AccessControlAttribute struct {
	_ byte // padding
}

func (j *jsiiProxy_AccessControlAttribute) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessControlAttribute) Sources() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"sources",
		&returns,
	)
	return returns
}


// Creates a new instance of the AccessControlAttribute class.
func NewAccessControlAttribute(options *AccessControlAttributeOptions) AccessControlAttribute {
	_init_.Initialize()

	if err := validateNewAccessControlAttributeParameters(options); err != nil {
		panic(err)
	}
	j := jsiiProxy_AccessControlAttribute{}

	_jsii_.Create(
		"cdk-extensions.sso.AccessControlAttribute",
		[]interface{}{options},
		&j,
	)

	return &j
}

// Creates a new instance of the AccessControlAttribute class.
func NewAccessControlAttribute_Override(a AccessControlAttribute, options *AccessControlAttributeOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk-extensions.sso.AccessControlAttribute",
		[]interface{}{options},
		a,
	)
}

func (a *jsiiProxy_AccessControlAttribute) AddSource(source *string) AccessControlAttribute {
	if err := a.validateAddSourceParameters(source); err != nil {
		panic(err)
	}
	var returns AccessControlAttribute

	_jsii_.Invoke(
		a,
		"addSource",
		[]interface{}{source},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AccessControlAttribute) Bind(scope constructs.IConstruct) *awssso.CfnInstanceAccessControlAttributeConfiguration_AccessControlAttributeProperty {
	if err := a.validateBindParameters(scope); err != nil {
		panic(err)
	}
	var returns *awssso.CfnInstanceAccessControlAttributeConfiguration_AccessControlAttributeProperty

	_jsii_.Invoke(
		a,
		"bind",
		[]interface{}{scope},
		&returns,
	)

	return returns
}

