package glue

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsglue"
)

type CrawlerTargetCollection struct {
	CatalogTargets *[]*awsglue.CfnCrawler_CatalogTargetProperty `field:"optional" json:"catalogTargets" yaml:"catalogTargets"`
	DynamoDbTargets *[]*awsglue.CfnCrawler_DynamoDBTargetProperty `field:"optional" json:"dynamoDbTargets" yaml:"dynamoDbTargets"`
	JdbcTargets *[]*awsglue.CfnCrawler_JdbcTargetProperty `field:"optional" json:"jdbcTargets" yaml:"jdbcTargets"`
	S3Targets *[]*awsglue.CfnCrawler_S3TargetProperty `field:"optional" json:"s3Targets" yaml:"s3Targets"`
}

