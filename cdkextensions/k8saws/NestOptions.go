package k8saws


type NestOptions struct {
	// Nest records matching `wildcard` under this key.
	NestUnder *string `field:"required" json:"nestUnder" yaml:"nestUnder"`
	// Nest records which field matches this wildcard,.
	Wildcards *[]*string `field:"required" json:"wildcards" yaml:"wildcards"`
}

