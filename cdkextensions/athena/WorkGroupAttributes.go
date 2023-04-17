package athena


type WorkGroupAttributes struct {
	Arn *string `field:"optional" json:"arn" yaml:"arn"`
	CreationTime *string `field:"optional" json:"creationTime" yaml:"creationTime"`
	EffectiveEngineVersion *string `field:"optional" json:"effectiveEngineVersion" yaml:"effectiveEngineVersion"`
	Name *string `field:"optional" json:"name" yaml:"name"`
}

