package alerting


type DescriptionBuilderSectionProps struct {
	Title *string `field:"required" json:"title" yaml:"title"`
	ReferenceChecks *[]*string `field:"optional" json:"referenceChecks" yaml:"referenceChecks"`
}

