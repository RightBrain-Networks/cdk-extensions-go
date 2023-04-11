package ec2

import (
	"github.com/vibe-io/cdk-extensions-go/cdkextensions/ram"
)

type SharingOptions struct {
	AllowExternalPrincipals *bool `field:"optional" json:"allowExternalPrincipals" yaml:"allowExternalPrincipals"`
	AutoDiscoverAccounts *bool `field:"optional" json:"autoDiscoverAccounts" yaml:"autoDiscoverAccounts"`
	Principals *[]ram.ISharedPrincipal `field:"optional" json:"principals" yaml:"principals"`
}

