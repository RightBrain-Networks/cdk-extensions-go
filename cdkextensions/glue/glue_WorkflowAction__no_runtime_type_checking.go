//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package glue

// Building without runtime type checking enabled, so all the below just return nil

func validateWorkflowAction_CrawlerParameters(crawler ICrawler, options *WorkflowCrawlerActionOptions) error {
	return nil
}

func validateWorkflowAction_JobParameters(job IJob, options *WorkflowJobActionOptions) error {
	return nil
}

