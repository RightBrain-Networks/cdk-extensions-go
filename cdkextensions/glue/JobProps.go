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
	//   CloudFormation during deployment.
	// - a concrete value implies a specific physical name
	// - `PhysicalName.GENERATE_IF_NEEDED` is a marker that indicates that a physical will only be generated
	//   by the CDK if it is needed for cross-environment references. Otherwise, it will be allocated by CloudFormation.
	PhysicalName *string `field:"optional" json:"physicalName" yaml:"physicalName"`
	// The AWS region this resource belongs to.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Executable properties for the Job.
	Executable JobExecutable `field:"required" json:"executable" yaml:"executable"`
	// The number of capacity units that are allocated to this job.
	// See: [AWS::Glue::Job](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-job.html#cfn-glue-job-allocatedcapacity)
	//
	AllocatedCapacity *float64 `field:"optional" json:"allocatedCapacity" yaml:"allocatedCapacity"`
	// List of Connections for use with this job.
	// See: [AWS::Glue::Job](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-job.html#cfn-glue-job-connections)
	//
	Connections *[]Connection `field:"optional" json:"connections" yaml:"connections"`
	// Set of properties for configuration of Continuous Logging.
	ContinuousLogging *ContinuousLoggingProps `field:"optional" json:"continuousLogging" yaml:"continuousLogging"`
	// The default arguments for this job, specified as name-value pairs.
	//
	// You can specify arguments here that your own job-execution script consumes, in addition to arguments that AWS Glue itself consumes.
	// See: [AWS::Glue::Job](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-job.html#cfn-glue-job-defaultarguments)
	//
	DefaultArguments *map[string]*string `field:"optional" json:"defaultArguments" yaml:"defaultArguments"`
	// A description of the job.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Boolean value for whether to enable Profiling Metrics.
	EnableProfilingMetrics *bool `field:"optional" json:"enableProfilingMetrics" yaml:"enableProfilingMetrics"`
	// The number of AWS Glue data processing units (DPUs) that can be allocated when this job runs.
	//
	// A DPU is a relative measure of processing power that consists of 4 vCPUs of compute capacity and 16 GB of memory.
	//
	// Do not set Max Capacity if using WorkerType and NumberOfWorkers.
	//
	// The value that can be allocated for MaxCapacity depends on whether you are running a Python shell job or an Apache Spark ETL job:
	//
	//    - When you specify a Python shell job (JobCommand.Name="pythonshell"), you can allocate either 0.0625 or 1 DPU. The default is 0.0625 DPU.
	//
	//    - When you specify an Apache Spark ETL job (JobCommand.Name="glueetl"), you can allocate from 2 to 100 DPUs. The default is 10 DPUs. This job type cannot have a fractional DPU allocation.
	MaxCapacity *float64 `field:"optional" json:"maxCapacity" yaml:"maxCapacity"`
	// Maximum number of concurrent executions.
	// See: [AWS::Glue::Job ExecutionProperty](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-job-executionproperty.html)
	//
	MaxConcurrentRuns *float64 `field:"optional" json:"maxConcurrentRuns" yaml:"maxConcurrentRuns"`
	// The maximum number of times to retry this job after a JobRun fails.
	// See: [AWS::Glue::Job](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-job.html#cfn-glue-job-maxretries)
	//
	MaxRetries *float64 `field:"optional" json:"maxRetries" yaml:"maxRetries"`
	// A name for the Job.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// After a job run starts, the number of minutes to wait before sending a job run delay notification.
	// See: [AWS::Glue::Job NotificationProperty](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-job-notificationproperty.html)
	//
	NotifyDelayAfter awscdk.Duration `field:"optional" json:"notifyDelayAfter" yaml:"notifyDelayAfter"`
	// The name or Amazon Resource Name (ARN) of the IAM role associated with this job.
	Role awsiam.IRole `field:"optional" json:"role" yaml:"role"`
	// The Security Configuration object to be applied to the Job.
	SecurityConfiguration SecurityConfiguration `field:"optional" json:"securityConfiguration" yaml:"securityConfiguration"`
	// The job timeout in minutes.
	//
	// This is the maximum time that a job run can consume resources before it is terminated and enters TIMEOUT status. The default is 2,880 minutes (48 hours).
	// See: [AWS::Glue::Job](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-job.html#cfn-glue-job-timeout)
	//
	Timeout awscdk.Duration `field:"optional" json:"timeout" yaml:"timeout"`
	// The number of worker available the Job.
	WorkerCount *float64 `field:"optional" json:"workerCount" yaml:"workerCount"`
	// The type of predefined worker that is allocated when a job runs.
	//
	// Accepts a value of Standard, G.1X, or G.2X.
	// See: [AWS::Glue::Job](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-job.html#cfn-glue-job-workertype)
	//
	WorkerType WorkerType `field:"optional" json:"workerType" yaml:"workerType"`
}

