package k8saws


type NamespacedExternalSecretOptions struct {
	Fields *[]*SecretFieldReference `field:"optional" json:"fields" yaml:"fields"`
	Name *string `field:"optional" json:"name" yaml:"name"`
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
}

