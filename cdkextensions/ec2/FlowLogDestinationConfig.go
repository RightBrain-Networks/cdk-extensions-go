package ec2

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslogs"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
)

// A configuration object providing the details necessary to set up log delivery to a given destination.
type FlowLogDestinationConfig struct {
	// The type of destination for the flow log data.
	// See: [FlowLog LogDestinationType](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-flowlog.html#cfn-ec2-flowlog-logdestinationtype)
	//
	DestinationType awsec2.FlowLogDestinationType `field:"required" json:"destinationType" yaml:"destinationType"`
	// An S3 bucket where logs should be delivered.
	// See: [FlowLog LogDestination](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-flowlog.html#cfn-ec2-flowlog-logdestination)
	//
	Bucket awss3.IBucket `field:"optional" json:"bucket" yaml:"bucket"`
	// Additional options that control the format and behavior of logs delivered to the destination.
	DestinationOptions *map[string]interface{} `field:"optional" json:"destinationOptions" yaml:"destinationOptions"`
	// A CloudWatch LogGroup where logs should be delivered.
	// See: [FlowLog LogDestination](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-flowlog.html#cfn-ec2-flowlog-logdestination)
	//
	LogGroup awslogs.ILogGroup `field:"optional" json:"logGroup" yaml:"logGroup"`
	// The ARN of the IAM role that allows Amazon EC2 to publish flow logs in your account.
	// See: [FlowLog DeliverLogsPermissionArn](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-flowlog.html#cfn-ec2-flowlog-deliverlogspermissionarn)
	//
	Role awsiam.IRole `field:"optional" json:"role" yaml:"role"`
	// An Amazon Resource Name (ARN) for the S3 destination where log files are to be delivered.
	//
	// If a custom prefix is being added the ARN should reflect that prefix.
	// See: [FlowLog LogDestination](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-flowlog.html#cfn-ec2-flowlog-logdestination)
	//
	S3Path *string `field:"optional" json:"s3Path" yaml:"s3Path"`
}

