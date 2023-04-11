package ec2


type FlowLogDataType string

const (
	// 32 bit signed int.
	FlowLogDataType_INT_32 FlowLogDataType = "INT_32"
	// 64 bit signed int.
	FlowLogDataType_INT_64 FlowLogDataType = "INT_64"
	// UTF-8 encoded character string.
	FlowLogDataType_STRING FlowLogDataType = "STRING"
)

