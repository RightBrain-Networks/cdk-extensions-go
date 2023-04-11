package glue


type ConnectionType string

const (
	// JDBC - Designates a connection to a database through Java Database Connectivity (JDBC).
	ConnectionType_JDBC ConnectionType = "JDBC"
	// KAFKA - Designates a connection to an Apache Kafka streaming platform.
	ConnectionType_KAFKA ConnectionType = "KAFKA"
	// MONGODB - Designates a connection to a MongoDB document database.
	ConnectionType_MONGODB ConnectionType = "MONGODB"
	// NETWORK - Designates a network connection to a data source within an Amazon Virtual Private Cloud environment (Amazon VPC).
	ConnectionType_NETWORK ConnectionType = "NETWORK"
)

