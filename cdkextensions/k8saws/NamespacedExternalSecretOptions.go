package k8saws


// Configuration options for adding a Kubernetes secret synced from an external provider to a specific Kubernetes namespace.
type NamespacedExternalSecretOptions struct {
	// A collection of field mappings that tells the external secrets operator the structure of the Kubernetes secret to create and which how fields in the Kubernetes secret should map to fields in the secret from the external secret provider.
	Fields *[]*SecretFieldReference `field:"optional" json:"fields" yaml:"fields"`
	// The name of the Kubernetes secret that will be created, as it will appear from within the Kubernetes cluster.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The Kubernetes namespace where the synced secret should be created.
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
}

