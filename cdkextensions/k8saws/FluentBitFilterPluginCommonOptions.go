package k8saws


// Configuration options that apply to all Fluent Bit output plugins.
type FluentBitFilterPluginCommonOptions struct {
	// The pattern to match for records that this output should apply to.
	Match FluentBitMatch `field:"optional" json:"match" yaml:"match"`
}

