package k8saws


// The modes that the Fluent Bit Nest filter plugin can work in.
type NestFilterOperationType string

const (
	// Lift data from a nested object.
	NestFilterOperationType_LIFT NestFilterOperationType = "LIFT"
	// Nest data into a specified object.
	NestFilterOperationType_NEST NestFilterOperationType = "NEST"
)

