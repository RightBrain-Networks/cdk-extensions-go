package k8saws


// Configuration options for referencing a Secrets Manager secret as a Kubernetes secret.
type SecretsManagerReferenceOptions struct {
	// Defines a mapping of how JSON keys in the Secrets Manager secret should appear in the imported Kubernetes secret.
	Fields *[]*SecretFieldReference `field:"optional" json:"fields" yaml:"fields"`
}

