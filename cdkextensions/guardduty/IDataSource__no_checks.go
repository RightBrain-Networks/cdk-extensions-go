//go:build no_runtime_type_checking

package guardduty

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_IDataSource) validateBindParameters(scope constructs.IConstruct) error {
	return nil
}

