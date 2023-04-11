package ec2


type FlowLogS3Options struct {
	// The file format in which flow logs should be delivered to S3.
	// See: [Flow log files](https://docs.aws.amazon.com/vpc/latest/tgw/flow-logs-s3.html#flow-logs-s3-path)
	//
	FileFormat FlowLogFileFormat `field:"optional" json:"fileFormat" yaml:"fileFormat"`
	// Controls the format of partitions ("folders") when the flow logs are delivered to S3.
	//
	// By default, flow logs are delivered partitioned such that each part of
	// the S3 path represents a values pertaining to details of the log.
	//
	// When hive compatible partitions are enabled, partitions will be
	// structured such that keys declaring the partition name are added at
	// each level.
	//
	// An example of standard partitioning:
	// ```
	// /us-east-1/2020/03/08/log.tar.gz
	// ```
	//
	// An example with Hive compatible partitions:
	// ```
	// /region=us-east-1/year=2020/month=03/day=08/log.tar.gz
	// ```.
	// See: [AWS Big Data Blog](https://aws.amazon.com/blogs/big-data/optimize-performance-and-reduce-costs-for-network-analytics-with-vpc-flow-logs-in-apache-parquet-format/)
	//
	HiveCompatiblePartitions *bool `field:"optional" json:"hiveCompatiblePartitions" yaml:"hiveCompatiblePartitions"`
	// An optional prefix that will be added to the start of all flow log files delivered to the S3 bucket.
	// See: [FlowLog LogDestination](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-flowlog.html#cfn-ec2-flowlog-logdestination)
	//
	KeyPrefix *string `field:"optional" json:"keyPrefix" yaml:"keyPrefix"`
	// Indicates whether to partition the flow log per hour.
	//
	// By default, flow logs are partitioned (organized into S3 "folders") by
	// day.
	//
	// Setting this to true will add an extra layer of directories splitting
	// flow log files by the hour in which they were delivered.
	// See: [Flow log files](https://docs.aws.amazon.com/vpc/latest/tgw/flow-logs-s3.html#flow-logs-s3-path)
	//
	PerHourPartition *bool `field:"optional" json:"perHourPartition" yaml:"perHourPartition"`
}

