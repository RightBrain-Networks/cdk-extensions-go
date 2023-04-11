package k8saws

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/constructs-go/constructs/v10"
)

// Represents a resource the can be synchronized into a Kubernetes secret.
type ISecretReference interface {
	// Gets the configuration details for the resource being sychronized in a form that can be universally used to create the synchronization configuration.
	Bind(scope constructs.IConstruct) *SecretReferenceConfiguration
}

// The jsii proxy for ISecretReference
type jsiiProxy_ISecretReference struct {
	_ byte // padding
}

func (i *jsiiProxy_ISecretReference) Bind(scope constructs.IConstruct) *SecretReferenceConfiguration {
	if err := i.validateBindParameters(scope); err != nil {
		panic(err)
	}
	var returns *SecretReferenceConfiguration

	_jsii_.Invoke(
		i,
		"bind",
		[]interface{}{scope},
		&returns,
	)

	return returns
}

