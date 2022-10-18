//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package glue

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_StructColumn) validateAddColumnParameters(column Column) error {
	return nil
}

func (s *jsiiProxy_StructColumn) validateBindParameters(scope constructs.IConstruct) error {
	return nil
}

func validateNewStructColumnParameters(props *StructColumnProps) error {
	return nil
}

