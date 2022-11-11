package k8saws


// Options for configuring the Regex Fluent Bit parser plugin.
// See: [Regex Plugin Documention](https://docs.fluentbit.io/manual/pipeline/parsers/regular-expression)
//
type FluentBitRegexParserOptions struct {
	// The regular expression to use to parse the incoming records.
	//
	// Use regex group names to define the name of fields being captured.
	Regex *string `field:"required" json:"regex" yaml:"regex"`
	// If enabled, the parser ignores empty value of the record.
	SkipEmptyValues *bool `field:"optional" json:"skipEmptyValues" yaml:"skipEmptyValues"`
	// Defines the format of the timestamp on the inbound record.
	// See: [strftime](http://man7.org/linux/man-pages/man3/strftime.3.html)
	//
	TimeFormat *string `field:"optional" json:"timeFormat" yaml:"timeFormat"`
	// The key under which timestamp information for the inbound record is given.
	TimeKey *string `field:"optional" json:"timeKey" yaml:"timeKey"`
	// Maps group names matched by the regex to the data types they should be interpreted as.
	Types *map[string]ParserPluginDataType `field:"optional" json:"types" yaml:"types"`
}

