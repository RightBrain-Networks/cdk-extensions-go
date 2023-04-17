package athena

import (
	"github.com/vibe-io/cdk-extensions-go/cdkextensions/core"
)

type AthenaSqlEngineOptions struct {
	EnforceConfiguration *bool `field:"optional" json:"enforceConfiguration" yaml:"enforceConfiguration"`
	EngineVersion AthenaSqlEngineVersion `field:"optional" json:"engineVersion" yaml:"engineVersion"`
	Output *AnalyticsEngineOutputOptions `field:"optional" json:"output" yaml:"output"`
	PublishMetrics *bool `field:"optional" json:"publishMetrics" yaml:"publishMetrics"`
	QueryScannedBytesLimit core.DataSize `field:"optional" json:"queryScannedBytesLimit" yaml:"queryScannedBytesLimit"`
	RequesterPays *bool `field:"optional" json:"requesterPays" yaml:"requesterPays"`
}

