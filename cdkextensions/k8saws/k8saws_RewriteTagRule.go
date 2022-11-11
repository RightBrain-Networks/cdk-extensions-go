package k8saws


// Defines the matching criteria and the format of the Tag for the rewrite tag Fluent Bit filter plugin.
// See: [Rules](https://docs.fluentbit.io/manual/pipeline/filters/rewrite-tag#rules)
//
type RewriteTagRule struct {
	// If a rule matches a rule the filter will emit a copy of the record with the new defined Tag.
	//
	// The property keep takes a boolean value to define if the original
	// record with the old Tag must be preserved and continue in the pipeline
	// or just be discarded.
	// See: [Keep](https://docs.fluentbit.io/manual/pipeline/filters/rewrite-tag#keep)
	//
	Keep *bool `field:"required" json:"keep" yaml:"keep"`
	// The key represents the name of the record key that holds the value that we want to use to match our regular expression.
	//
	// A key name is specified and prefixed with a `$`.
	// See: [Key](https://docs.fluentbit.io/manual/pipeline/filters/rewrite-tag#key)
	//
	Key *string `field:"required" json:"key" yaml:"key"`
	// If a regular expression has matched the value of the defined key in the rule, we are ready to compose a new Tag for that specific record.
	//
	// The tag is a concatenated string that can contain any of the following
	// characters: `a-z,A-Z,0-9` and `.-,`.
	// See: [New Tag](https://docs.fluentbit.io/manual/pipeline/filters/rewrite-tag#new-tag)
	//
	NewTag *string `field:"required" json:"newTag" yaml:"newTag"`
	// Using a simple regular expression we can specify a matching pattern to use against the value of the key specified, also we can take advantage of group capturing to create custom placeholder values.
	// See: [Rubular regex tester](https://rubular.com/)
	//
	Regex *string `field:"required" json:"regex" yaml:"regex"`
}

