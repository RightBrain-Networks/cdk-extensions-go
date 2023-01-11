package glue


// State a Glue crawler must be in in order to satisfy a predicate condition to trigger a part of a workflow.
type CrawlerState string

const (
	// A crawler execution was cancelled before it could finish.
	CrawlerState_CANCELLED CrawlerState = "CANCELLED"
	// A crawler that has finished and ended in an error.
	CrawlerState_FAILED CrawlerState = "FAILED"
	// A crawler which has finished successfully.
	CrawlerState_SUCCEEDED CrawlerState = "SUCCEEDED"
)

