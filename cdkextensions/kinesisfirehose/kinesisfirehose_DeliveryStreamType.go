package kinesisfirehose


type DeliveryStreamType string

const (
	DeliveryStreamType_DIRECT_PUT DeliveryStreamType = "DIRECT_PUT"
	DeliveryStreamType_KINESIS_STREAM_AS_SOURCE DeliveryStreamType = "KINESIS_STREAM_AS_SOURCE"
)

