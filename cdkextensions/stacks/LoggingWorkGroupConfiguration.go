package stacks

import (
	"github.com/vibe-io/cdk-extensions-go/cdkextensions/core"
)

type LoggingWorkGroupConfiguration struct {
	Enabled *bool `field:"optional" json:"enabled" yaml:"enabled"`
	QueryScannedBytesLimit core.DataSize `field:"optional" json:"queryScannedBytesLimit" yaml:"queryScannedBytesLimit"`
}

