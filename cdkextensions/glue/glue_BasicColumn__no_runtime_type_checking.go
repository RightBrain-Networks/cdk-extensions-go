//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package glue

// Building without runtime type checking enabled, so all the below just return nil

func (b *jsiiProxy_BasicColumn) validateBindParameters(scope constructs.IConstruct) error {
	return nil
}

func validateNewBasicColumnParameters(props *BasicColumnProps) error {
	return nil
}

