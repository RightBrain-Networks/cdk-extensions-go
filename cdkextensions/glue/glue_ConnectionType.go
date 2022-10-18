package glue


type ConnectionType string

const (
	ConnectionType_JDBC ConnectionType = "JDBC"
	ConnectionType_KAFKA ConnectionType = "KAFKA"
	ConnectionType_MONGODB ConnectionType = "MONGODB"
	ConnectionType_NETWORK ConnectionType = "NETWORK"
)

