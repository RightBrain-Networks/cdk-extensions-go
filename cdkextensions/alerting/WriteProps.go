package alerting


type WriteProps struct {
	Value *string `field:"required" json:"value" yaml:"value"`
	DefaultDelimiter AppendDelimiter `field:"optional" json:"defaultDelimiter" yaml:"defaultDelimiter"`
	Delimiter AppendDelimiter `field:"optional" json:"delimiter" yaml:"delimiter"`
	Prefix *string `field:"optional" json:"prefix" yaml:"prefix"`
	Required *bool `field:"optional" json:"required" yaml:"required"`
	Suffix *string `field:"optional" json:"suffix" yaml:"suffix"`
}

