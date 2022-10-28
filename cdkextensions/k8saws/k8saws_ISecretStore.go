package k8saws

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/vibe-io/cdk-extensions-go/cdkextensions/k8saws/internal"
)

// Represents a Kubernetes secret store resource.
type ISecretStore interface {
	constructs.IDependable
	// The name of the secret store as it appears in Kubernetes.
	SecretStoreName() *string
}

// The jsii proxy for ISecretStore
type jsiiProxy_ISecretStore struct {
	internal.Type__constructsIDependable
}

func (j *jsiiProxy_ISecretStore) SecretStoreName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secretStoreName",
		&returns,
	)
	return returns
}

