package glue


// Properties of a DataFormat instance.
type DataFormatProps struct {
	// `InputFormat` for this data format.
	InputFormat InputFormat `field:"required" json:"inputFormat" yaml:"inputFormat"`
	// `OutputFormat` for this data format.
	OutputFormat OutputFormat `field:"required" json:"outputFormat" yaml:"outputFormat"`
	// Serialization library for this data format.
	SerializationLibrary SerializationLibrary `field:"required" json:"serializationLibrary" yaml:"serializationLibrary"`
	// Classification string given to tables with this data format.
	ClassificationString ClassificationString `field:"optional" json:"classificationString" yaml:"classificationString"`
}

