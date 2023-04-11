package k8saws


// Options for configuring the Nest Fluent Bit filter plugin.
// See: [Nest Plugin Documention](https://docs.fluentbit.io/manual/pipeline/filters/nest)
//
type FluentBitNestFilterOptions struct {
	// The pattern to match for records that this output should apply to.
	Match FluentBitMatch `field:"optional" json:"match" yaml:"match"`
	// The operation the filter will perform.
	Operation NestFilterOperation `field:"required" json:"operation" yaml:"operation"`
	// Prefix affected keys with this string.
	AddPrefix *string `field:"optional" json:"addPrefix" yaml:"addPrefix"`
	// Remove prefix from affected keys if it matches this string.
	RemovePrefix *string `field:"optional" json:"removePrefix" yaml:"removePrefix"`
}

