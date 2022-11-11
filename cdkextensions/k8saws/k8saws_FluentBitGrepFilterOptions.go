package k8saws


// Options for configuring the Grep Fluent Bit filter plugin.
// See: [Grep Plugin Documention](https://docs.fluentbit.io/manual/pipeline/filters/grep)
//
type FluentBitGrepFilterOptions struct {
	// The pattern to match for records that this output should apply to.
	Match FluentBitMatch `field:"optional" json:"match" yaml:"match"`
	// The pattern to use for filtering records processed by the plugin.
	Pattern *FluentBitGrepRegex `field:"required" json:"pattern" yaml:"pattern"`
}

