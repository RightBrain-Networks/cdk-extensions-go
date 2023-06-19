package alerting


type AddReferenceProps struct {
	Value *string `field:"required" json:"value" yaml:"value"`
	Delimiter AppendDelimiter `field:"optional" json:"delimiter" yaml:"delimiter"`
	Label *string `field:"optional" json:"label" yaml:"label"`
	Required *bool `field:"optional" json:"required" yaml:"required"`
}

