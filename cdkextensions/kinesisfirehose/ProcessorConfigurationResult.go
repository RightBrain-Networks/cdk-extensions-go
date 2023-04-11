package kinesisfirehose


type ProcessorConfigurationResult struct {
	Processors *[]DeliveryStreamProcessor `field:"required" json:"processors" yaml:"processors"`
	Enabled *bool `field:"optional" json:"enabled" yaml:"enabled"`
}

