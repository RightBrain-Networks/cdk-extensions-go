package kinesisfirehose


type OrcCompressionFormat string

const (
	OrcCompressionFormat_NONE OrcCompressionFormat = "NONE"
	OrcCompressionFormat_SNAPPY OrcCompressionFormat = "SNAPPY"
	OrcCompressionFormat_ZLIB OrcCompressionFormat = "ZLIB"
)

