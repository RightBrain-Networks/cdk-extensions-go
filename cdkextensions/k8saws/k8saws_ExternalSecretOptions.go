package k8saws


type ExternalSecretOptions struct {
	Fields *[]*SecretFieldReference `field:"optional" json:"fields" yaml:"fields"`
	Name *string `field:"optional" json:"name" yaml:"name"`
}

