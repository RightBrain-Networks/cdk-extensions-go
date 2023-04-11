package aps


// The details that are needed to configure an APS rule groups namespace that will consume a rule group configuration.
type RuleGroupConfigurationDetails struct {
	// The content of the rules configuration definition, in the format of an APS rules configuration file.
	// See: [Creating a rules file](https://docs.aws.amazon.com/prometheus/latest/userguide/AMP-ruler-rulesfile.html)
	//
	Content *string `field:"required" json:"content" yaml:"content"`
}

