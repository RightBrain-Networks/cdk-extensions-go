package ec2


// The file format options for flow log files delivered to S3.
// See: [Flow log files](https://docs.aws.amazon.com/vpc/latest/tgw/flow-logs-s3.html#flow-logs-s3-path)
//
type FlowLogFileFormat string

const (
	// Apache Parquet is a columnar data format.
	//
	// Queries on data in Parquet
	// format are 10 to 100 times faster compared to queries on data in plain
	// text. Data in Parquet format with Gzip compression takes 20 percent less
	// storage space than plain text with Gzip compression.
	FlowLogFileFormat_PARQUET FlowLogFileFormat = "PARQUET"
	// Plain text.
	//
	// This is the default format.
	FlowLogFileFormat_PLAIN_TEXT FlowLogFileFormat = "PLAIN_TEXT"
)

