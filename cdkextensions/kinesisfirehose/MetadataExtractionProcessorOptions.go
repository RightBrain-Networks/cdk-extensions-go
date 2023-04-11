package kinesisfirehose


type MetadataExtractionProcessorOptions struct {
	Query MetaDataExtractionQuery `field:"required" json:"query" yaml:"query"`
	Engine JsonParsingEngine `field:"optional" json:"engine" yaml:"engine"`
}

