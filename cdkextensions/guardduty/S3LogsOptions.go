package guardduty


// Options for configuring whether S3 data events should be used as a data source for GuardDuty.
type S3LogsOptions struct {
	// Controls whether S3 data events are enabled as a data source for GuardDuty.
	Enabled *bool `field:"optional" json:"enabled" yaml:"enabled"`
}

