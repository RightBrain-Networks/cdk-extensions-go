package kinesisfirehose


type ParquetCompressionFormat string

const (
	ParquetCompressionFormat_GZIP ParquetCompressionFormat = "GZIP"
	ParquetCompressionFormat_SNAPPY ParquetCompressionFormat = "SNAPPY"
	ParquetCompressionFormat_UNCOMPRESSED ParquetCompressionFormat = "UNCOMPRESSED"
)

