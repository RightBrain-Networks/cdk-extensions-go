package k8saws

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/vibe-io/cdk-extensions-go/cdkextensions/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsssm"
	"github.com/aws/constructs-go/constructs/v10"
)

// Defines a reference for importing and synchronizing an SSM parameter to a Kubernetes secret.
type SsmParameterReference interface {
	ISecretReference
	// An array of field mappings which will be applied to this secret reference when mapping keys from SSM parameter JSON objects to keys in the imported secret.
	Fields() *[]*SecretFieldReference
	// The SSM parameter being referenced to import into Kubernetes.
	Parameter() awsssm.IParameter
	// Adds a field mapping that specifies how a field from an SSM JSON parameter should be mapped into the imported Kubernetes secret.
	//
	// Returns: The `SsmParameterReference` where the mapping was added.
	AddFieldMapping(field *SecretFieldReference) SsmParameterReference
	// Binds the reference to an object that is in charge of generating the manifest for the external secret.
	//
	// Returns: A configuration object providing the details needed to build
	// the external secret Kubernetes resource.
	Bind(_scope constructs.IConstruct) *SecretReferenceConfiguration
}

// The jsii proxy struct for SsmParameterReference
type jsiiProxy_SsmParameterReference struct {
	jsiiProxy_ISecretReference
}

func (j *jsiiProxy_SsmParameterReference) Fields() *[]*SecretFieldReference {
	var returns *[]*SecretFieldReference
	_jsii_.Get(
		j,
		"fields",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmParameterReference) Parameter() awsssm.IParameter {
	var returns awsssm.IParameter
	_jsii_.Get(
		j,
		"parameter",
		&returns,
	)
	return returns
}


// Creates a new instance of the SsmParameterReference class.
func NewSsmParameterReference(parameter awsssm.IParameter, options *SsmParameterReferenceOptions) SsmParameterReference {
	_init_.Initialize()

	if err := validateNewSsmParameterReferenceParameters(parameter, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_SsmParameterReference{}

	_jsii_.Create(
		"cdk-extensions.k8s_aws.SsmParameterReference",
		[]interface{}{parameter, options},
		&j,
	)

	return &j
}

// Creates a new instance of the SsmParameterReference class.
func NewSsmParameterReference_Override(s SsmParameterReference, parameter awsssm.IParameter, options *SsmParameterReferenceOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk-extensions.k8s_aws.SsmParameterReference",
		[]interface{}{parameter, options},
		s,
	)
}

func (s *jsiiProxy_SsmParameterReference) AddFieldMapping(field *SecretFieldReference) SsmParameterReference {
	if err := s.validateAddFieldMappingParameters(field); err != nil {
		panic(err)
	}
	var returns SsmParameterReference

	_jsii_.Invoke(
		s,
		"addFieldMapping",
		[]interface{}{field},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SsmParameterReference) Bind(_scope constructs.IConstruct) *SecretReferenceConfiguration {
	if err := s.validateBindParameters(_scope); err != nil {
		panic(err)
	}
	var returns *SecretReferenceConfiguration

	_jsii_.Invoke(
		s,
		"bind",
		[]interface{}{_scope},
		&returns,
	)

	return returns
}

