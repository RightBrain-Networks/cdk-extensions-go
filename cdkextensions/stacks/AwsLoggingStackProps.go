package stacks

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/vibe-io/cdk-extensions-go/cdkextensions/ec2"
	"github.com/vibe-io/cdk-extensions-go/cdkextensions/s3buckets"
)

// Configuration for AwsLoggingStack.
type AwsLoggingStackProps struct {
	// Include runtime versioning information in this Stack.
	AnalyticsReporting *bool `field:"optional" json:"analyticsReporting" yaml:"analyticsReporting"`
	// Enable this flag to allow native cross region stack references.
	//
	// Enabling this will create a CloudFormation custom resource
	// in both the producing stack and consuming stack in order to perform the export/import
	//
	// This feature is currently experimental.
	CrossRegionReferences *bool `field:"optional" json:"crossRegionReferences" yaml:"crossRegionReferences"`
	// A description of the stack.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The AWS environment (account/region) where this stack will be deployed.
	//
	// Set the `region`/`account` fields of `env` to either a concrete value to
	// select the indicated environment (recommended for production stacks), or to
	// the values of environment variables
	// `CDK_DEFAULT_REGION`/`CDK_DEFAULT_ACCOUNT` to let the target environment
	// depend on the AWS credentials/configuration that the CDK CLI is executed
	// under (recommended for development stacks).
	//
	// If the `Stack` is instantiated inside a `Stage`, any undefined
	// `region`/`account` fields from `env` will default to the same field on the
	// encompassing `Stage`, if configured there.
	//
	// If either `region` or `account` are not set nor inherited from `Stage`, the
	// Stack will be considered "*environment-agnostic*"". Environment-agnostic
	// stacks can be deployed to any environment but may not be able to take
	// advantage of all features of the CDK. For example, they will not be able to
	// use environmental context lookups such as `ec2.Vpc.fromLookup` and will not
	// automatically translate Service Principals to the right format based on the
	// environment's AWS partition, and other such enhancements.
	//
	// Example:
	//   // Use a concrete account and region to deploy this stack to:
	//   // `.account` and `.region` will simply return these values.
	//   new Stack(app, 'Stack1', {
	//     env: {
	//       account: '123456789012',
	//       region: 'us-east-1'
	//     },
	//   });
	//
	//   // Use the CLI's current credentials to determine the target environment:
	//   // `.account` and `.region` will reflect the account+region the CLI
	//   // is configured to use (based on the user CLI credentials)
	//   new Stack(app, 'Stack2', {
	//     env: {
	//       account: process.env.CDK_DEFAULT_ACCOUNT,
	//       region: process.env.CDK_DEFAULT_REGION
	//     },
	//   });
	//
	//   // Define multiple stacks stage associated with an environment
	//   const myStage = new Stage(app, 'MyStage', {
	//     env: {
	//       account: '123456789012',
	//       region: 'us-east-1'
	//     }
	//   });
	//
	//   // both of these stacks will use the stage's account/region:
	//   // `.account` and `.region` will resolve to the concrete values as above
	//   new MyStack(myStage, 'Stack1');
	//   new YourStack(myStage, 'Stack2');
	//
	//   // Define an environment-agnostic stack:
	//   // `.account` and `.region` will resolve to `{ "Ref": "AWS::AccountId" }` and `{ "Ref": "AWS::Region" }` respectively.
	//   // which will only resolve to actual values by CloudFormation during deployment.
	//   new MyStack(app, 'Stack1');
	//
	Env *awscdk.Environment `field:"optional" json:"env" yaml:"env"`
	// Options for applying a permissions boundary to all IAM Roles and Users created within this Stage.
	PermissionsBoundary awscdk.PermissionsBoundary `field:"optional" json:"permissionsBoundary" yaml:"permissionsBoundary"`
	// Name to deploy the stack with.
	StackName *string `field:"optional" json:"stackName" yaml:"stackName"`
	// Synthesis method to use while deploying this stack.
	//
	// The Stack Synthesizer controls aspects of synthesis and deployment,
	// like how assets are referenced and what IAM roles to use. For more
	// information, see the README of the main CDK package.
	//
	// If not specified, the `defaultStackSynthesizer` from `App` will be used.
	// If that is not specified, `DefaultStackSynthesizer` is used if
	// `@aws-cdk/core:newStyleStackSynthesis` is set to `true` or the CDK major
	// version is v2. In CDK v1 `LegacyStackSynthesizer` is the default if no
	// other synthesizer is specified.
	Synthesizer awscdk.IStackSynthesizer `field:"optional" json:"synthesizer" yaml:"synthesizer"`
	// Stack tags that will be applied to all the taggable resources and the stack itself.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Whether to enable termination protection for this stack.
	TerminationProtection *bool `field:"optional" json:"terminationProtection" yaml:"terminationProtection"`
	// A cdk-extensions/s3-buckets {@link aws-s3-buckets !AlbLogsBucket} object.
	AlbLogsBucket s3buckets.AlbLogsBucket `field:"optional" json:"albLogsBucket" yaml:"albLogsBucket"`
	// A cdk-extensions/s3-buckets {@link aws-s3-buckets !CloudfrontLogsBucket} object.
	CloudfrontLogsBucket s3buckets.CloudfrontLogsBucket `field:"optional" json:"cloudfrontLogsBucket" yaml:"cloudfrontLogsBucket"`
	// A cdk-extensions/s3-buckets {@link aws-s3-buckets !CloudtrailBucket} object.
	CloudtrailLogsBucket s3buckets.CloudtrailBucket `field:"optional" json:"cloudtrailLogsBucket" yaml:"cloudtrailLogsBucket"`
	// Name used for the Glue Database that will be created.
	DatabaseName *string `field:"optional" json:"databaseName" yaml:"databaseName"`
	// A cdk-extensions/s3-buckets {@link aws-s3-buckets !FlowLogsBucket} object.
	FlowLogsBucket s3buckets.FlowLogsBucket `field:"optional" json:"flowLogsBucket" yaml:"flowLogsBucket"`
	// A cdk-extentions/ec2 {@link aws-ec2 !FlowLogFormat } object defining the desired formatting for Flow Logs.
	FlowLogsFormat ec2.FlowLogFormat `field:"optional" json:"flowLogsFormat" yaml:"flowLogsFormat"`
	// Boolean for adding "friendly names" for the created Athena queries.
	FriendlyQueryNames *bool `field:"optional" json:"friendlyQueryNames" yaml:"friendlyQueryNames"`
	// A cdk-extensions/s3-buckets {@link aws-s3-buckets !SesLogsBucket} object.
	SesLogsBucket s3buckets.SesLogsBucket `field:"optional" json:"sesLogsBucket" yaml:"sesLogsBucket"`
	// Boolean for using "standardized" naming (i.e. "aws-${service}-logs-${account} -${region}") for the created S3 Buckets.
	StandardizeNames *bool `field:"optional" json:"standardizeNames" yaml:"standardizeNames"`
	// A cdk-extensions/s3-buckets {@link aws-s3-buckets !WafLogsBucket} object.
	WafLogsBucket s3buckets.WafLogsBucket `field:"optional" json:"wafLogsBucket" yaml:"wafLogsBucket"`
	// Controls settings for an Athena WorkGroup used to query logs produced by AWS services.
	WorkGroupConfiguration *LoggingWorkGroupConfiguration `field:"optional" json:"workGroupConfiguration" yaml:"workGroupConfiguration"`
}

