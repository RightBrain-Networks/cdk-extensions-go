package k8saws

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/vibe-io/cdk-extensions-go/cdkextensions/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2/awssecretsmanager"
	"github.com/aws/constructs-go/constructs/v10"
)

// Defines a reference for importing and synchronizing a Secrets Manager secret to a Kubernetes secret.
type SecretsManagerReference interface {
	ISecretReference
	// An array of field mappings which will be applied to this secret reference when mapping keys from SecretsManager JSON objects to keys in the imported secret.
	Fields() *[]*SecretFieldReference
	// The secret being referenced to import into Kubernetes.
	Secret() awssecretsmanager.ISecret
	// Adds a field mapping that specifies how a field from a Secrets Manager JSON secret should be mapped into the imported Kubernetes secret.
	//
	// Returns: The `SecretsManagerReference` where the mapping was added.
	AddFieldMapping(field *SecretFieldReference) SecretsManagerReference
	// Binds the reference to an object that is in charge of generating the manifest for the external secret.
	//
	// Returns: A configuration object providing the details needed to build
	// the external secret Kubernetes resource.
	Bind(_scope constructs.IConstruct) *SecretReferenceConfiguration
}

// The jsii proxy struct for SecretsManagerReference
type jsiiProxy_SecretsManagerReference struct {
	jsiiProxy_ISecretReference
}

func (j *jsiiProxy_SecretsManagerReference) Fields() *[]*SecretFieldReference {
	var returns *[]*SecretFieldReference
	_jsii_.Get(
		j,
		"fields",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecretsManagerReference) Secret() awssecretsmanager.ISecret {
	var returns awssecretsmanager.ISecret
	_jsii_.Get(
		j,
		"secret",
		&returns,
	)
	return returns
}


// Creates a new instance of the SecretsManagerReference class.
func NewSecretsManagerReference(secret awssecretsmanager.ISecret, options *SecretsManagerReferenceOptions) SecretsManagerReference {
	_init_.Initialize()

	if err := validateNewSecretsManagerReferenceParameters(secret, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_SecretsManagerReference{}

	_jsii_.Create(
		"cdk-extensions.k8s_aws.SecretsManagerReference",
		[]interface{}{secret, options},
		&j,
	)

	return &j
}

// Creates a new instance of the SecretsManagerReference class.
func NewSecretsManagerReference_Override(s SecretsManagerReference, secret awssecretsmanager.ISecret, options *SecretsManagerReferenceOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk-extensions.k8s_aws.SecretsManagerReference",
		[]interface{}{secret, options},
		s,
	)
}

func (s *jsiiProxy_SecretsManagerReference) AddFieldMapping(field *SecretFieldReference) SecretsManagerReference {
	if err := s.validateAddFieldMappingParameters(field); err != nil {
		panic(err)
	}
	var returns SecretsManagerReference

	_jsii_.Invoke(
		s,
		"addFieldMapping",
		[]interface{}{field},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SecretsManagerReference) Bind(_scope constructs.IConstruct) *SecretReferenceConfiguration {
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

