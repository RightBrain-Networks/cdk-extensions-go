package k8saws

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/vibe-io/cdk-extensions-go/cdkextensions/jsii"

	"github.com/aws/constructs-go/constructs/v10"
)

// An ExternalDNS registry that tracks DNS record ownership information using DNS TXT records.
// See: [About TXT records](https://support.google.com/a/answer/2716800?hl=en)
//
type TxtRegistry interface {
	IExternalDnsRegistry
	// A unique identifier that is used to establish ownership of managed DNS records.
	//
	// Prevents conflicts in the event of multiple clusters running external-dns.
	OwnerId() *string
	// A prefix to be added top TXT ownership records.
	//
	// By default, the ownership record is a TXT record with the same name as the
	// managed record that was created. This causes issues as some record types
	// (CNAME's) do not allow duplicate records of a different type.
	//
	// This prefix is used to prevent such name collissions while still allowing
	// DNS ownership records to be created.
	Prefix() *string
	// The type name of ExternalDNS registry.
	RegistryType() *string
	// Generates an object with all the information needed to use the registry in a given CDK scope.
	//
	// Returns: A configuration object representing the implementation of this
	// registry.
	Bind(scope constructs.IConstruct) *ExternalDnsRegistryConfiguration
}

// The jsii proxy struct for TxtRegistry
type jsiiProxy_TxtRegistry struct {
	jsiiProxy_IExternalDnsRegistry
}

func (j *jsiiProxy_TxtRegistry) OwnerId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ownerId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TxtRegistry) Prefix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"prefix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TxtRegistry) RegistryType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"registryType",
		&returns,
	)
	return returns
}


// Creates a new instance of the NoopRegistry class.
func NewTxtRegistry(options *TxtRegistryOptions) TxtRegistry {
	_init_.Initialize()

	if err := validateNewTxtRegistryParameters(options); err != nil {
		panic(err)
	}
	j := jsiiProxy_TxtRegistry{}

	_jsii_.Create(
		"cdk-extensions.k8s_aws.TxtRegistry",
		[]interface{}{options},
		&j,
	)

	return &j
}

// Creates a new instance of the NoopRegistry class.
func NewTxtRegistry_Override(t TxtRegistry, options *TxtRegistryOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk-extensions.k8s_aws.TxtRegistry",
		[]interface{}{options},
		t,
	)
}

func TxtRegistry_DEFAULT_PREFIX() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"cdk-extensions.k8s_aws.TxtRegistry",
		"DEFAULT_PREFIX",
		&returns,
	)
	return returns
}

func TxtRegistry_NO_PREFIX() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"cdk-extensions.k8s_aws.TxtRegistry",
		"NO_PREFIX",
		&returns,
	)
	return returns
}

func (t *jsiiProxy_TxtRegistry) Bind(scope constructs.IConstruct) *ExternalDnsRegistryConfiguration {
	if err := t.validateBindParameters(scope); err != nil {
		panic(err)
	}
	var returns *ExternalDnsRegistryConfiguration

	_jsii_.Invoke(
		t,
		"bind",
		[]interface{}{scope},
		&returns,
	)

	return returns
}

