package glue


// State a Glue job must be in in order to satisfy a predicate condition to trigger a part of a workflow.
type JobState string

const (
	// A job that has finished and ended with an error.
	JobState_FAILED JobState = "FAILED"
	// A job which was stopped before completion.
	JobState_STOPPED JobState = "STOPPED"
	// A job which has finished successfully.
	JobState_SUCCEEDED JobState = "SUCCEEDED"
	// A job which timed out without completing.
	JobState_TIMEOUT JobState = "TIMEOUT"
)

