package k8saws

import (
	"github.com/vibe-io/cdk-extensions-go/cdkextensions/core"
)

// Options for configuring the Parser Fluent Bit filter plugin.
// See: [Parser Plugin Documention](https://docs.fluentbit.io/manual/pipeline/filters/parser)
//
type FluentBitRewriteTagFilterOptions struct {
	// The pattern to match for records that this output should apply to.
	Match FluentBitMatch `field:"optional" json:"match" yaml:"match"`
	// Set a limit on the amount of memory the tag rewrite emitter can consume if the outputs provide backpressure.
	// Default: 10M.
	//
	EmitterMemBufLimit core.DataSize `field:"optional" json:"emitterMemBufLimit" yaml:"emitterMemBufLimit"`
	// When the filter emits a record under the new Tag, there is an internal emitter plugin that takes care of the job.
	//
	// Since this emitter expose
	// metrics as any other component of the pipeline, you can use this
	// property to configure an optional name for it.
	EmitterName *string `field:"optional" json:"emitterName" yaml:"emitterName"`
	// Define a buffering mechanism for the new records created.
	//
	// Note these records are part of the emitter plugin.
	EmitterStorageType EmitterStorageType `field:"optional" json:"emitterStorageType" yaml:"emitterStorageType"`
	// Defines the matching criteria and the format of the Tag for the matching record.
	Rules *[]*RewriteTagRule `field:"optional" json:"rules" yaml:"rules"`
}

