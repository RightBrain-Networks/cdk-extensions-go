//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package glue

// Building without runtime type checking enabled, so all the below just return nil

func (p *jsiiProxy_PredicateBase) validateBindOptionsParameters(_scope constructs.IConstruct) error {
	return nil
}

func validateNewPredicateBaseParameters(_options *PredicateOptions) error {
	return nil
}

