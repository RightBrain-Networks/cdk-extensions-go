package k8saws


// Options for how to synchronize a specific field in a secret being imported.
type SecretFieldReference struct {
	// The name of the data key to be used for the field in the imported Kubernetes secret.
	KubernetesKey *string `field:"required" json:"kubernetesKey" yaml:"kubernetesKey"`
	// Policy for fetching tags/labels from provider secrets.
	MetadataPolicy MetadataPolicy `field:"optional" json:"metadataPolicy" yaml:"metadataPolicy"`
	// The JSON key for the field in the secret being imported.
	RemoteKey *string `field:"optional" json:"remoteKey" yaml:"remoteKey"`
}

