package glue


// Configuration for Crawler S3 target.
type S3TargetOptions struct {
	Connection Connection `field:"optional" json:"connection" yaml:"connection"`
	Exclusions *[]*string `field:"optional" json:"exclusions" yaml:"exclusions"`
	KeyPrefix *string `field:"optional" json:"keyPrefix" yaml:"keyPrefix"`
	SampleSize *string `field:"optional" json:"sampleSize" yaml:"sampleSize"`
}

