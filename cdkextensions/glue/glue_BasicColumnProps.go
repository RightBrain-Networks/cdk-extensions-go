package glue


type BasicColumnProps struct {
	Comment *string `field:"optional" json:"comment" yaml:"comment"`
	Name *string `field:"optional" json:"name" yaml:"name"`
	Type *string `field:"required" json:"type" yaml:"type"`
}

