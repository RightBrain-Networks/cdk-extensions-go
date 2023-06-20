package stepfunctions


type StringReplaceProps struct {
	InputString *string `field:"required" json:"inputString" yaml:"inputString"`
	OutputKey *string `field:"required" json:"outputKey" yaml:"outputKey"`
	Replace *string `field:"required" json:"replace" yaml:"replace"`
	Search *string `field:"required" json:"search" yaml:"search"`
}

