package securityhub


type ScopedRuleSet struct {
	Arn *string `field:"required" json:"arn" yaml:"arn"`
	Id *string `field:"required" json:"id" yaml:"id"`
	Version *string `field:"required" json:"version" yaml:"version"`
	Default *bool `field:"optional" json:"default" yaml:"default"`
	Description *string `field:"optional" json:"description" yaml:"description"`
	Name *string `field:"optional" json:"name" yaml:"name"`
}

