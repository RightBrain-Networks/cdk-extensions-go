package k8saws


// Options for configuring the Parser Fluent Bit filter plugin.
// See: [Parser Plugin Documention](https://docs.fluentbit.io/manual/pipeline/filters/parser)
//
type FluentBitParserFilterOptions struct {
	// The pattern to match for records that this output should apply to.
	Match FluentBitMatch `field:"optional" json:"match" yaml:"match"`
	// Specify field name in record to parse.
	KeyName *string `field:"required" json:"keyName" yaml:"keyName"`
	// The parsers to use to interpret the field.
	Parsers *[]IFluentBitParserPlugin `field:"optional" json:"parsers" yaml:"parsers"`
	// Keep original `keyName` field in the parsed result.
	//
	// If `false`, the field will be removed.
	// Default: false.
	//
	PreserveKey *bool `field:"optional" json:"preserveKey" yaml:"preserveKey"`
	// Keep all other original fields in the parsed result.
	//
	// If `false`, all other original fields will be removed.
	// Default: false.
	//
	ReserveData *bool `field:"optional" json:"reserveData" yaml:"reserveData"`
}

