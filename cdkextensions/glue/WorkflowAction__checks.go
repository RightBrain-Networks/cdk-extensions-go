//go:build !no_runtime_type_checking

package glue

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func validateWorkflowAction_CrawlerParameters(crawler ICrawler, options *WorkflowCrawlerActionOptions) error {
	if crawler == nil {
		return fmt.Errorf("parameter crawler is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func validateWorkflowAction_JobParameters(job IJob, options *WorkflowJobActionOptions) error {
	if job == nil {
		return fmt.Errorf("parameter job is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

