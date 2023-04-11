package k8saws


// Options for configuring the Modify Fluent Bit filter plugin.
// See: [Modify Plugin Documention](https://docs.fluentbit.io/manual/pipeline/filters/modify)
//
type FluentBitModifyFilterOptions struct {
	// The pattern to match for records that this output should apply to.
	Match FluentBitMatch `field:"optional" json:"match" yaml:"match"`
	Conditions *[]ModifyCondition `field:"optional" json:"conditions" yaml:"conditions"`
	Operations *[]ModifyOperation `field:"optional" json:"operations" yaml:"operations"`
}

