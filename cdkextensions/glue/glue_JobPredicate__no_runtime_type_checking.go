//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package glue

// Building without runtime type checking enabled, so all the below just return nil

func (j *jsiiProxy_JobPredicate) validateBindParameters(scope constructs.IConstruct) error {
	return nil
}

func (j *jsiiProxy_JobPredicate) validateBindOptionsParameters(_scope constructs.IConstruct) error {
	return nil
}

func validateNewJobPredicateParameters(job IJob, options *JobPredicateOptions) error {
	return nil
}

