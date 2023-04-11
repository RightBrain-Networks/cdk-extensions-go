package k8saws


// Configuration options for referencing an SSM parameter as a Kubernetes secret.
type SsmParameterReferenceOptions struct {
	// Defines a mapping of how JSON keys in the SSM parameter should appear in the imported Kubernetes secret.
	Fields *[]*SecretFieldReference `field:"optional" json:"fields" yaml:"fields"`
}

