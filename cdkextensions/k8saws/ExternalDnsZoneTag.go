package k8saws


// Specifies a tag that can be used to restrict which Hosted Zone external-dns will have access to.
type ExternalDnsZoneTag struct {
	// The name of the tag to filter on.
	Key *string `field:"required" json:"key" yaml:"key"`
	// The value of the tag to filter on.
	Value *string `field:"required" json:"value" yaml:"value"`
}

