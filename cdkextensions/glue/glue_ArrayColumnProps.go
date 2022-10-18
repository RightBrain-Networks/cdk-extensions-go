package glue


type ArrayColumnProps struct {
	Comment *string `field:"optional" json:"comment" yaml:"comment"`
	Name *string `field:"optional" json:"name" yaml:"name"`
	Data Column `field:"required" json:"data" yaml:"data"`
}

