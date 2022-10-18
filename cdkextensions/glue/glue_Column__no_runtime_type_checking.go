//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package glue

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_Column) validateBindParameters(scope constructs.IConstruct) error {
	return nil
}

func validateNewColumnParameters(props *ColumnProps) error {
	return nil
}

