package k8saws

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/vibe-io/cdk-extensions-go/cdkextensions/jsii"

	"github.com/aws/constructs-go/constructs/v10"
)

// A placeholder ExternalDNS registry that says ExternalDNS should use not use a registry.
//
// When configuring ExternalDNS without a registry, the service has no idea the
// original creator and maintainer of DNS records. This means that there are
// likely to be conflicts if there are multiple services that could create or
// change DNS records in the same zone.
type NoopRegistry interface {
	IExternalDnsRegistry
	// The type name of ExternalDNS registry.
	RegistryType() *string
	// Generates an object with all the information needed to use the registry in a given CDK scope.
	//
	// Returns: A configuration object representing the implementation of this
	// registry.
	Bind(_scope constructs.IConstruct) *ExternalDnsRegistryConfiguration
}

// The jsii proxy struct for NoopRegistry
type jsiiProxy_NoopRegistry struct {
	jsiiProxy_IExternalDnsRegistry
}

func (j *jsiiProxy_NoopRegistry) RegistryType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"registryType",
		&returns,
	)
	return returns
}


// Creates a new instance of the NoopRegistry class.
func NewNoopRegistry() NoopRegistry {
	_init_.Initialize()

	j := jsiiProxy_NoopRegistry{}

	_jsii_.Create(
		"cdk-extensions.k8s_aws.NoopRegistry",
		nil, // no parameters
		&j,
	)

	return &j
}

// Creates a new instance of the NoopRegistry class.
func NewNoopRegistry_Override(n NoopRegistry) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk-extensions.k8s_aws.NoopRegistry",
		nil, // no parameters
		n,
	)
}

func (n *jsiiProxy_NoopRegistry) Bind(_scope constructs.IConstruct) *ExternalDnsRegistryConfiguration {
	if err := n.validateBindParameters(_scope); err != nil {
		panic(err)
	}
	var returns *ExternalDnsRegistryConfiguration

	_jsii_.Invoke(
		n,
		"bind",
		[]interface{}{_scope},
		&returns,
	)

	return returns
}

