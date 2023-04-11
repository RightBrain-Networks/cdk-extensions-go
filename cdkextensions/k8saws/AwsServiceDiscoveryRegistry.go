package k8saws

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/vibe-io/cdk-extensions-go/cdkextensions/jsii"

	"github.com/aws/constructs-go/constructs/v10"
)

// An ExternalDNS registry that tracks DNS record ownership information using AWS Service Discovery.
// See: [AWS Cloud Map](https://docs.aws.amazon.com/cloud-map/latest/dg/what-is-cloud-map.html)
//
type AwsServiceDiscoveryRegistry interface {
	IExternalDnsRegistry
	// The type name of ExternalDNS registry.
	RegistryType() *string
	// Generates an object with all the information needed to use the registry in a given CDK scope.
	//
	// Returns: A configuration object representing the implementation of this
	// registry.
	Bind(scope constructs.IConstruct) *ExternalDnsRegistryConfiguration
}

// The jsii proxy struct for AwsServiceDiscoveryRegistry
type jsiiProxy_AwsServiceDiscoveryRegistry struct {
	jsiiProxy_IExternalDnsRegistry
}

func (j *jsiiProxy_AwsServiceDiscoveryRegistry) RegistryType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"registryType",
		&returns,
	)
	return returns
}


// Creates a new instance of the AwsServiceDiscoveryRegistry class.
func NewAwsServiceDiscoveryRegistry() AwsServiceDiscoveryRegistry {
	_init_.Initialize()

	j := jsiiProxy_AwsServiceDiscoveryRegistry{}

	_jsii_.Create(
		"cdk-extensions.k8s_aws.AwsServiceDiscoveryRegistry",
		nil, // no parameters
		&j,
	)

	return &j
}

// Creates a new instance of the AwsServiceDiscoveryRegistry class.
func NewAwsServiceDiscoveryRegistry_Override(a AwsServiceDiscoveryRegistry) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk-extensions.k8s_aws.AwsServiceDiscoveryRegistry",
		nil, // no parameters
		a,
	)
}

func (a *jsiiProxy_AwsServiceDiscoveryRegistry) Bind(scope constructs.IConstruct) *ExternalDnsRegistryConfiguration {
	if err := a.validateBindParameters(scope); err != nil {
		panic(err)
	}
	var returns *ExternalDnsRegistryConfiguration

	_jsii_.Invoke(
		a,
		"bind",
		[]interface{}{scope},
		&returns,
	)

	return returns
}

