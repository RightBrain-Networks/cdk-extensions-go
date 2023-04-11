package kinesisfirehose


type ParquetOutputSerDeOptions struct {
	BlockSizeBytes *float64 `field:"optional" json:"blockSizeBytes" yaml:"blockSizeBytes"`
	Compression ParquetCompressionFormat `field:"optional" json:"compression" yaml:"compression"`
	EnableDictionaryCompression *bool `field:"optional" json:"enableDictionaryCompression" yaml:"enableDictionaryCompression"`
	MaxPaddingBytes *float64 `field:"optional" json:"maxPaddingBytes" yaml:"maxPaddingBytes"`
	PageSizeBytes *float64 `field:"optional" json:"pageSizeBytes" yaml:"pageSizeBytes"`
	WriterVersion ParquetWriterVersion `field:"optional" json:"writerVersion" yaml:"writerVersion"`
}

