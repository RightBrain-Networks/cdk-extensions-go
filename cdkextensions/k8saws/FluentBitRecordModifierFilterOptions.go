package k8saws


// Options for configuring the Record Modifier Fluent Bit filter plugin.
// See: [Record Modifier Plugin Documention](https://docs.fluentbit.io/manual/pipeline/filters/record-modifier)
//
type FluentBitRecordModifierFilterOptions struct {
	// The pattern to match for records that this output should apply to.
	Match FluentBitMatch `field:"optional" json:"match" yaml:"match"`
	// If a tag is not match, that field is removed.
	Allow *[]*string `field:"optional" json:"allow" yaml:"allow"`
	// Add fields to the output.
	Records *[]*AppendedRecord `field:"optional" json:"records" yaml:"records"`
	// If a tag is match, that field is removed.
	Remove *[]*string `field:"optional" json:"remove" yaml:"remove"`
}

