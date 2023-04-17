package athena


type WorkGroupOptions struct {
	Description *string `field:"optional" json:"description" yaml:"description"`
	Engine IAnalyticsEngine `field:"optional" json:"engine" yaml:"engine"`
	Name *string `field:"optional" json:"name" yaml:"name"`
	RecursiveDelete *bool `field:"optional" json:"recursiveDelete" yaml:"recursiveDelete"`
	State WorkGroupState `field:"optional" json:"state" yaml:"state"`
}

