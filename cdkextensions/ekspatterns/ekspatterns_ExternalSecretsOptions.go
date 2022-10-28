package ekspatterns


type ExternalSecretsOptions struct {
	CreateNamespace *bool `field:"optional" json:"createNamespace" yaml:"createNamespace"`
	Enabled *bool `field:"optional" json:"enabled" yaml:"enabled"`
	Name *string `field:"optional" json:"name" yaml:"name"`
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
}

