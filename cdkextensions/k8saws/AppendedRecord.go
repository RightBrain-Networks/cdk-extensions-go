package k8saws


// Represents a record field to be added by the record modifier Fluent Bit filter plugin.
type AppendedRecord struct {
	// The name of the field to be added.
	FieldName *string `field:"required" json:"fieldName" yaml:"fieldName"`
	// The value that the added field should be set to.
	Value *string `field:"required" json:"value" yaml:"value"`
}

