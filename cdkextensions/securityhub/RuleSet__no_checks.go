//go:build no_runtime_type_checking

package securityhub

// Building without runtime type checking enabled, so all the below just return nil

func (r *jsiiProxy_RuleSet) validateBindParameters(scope constructs.IConstruct) error {
	return nil
}

