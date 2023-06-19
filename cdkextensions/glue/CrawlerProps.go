package glue

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
)

// Configuration for Crawler.
type CrawlerProps struct {
	// The AWS account ID this resource belongs to.
	Account *string `field:"optional" json:"account" yaml:"account"`
	// ARN to deduce region and account from.
	//
	// The ARN is parsed and the account and region are taken from the ARN.
	// This should be used for imported resources.
	//
	// Cannot be supplied together with either `account` or `region`.
	EnvironmentFromArn *string `field:"optional" json:"environmentFromArn" yaml:"environmentFromArn"`
	// The value passed in by users to the physical name prop of the resource.
	//
	// - `undefined` implies that a physical name will be allocated by
	//   CloudFormation during deployment.
	// - a concrete value implies a specific physical name
	// - `PhysicalName.GENERATE_IF_NEEDED` is a marker that indicates that a physical will only be generated
	//   by the CDK if it is needed for cross-environment references. Otherwise, it will be allocated by CloudFormation.
	PhysicalName *string `field:"optional" json:"physicalName" yaml:"physicalName"`
	// The AWS region this resource belongs to.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// A list of UTF-8 strings that specify the names of custom classifiers that are associated with the crawler.
	// See: [AWS::Glue::Crawler](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-crawler.html#cfn-glue-crawler-classifiers)
	//
	Classifiers *[]*string `field:"optional" json:"classifiers" yaml:"classifiers"`
	// Crawler configuration information.
	//
	// This versioned JSON string allows users to specify aspects of a crawler's behavior. For more information, see Configuring a Crawler.
	// See: [AWS::Glue::Crawler](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-crawler.html#cfn-glue-crawler-configuration)
	//
	Configuration *CrawlerConfiguration `field:"optional" json:"configuration" yaml:"configuration"`
	// The {@link aws-glue.Database Database } object in which the crawler's output is stored.
	Database Database `field:"optional" json:"database" yaml:"database"`
	// The deletion behavior when the crawler finds a deleted object.
	// See: [AWS::Glue::Crawler SchemaChangePolicy](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-crawler-schemachangepolicy.html#cfn-glue-crawler-schemachangepolicy-deletebehavior)
	//
	DeleteBehavior DeleteBehavior `field:"optional" json:"deleteBehavior" yaml:"deleteBehavior"`
	// Description of the Crawler.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Name of the Crawler.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// When crawling an Amazon S3 data source after the first crawl is complete, specifies whether to crawl the entire dataset again or to crawl only folders that were added since the last crawler run.
	// See: [AWS::Glue::Crawler RecrawlPolicy](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-crawler-recrawlpolicy.html)
	//
	RecrawlBehavior RecrawlBehavior `field:"optional" json:"recrawlBehavior" yaml:"recrawlBehavior"`
	// For scheduled crawlers, the schedule when the crawler runs.
	// See: [AWS::Glue::Crawler Schedule](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-crawler-schedule.html)
	//
	ScheduleExpression awsevents.Schedule `field:"optional" json:"scheduleExpression" yaml:"scheduleExpression"`
	// A {@link aws-glue.SecurityConfiguration SecurityConfiguration } object to apply to the Crawler.
	SecurityConfiguration SecurityConfiguration `field:"optional" json:"securityConfiguration" yaml:"securityConfiguration"`
	// The prefix added to the names of tables that are created.
	// See: [AWS::Glue::Crawler](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-crawler.html#cfn-glue-crawler-tableprefix)
	//
	TablePrefix *string `field:"optional" json:"tablePrefix" yaml:"tablePrefix"`
	// A collection of targets to crawl.
	// See: [AWS::Glue::Crawler](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-crawler.html#cfn-glue-crawler-targets)
	//
	Targets *[]ICrawlerTarget `field:"optional" json:"targets" yaml:"targets"`
	// The update behavior when the crawler finds a changed schema.
	// See: [AWS::Glue::Crawler SchemaChangePolicy](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-crawler-schemachangepolicy.html#cfn-glue-crawler-schemachangepolicy-updatebehavior)
	//
	UpdateBehavior UpdateBehavior `field:"optional" json:"updateBehavior" yaml:"updateBehavior"`
}

