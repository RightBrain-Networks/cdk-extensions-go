//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package glue

// Building without runtime type checking enabled, so all the below just return nil

func (w *jsiiProxy_WorkflowCrawlerPredicate) validateBindParameters(scope constructs.IConstruct) error {
	return nil
}

func (w *jsiiProxy_WorkflowCrawlerPredicate) validateBindOptionsParameters(_scope constructs.IConstruct) error {
	return nil
}

func validateNewWorkflowCrawlerPredicateParameters(crawler ICrawler, options *WorkflowCrawlerPredicateOptions) error {
	return nil
}

