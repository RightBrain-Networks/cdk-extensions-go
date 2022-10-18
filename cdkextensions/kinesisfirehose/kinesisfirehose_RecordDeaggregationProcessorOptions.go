package kinesisfirehose


type RecordDeaggregationProcessorOptions struct {
	SubRecordType SubRecordType `field:"required" json:"subRecordType" yaml:"subRecordType"`
	Delimiter *string `field:"optional" json:"delimiter" yaml:"delimiter"`
}

