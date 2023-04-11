package ec2patterns

import (
	"github.com/vibe-io/cdk-extensions-go/cdkextensions/ram"
)

type FourTierNetworkShareProperties struct {
	AllowExternalPrincipals *bool `field:"optional" json:"allowExternalPrincipals" yaml:"allowExternalPrincipals"`
	AutoAddAccounts *bool `field:"optional" json:"autoAddAccounts" yaml:"autoAddAccounts"`
	Pricipals *[]ram.ISharedPrincipal `field:"optional" json:"pricipals" yaml:"pricipals"`
}

