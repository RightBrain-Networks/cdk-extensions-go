package glue

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
)

// Configuration for the Glue Job resource.
type JobProps struct {
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
	Executable JobExecutable `field:"required" json:"executable" yaml:"executable"`
	AllocatedCapacity *float64 `field:"optional" json:"allocatedCapacity" yaml:"allocatedCapacity"`
	Connections *[]Connection `field:"optional" json:"connections" yaml:"connections"`
	ContinuousLogging *ContinuousLoggingProps `field:"optional" json:"continuousLogging" yaml:"continuousLogging"`
	DefaultArguments *map[string]*string `field:"optional" json:"defaultArguments" yaml:"defaultArguments"`
	Description *string `field:"optional" json:"description" yaml:"description"`
	EnableProfilingMetrics *bool `field:"optional" json:"enableProfilingMetrics" yaml:"enableProfilingMetrics"`
	MaxCapacity *float64 `field:"optional" json:"maxCapacity" yaml:"maxCapacity"`
	MaxConcurrentRuns *float64 `field:"optional" json:"maxConcurrentRuns" yaml:"maxConcurrentRuns"`
	MaxRetries *float64 `field:"optional" json:"maxRetries" yaml:"maxRetries"`
	Name *string `field:"optional" json:"name" yaml:"name"`
	NotifyDelayAfter awscdk.Duration `field:"optional" json:"notifyDelayAfter" yaml:"notifyDelayAfter"`
	Role awsiam.IRole `field:"optional" json:"role" yaml:"role"`
	SecurityConfiguration SecurityConfiguration `field:"optional" json:"securityConfiguration" yaml:"securityConfiguration"`
	Timeout awscdk.Duration `field:"optional" json:"timeout" yaml:"timeout"`
	WorkerCount *float64 `field:"optional" json:"workerCount" yaml:"workerCount"`
	WorkerType WorkerType `field:"optional" json:"workerType" yaml:"workerType"`
}

