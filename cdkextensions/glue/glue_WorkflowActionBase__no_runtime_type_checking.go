//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package glue

// Building without runtime type checking enabled, so all the below just return nil

func (w *jsiiProxy_WorkflowActionBase) validateAddArgumentParameters(key *string, value *string) error {
	return nil
}

func (w *jsiiProxy_WorkflowActionBase) validateBindOptionsParameters(_scope constructs.IConstruct) error {
	return nil
}

func validateNewWorkflowActionBaseParameters(options *WorkflowActionOptions) error {
	return nil
}

