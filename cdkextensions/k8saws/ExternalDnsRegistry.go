package k8saws

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/vibe-io/cdk-extensions-go/cdkextensions/jsii"
)

// Helper class that provides access to the available ExternalDns registry options.
type ExternalDnsRegistry interface {
}

// The jsii proxy struct for ExternalDnsRegistry
type jsiiProxy_ExternalDnsRegistry struct {
	_ byte // padding
}

func NewExternalDnsRegistry() ExternalDnsRegistry {
	_init_.Initialize()

	j := jsiiProxy_ExternalDnsRegistry{}

	_jsii_.Create(
		"cdk-extensions.k8s_aws.ExternalDnsRegistry",
		nil, // no parameters
		&j,
	)

	return &j
}

func NewExternalDnsRegistry_Override(e ExternalDnsRegistry) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk-extensions.k8s_aws.ExternalDnsRegistry",
		nil, // no parameters
		e,
	)
}

// An ExternalDNS registry that tracks DNS record ownership information using AWS Service Discovery.
//
// Returns: A ExternalDNS registry object configured to use AWS Cloud Map
// for ownership information.
// See: [AWS Cloud Map](https://docs.aws.amazon.com/cloud-map/latest/dg/what-is-cloud-map.html)
//
func ExternalDnsRegistry_AwsServiceDiscovery() AwsServiceDiscoveryRegistry {
	_init_.Initialize()

	var returns AwsServiceDiscoveryRegistry

	_jsii_.StaticInvoke(
		"cdk-extensions.k8s_aws.ExternalDnsRegistry",
		"awsServiceDiscovery",
		nil, // no parameters
		&returns,
	)

	return returns
}

// A placeholder ExternalDNS registry that says ExternalDNS should use not use a registry.
//
// When configuring ExternalDNS without a registry, the service has no idea
// the original creator and maintainer of DNS records. This means that
// there are likely to be conflicts if there are multiple services that
// could create or change DNS records in the same zone.
//
// Returns: An object that instructs ExternalDNS to not store record
// ownership information and will perform record updates without
// validation.
func ExternalDnsRegistry_Noop() NoopRegistry {
	_init_.Initialize()

	var returns NoopRegistry

	_jsii_.StaticInvoke(
		"cdk-extensions.k8s_aws.ExternalDnsRegistry",
		"noop",
		nil, // no parameters
		&returns,
	)

	return returns
}

// An ExternalDNS registry that tracks DNS record ownership information using DNS TXT records.
//
// Returns: A ExternalDNS registry object configured to use DNS TXT records
// for ownership information.
// See: [About TXT records](https://support.google.com/a/answer/2716800?hl=en)
//
func ExternalDnsRegistry_Txt(options *TxtRegistryOptions) TxtRegistry {
	_init_.Initialize()

	if err := validateExternalDnsRegistry_TxtParameters(options); err != nil {
		panic(err)
	}
	var returns TxtRegistry

	_jsii_.StaticInvoke(
		"cdk-extensions.k8s_aws.ExternalDnsRegistry",
		"txt",
		[]interface{}{options},
		&returns,
	)

	return returns
}

