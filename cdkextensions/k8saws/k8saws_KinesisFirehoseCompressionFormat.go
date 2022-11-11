package k8saws


type KinesisFirehoseCompressionFormat string

const (
	// The Apache Arrow compression format.
	//
	// Only available if the Fluent Fit service being used to send logs to
	// Firehose had Apache Arrow enabled at compile time.
	KinesisFirehoseCompressionFormat_ARROW KinesisFirehoseCompressionFormat = "ARROW"
	// Gzip compression format.
	KinesisFirehoseCompressionFormat_GZIP KinesisFirehoseCompressionFormat = "GZIP"
)

