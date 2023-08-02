package k8saws


// Configures a pattern to match against a Fluent Bit record.
type FluentBitGrepRegex struct {
	// The key of the fields which you want to filter using the regex.
	Key *string `field:"required" json:"key" yaml:"key"`
	// The regular expression to apply to the specified field.
	Regex *string `field:"required" json:"regex" yaml:"regex"`
	// Whether the matched expression should exclude or include records from being output.
	//
	// When this is true, only records that match the given expression will be
	// output.
	//
	// When this is false, only records that do not match the given expression
	// will be output.
	// Default: false.
	//
	Exclude *bool `field:"optional" json:"exclude" yaml:"exclude"`
}

