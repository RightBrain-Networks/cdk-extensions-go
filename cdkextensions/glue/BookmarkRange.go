package glue


// A range of job run ID's that specify the job bookmark state of a Glue job which has had its bookmark state set to paused.
// See: [Using job bookmarks in AWS Glue](https://docs.aws.amazon.com/glue/latest/dg/monitor-continuations.html#monitor-continuations-implement)
//
type BookmarkRange struct {
	// The run ID which represents all the input that was processed until the last successful run before and including the specified run ID.
	//
	// The
	// corresponding input is ignored.
	From *string `field:"required" json:"from" yaml:"from"`
	// The run ID which represents all the input that was processed until the last successful run before and including the specified run ID.
	//
	// The
	// corresponding input excluding the input identified by the
	// {@link BookmarkRange.from | from} is processed by the job. Any input later
	// than this input is also excluded for processing.
	To *string `field:"required" json:"to" yaml:"to"`
}

