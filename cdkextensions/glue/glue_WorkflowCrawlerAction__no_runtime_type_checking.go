//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package glue

// Building without runtime type checking enabled, so all the below just return nil

func (w *jsiiProxy_WorkflowCrawlerAction) validateAddArgumentParameters(key *string, value *string) error {
	return nil
}

func (w *jsiiProxy_WorkflowCrawlerAction) validateBindParameters(scope constructs.IConstruct) error {
	return nil
}

func (w *jsiiProxy_WorkflowCrawlerAction) validateBindOptionsParameters(_scope constructs.IConstruct) error {
	return nil
}

func validateNewWorkflowCrawlerActionParameters(crawler ICrawler, options *WorkflowCrawlerActionOptions) error {
	return nil
}

