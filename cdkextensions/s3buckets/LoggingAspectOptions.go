package s3buckets

import (
	"github.com/aws/constructs-go/constructs/v10"
)

type LoggingAspectOptions struct {
	Exclusions *[]constructs.IConstruct `field:"optional" json:"exclusions" yaml:"exclusions"`
	Force *bool `field:"optional" json:"force" yaml:"force"`
	Prefix *string `field:"optional" json:"prefix" yaml:"prefix"`
}

