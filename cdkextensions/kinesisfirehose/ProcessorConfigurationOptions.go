package kinesisfirehose


type ProcessorConfigurationOptions struct {
	Enabled *bool `field:"optional" json:"enabled" yaml:"enabled"`
	Processors *[]DeliveryStreamProcessor `field:"optional" json:"processors" yaml:"processors"`
}

