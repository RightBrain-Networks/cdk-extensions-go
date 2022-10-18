package glue

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
)

// Configuration for Crawlner.
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
	//    CloudFormation during deployment.
	// - a concrete value implies a specific physical name
	// - `PhysicalName.GENERATE_IF_NEEDED` is a marker that indicates that a physical will only be generated
	//    by the CDK if it is needed for cross-environment references. Otherwise, it will be allocated by CloudFormation.
	PhysicalName *string `field:"optional" json:"physicalName" yaml:"physicalName"`
	// The AWS region this resource belongs to.
	Region *string `field:"optional" json:"region" yaml:"region"`
	Classifiers *[]*string `field:"optional" json:"classifiers" yaml:"classifiers"`
	Configuration *CrawlerConfiguration `field:"optional" json:"configuration" yaml:"configuration"`
	Database Database `field:"optional" json:"database" yaml:"database"`
	DeleteBehavior DeleteBehavior `field:"optional" json:"deleteBehavior" yaml:"deleteBehavior"`
	Description *string `field:"optional" json:"description" yaml:"description"`
	Name *string `field:"optional" json:"name" yaml:"name"`
	RecrawlBehavior RecrawlBehavior `field:"optional" json:"recrawlBehavior" yaml:"recrawlBehavior"`
	ScheduleExpression awsevents.Schedule `field:"optional" json:"scheduleExpression" yaml:"scheduleExpression"`
	SecurityConfiguration SecurityConfiguration `field:"optional" json:"securityConfiguration" yaml:"securityConfiguration"`
	TablePrefix *string `field:"optional" json:"tablePrefix" yaml:"tablePrefix"`
	Targets *[]ICrawlerTarget `field:"optional" json:"targets" yaml:"targets"`
	UpdateBehavior UpdateBehavior `field:"optional" json:"updateBehavior" yaml:"updateBehavior"`
}

