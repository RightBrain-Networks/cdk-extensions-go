//go:build no_runtime_type_checking

package glue

// Building without runtime type checking enabled, so all the below just return nil

func validateWorkflowPredicate_CrawlerParameters(crawler ICrawler, options *WorkflowCrawlerPredicateOptions) error {
	return nil
}

func validateWorkflowPredicate_JobParameters(job IJob, options *WorkflowJobPredicateOptions) error {
	return nil
}

