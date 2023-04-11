package k8saws

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/constructs-go/constructs/v10"
)

type IExternalDnsRegistry interface {
	Bind(scope constructs.IConstruct) *ExternalDnsRegistryConfiguration
	RegistryType() *string
}

// The jsii proxy for IExternalDnsRegistry
type jsiiProxy_IExternalDnsRegistry struct {
	_ byte // padding
}

func (i *jsiiProxy_IExternalDnsRegistry) Bind(scope constructs.IConstruct) *ExternalDnsRegistryConfiguration {
	if err := i.validateBindParameters(scope); err != nil {
		panic(err)
	}
	var returns *ExternalDnsRegistryConfiguration

	_jsii_.Invoke(
		i,
		"bind",
		[]interface{}{scope},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_IExternalDnsRegistry) RegistryType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"registryType",
		&returns,
	)
	return returns
}

