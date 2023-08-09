package securityhub


type RuleSetProps struct {
	Id *string `field:"required" json:"id" yaml:"id"`
	Version *string `field:"required" json:"version" yaml:"version"`
	Default *bool `field:"optional" json:"default" yaml:"default"`
	Description *string `field:"optional" json:"description" yaml:"description"`
	Global *bool `field:"optional" json:"global" yaml:"global"`
	Name *string `field:"optional" json:"name" yaml:"name"`
}

