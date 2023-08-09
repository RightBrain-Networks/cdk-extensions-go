package config

import (
	"github.com/vibe-io/cdk-extensions-go/cdkextensions/ssm"
)

type AutomationDocumentRemediationProps struct {
	Document ssm.IAutomationDocument `field:"required" json:"document" yaml:"document"`
	ConcurrencyPercentage *float64 `field:"optional" json:"concurrencyPercentage" yaml:"concurrencyPercentage"`
	ErrorPercentage *float64 `field:"optional" json:"errorPercentage" yaml:"errorPercentage"`
	Version *string `field:"optional" json:"version" yaml:"version"`
}

