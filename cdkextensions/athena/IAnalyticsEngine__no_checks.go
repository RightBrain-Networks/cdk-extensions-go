//go:build no_runtime_type_checking

package athena

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_IAnalyticsEngine) validateBindParameters(scope constructs.IConstruct, options *AnalyticsEngineBindProps) error {
	return nil
}

