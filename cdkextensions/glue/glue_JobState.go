package glue


type JobState string

const (
	JobState_FAILED JobState = "FAILED"
	JobState_STOPPED JobState = "STOPPED"
	JobState_SUCCEEDED JobState = "SUCCEEDED"
	JobState_TIMEOUT JobState = "TIMEOUT"
)

