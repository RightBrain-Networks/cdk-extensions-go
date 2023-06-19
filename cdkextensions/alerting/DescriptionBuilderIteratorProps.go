package alerting


type DescriptionBuilderIteratorProps struct {
	ArrayRef *string `field:"required" json:"arrayRef" yaml:"arrayRef"`
	ResultPath *string `field:"required" json:"resultPath" yaml:"resultPath"`
	FieldDelimiter AppendDelimiter `field:"optional" json:"fieldDelimiter" yaml:"fieldDelimiter"`
	RecordDelimiter AppendDelimiter `field:"optional" json:"recordDelimiter" yaml:"recordDelimiter"`
	SectionDelimiter AppendDelimiter `field:"optional" json:"sectionDelimiter" yaml:"sectionDelimiter"`
	Title *string `field:"optional" json:"title" yaml:"title"`
}

