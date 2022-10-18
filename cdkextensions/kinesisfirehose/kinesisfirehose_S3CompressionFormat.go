package kinesisfirehose


type S3CompressionFormat string

const (
	S3CompressionFormat_GZIP S3CompressionFormat = "GZIP"
	S3CompressionFormat_HADOOP_SNAPPY S3CompressionFormat = "HADOOP_SNAPPY"
	S3CompressionFormat_SNAPPY S3CompressionFormat = "SNAPPY"
	S3CompressionFormat_UNCOMPRESSED S3CompressionFormat = "UNCOMPRESSED"
	S3CompressionFormat_ZIP S3CompressionFormat = "ZIP"
)

