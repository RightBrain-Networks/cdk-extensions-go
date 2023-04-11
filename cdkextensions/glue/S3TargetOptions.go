package glue


// Configuration for Crawler S3 target.
type S3TargetOptions struct {
	// A {@link aws-glue.Connection | "Connection" } object to connect to the target with.
	Connection Connection `field:"optional" json:"connection" yaml:"connection"`
	// A list of glob patterns used to exclude from the crawl.
	// See: [For More Information](https://docs.aws.amazon.com/glue/latest/dg/add-crawler.html)
	//
	Exclusions *[]*string `field:"optional" json:"exclusions" yaml:"exclusions"`
	// A Prefix Key for identification and organization of objects in the bucket.
	KeyPrefix *string `field:"optional" json:"keyPrefix" yaml:"keyPrefix"`
	// Sets the number of files in each leaf folder to be crawled when crawling sample files in a dataset.
	//
	// If not set, all the files are crawled. A valid value is an integer between 1 and 249.
	// See: [AWS::Glue::Crawler S3Target](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-crawler-s3target.html#cfn-glue-crawler-s3target-samplesize)
	//
	SampleSize *string `field:"optional" json:"sampleSize" yaml:"sampleSize"`
}

