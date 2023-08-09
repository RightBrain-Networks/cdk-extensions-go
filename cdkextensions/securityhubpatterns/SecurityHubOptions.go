package securityhubpatterns

import (
	"github.com/vibe-io/cdk-extensions-go/cdkextensions/securityhub"
)

type SecurityHubOptions struct {
	Enabled *bool `field:"optional" json:"enabled" yaml:"enabled"`
	Hub securityhub.IHub `field:"optional" json:"hub" yaml:"hub"`
}

