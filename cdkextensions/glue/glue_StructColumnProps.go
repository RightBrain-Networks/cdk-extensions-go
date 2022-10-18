package glue


type StructColumnProps struct {
	Comment *string `field:"optional" json:"comment" yaml:"comment"`
	Name *string `field:"optional" json:"name" yaml:"name"`
	Data *[]Column `field:"optional" json:"data" yaml:"data"`
}

