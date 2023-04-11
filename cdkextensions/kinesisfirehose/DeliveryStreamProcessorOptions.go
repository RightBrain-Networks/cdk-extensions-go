package kinesisfirehose


type DeliveryStreamProcessorOptions struct {
	ProcessorType ProcessorType `field:"required" json:"processorType" yaml:"processorType"`
	Parameters *map[string]*string `field:"optional" json:"parameters" yaml:"parameters"`
}

