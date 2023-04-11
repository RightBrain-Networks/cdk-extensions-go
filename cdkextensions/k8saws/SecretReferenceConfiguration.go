package k8saws


// Configuration detailing how secrets are to be synchronized.
type SecretReferenceConfiguration struct {
	// The ID of the secret to be imported from the provider.
	RemoteRef *string `field:"required" json:"remoteRef" yaml:"remoteRef"`
	// A mapping of fields and per field options to use when synchronizing a secret from a provider.
	Fields *[]*SecretFieldReference `field:"optional" json:"fields" yaml:"fields"`
}

