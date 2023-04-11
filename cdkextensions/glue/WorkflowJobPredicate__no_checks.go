//go:build no_runtime_type_checking

package glue

// Building without runtime type checking enabled, so all the below just return nil

func (w *jsiiProxy_WorkflowJobPredicate) validateBindParameters(scope constructs.IConstruct) error {
	return nil
}

func (w *jsiiProxy_WorkflowJobPredicate) validateBindOptionsParameters(_scope constructs.IConstruct) error {
	return nil
}

func validateNewWorkflowJobPredicateParameters(job IJob, options *WorkflowJobPredicateOptions) error {
	return nil
}

