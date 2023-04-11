package kinesisfirehose


type BackupConfigurationOptions struct {
	Destination IDeliveryStreamBackupDestination `field:"required" json:"destination" yaml:"destination"`
	Enabled *bool `field:"optional" json:"enabled" yaml:"enabled"`
}

