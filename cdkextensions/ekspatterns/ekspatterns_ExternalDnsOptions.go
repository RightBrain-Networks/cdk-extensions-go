package ekspatterns


type ExternalDnsOptions struct {
	DomainFilter *[]*string `field:"optional" json:"domainFilter" yaml:"domainFilter"`
	Enabled *bool `field:"optional" json:"enabled" yaml:"enabled"`
}

