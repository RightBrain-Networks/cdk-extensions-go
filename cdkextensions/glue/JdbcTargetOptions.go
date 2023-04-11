package glue


// Configuration for Crawler JDBC target.
type JdbcTargetOptions struct {
	// A list of glob patterns used to exclude from the crawl.
	//
	// For more information.
	// See: [Catalog Tables with a Crawler](https://docs.aws.amazon.com/glue/latest/dg/add-crawler.html)
	//
	Exclusions *[]*string `field:"optional" json:"exclusions" yaml:"exclusions"`
	// The path of the JDBC target.
	// See: [AWS::Glue::Crawler JdbcTarget](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-crawler-jdbctarget.html#cfn-glue-crawler-jdbctarget-path)
	//
	Paths *[]*string `field:"optional" json:"paths" yaml:"paths"`
}

