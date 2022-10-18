package kinesisfirehose


type OrcOutputSerDeOptions struct {
	BlockSizeBytes *float64 `field:"optional" json:"blockSizeBytes" yaml:"blockSizeBytes"`
	BloomFilterColumns *[]*string `field:"optional" json:"bloomFilterColumns" yaml:"bloomFilterColumns"`
	BloomFilterFalsePositiveProbability *float64 `field:"optional" json:"bloomFilterFalsePositiveProbability" yaml:"bloomFilterFalsePositiveProbability"`
	Compression OrcCompressionFormat `field:"optional" json:"compression" yaml:"compression"`
	DictionaryKeyThreshold *float64 `field:"optional" json:"dictionaryKeyThreshold" yaml:"dictionaryKeyThreshold"`
	EnablePadding *bool `field:"optional" json:"enablePadding" yaml:"enablePadding"`
	FormatVersion OrcFormatVersion `field:"optional" json:"formatVersion" yaml:"formatVersion"`
	PaddingTolerance *float64 `field:"optional" json:"paddingTolerance" yaml:"paddingTolerance"`
	RowIndexStride *float64 `field:"optional" json:"rowIndexStride" yaml:"rowIndexStride"`
	StripeSizeBytes *float64 `field:"optional" json:"stripeSizeBytes" yaml:"stripeSizeBytes"`
}

