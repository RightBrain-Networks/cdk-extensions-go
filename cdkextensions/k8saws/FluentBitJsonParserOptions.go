package k8saws


// Options for configuring the JSON Fluent Bit parser plugin.
// See: [JSON Plugin Documention](https://docs.fluentbit.io/manual/pipeline/parsers/json)
//
type FluentBitJsonParserOptions struct {
	// Defines the format of the timestamp on the inbound record.
	// See: [strftime](http://man7.org/linux/man-pages/man3/strftime.3.html)
	//
	TimeFormat *string `field:"optional" json:"timeFormat" yaml:"timeFormat"`
	// The key under which timestamp information for the inbound record is given.
	TimeKey *string `field:"optional" json:"timeKey" yaml:"timeKey"`
}

