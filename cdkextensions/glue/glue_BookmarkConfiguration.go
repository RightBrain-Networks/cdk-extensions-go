package glue

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/vibe-io/cdk-extensions-go/cdkextensions/jsii"
)

// Controls the bookmark state of a Glue job.
// See: [Using job bookmarks in AWS Glue](https://docs.aws.amazon.com/glue/latest/dg/monitor-continuations.html#monitor-continuations-implement)
//
type BookmarkConfiguration interface {
	// An optional range of job ID's that will correspond to the `job-bookmark-from` and `job-bookmark-to` arguments.
	Range() *BookmarkRange
	// The value to pass to the `job-bookmark-option` argument.
	Value() *string
}

// The jsii proxy struct for BookmarkConfiguration
type jsiiProxy_BookmarkConfiguration struct {
	_ byte // padding
}

func (j *jsiiProxy_BookmarkConfiguration) Range() *BookmarkRange {
	var returns *BookmarkRange
	_jsii_.Get(
		j,
		"range",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BookmarkConfiguration) Value() *string {
	var returns *string
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


// Job bookmarks are not used, and the job always processes the entire dataset.
//
// You are responsible for managing the output from previous job
// runs.
//
// Returns: A configuration object that disabled job bookmarks.
func BookmarkConfiguration_Disable() BookmarkConfiguration {
	_init_.Initialize()

	var returns BookmarkConfiguration

	_jsii_.StaticInvoke(
		"cdk-extensions.glue.BookmarkConfiguration",
		"disable",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Causes the job to update the state after a run to keep track of previously processed data.
//
// If your job has a source with job bookmark support, it
// will keep track of processed data, and when a job runs, it processes new
// data since the last checkpoint.
//
// Returns: A configuration object that enables job bookmarks.
func BookmarkConfiguration_Enable() BookmarkConfiguration {
	_init_.Initialize()

	var returns BookmarkConfiguration

	_jsii_.StaticInvoke(
		"cdk-extensions.glue.BookmarkConfiguration",
		"enable",
		nil, // no parameters
		&returns,
	)

	return returns
}

// An escape hatch method that allows specifying arbitrary values for the `job-bookmark-option` argument of a Glue job.
//
// Returns: A configuration object that represents the provided bookmark configuration.
func BookmarkConfiguration_Of(value *string, range_ *BookmarkRange) BookmarkConfiguration {
	_init_.Initialize()

	if err := validateBookmarkConfiguration_OfParameters(value, range_); err != nil {
		panic(err)
	}
	var returns BookmarkConfiguration

	_jsii_.StaticInvoke(
		"cdk-extensions.glue.BookmarkConfiguration",
		"of",
		[]interface{}{value, range_},
		&returns,
	)

	return returns
}

// Process incremental data since the last successful run or the data in a specified range, without updating the state of last bookmark.
//
// You are
// responsible for managing the output from previous job runs.
//
// Returns: A configuration object that pauses job bookmarks.
func BookmarkConfiguration_Pause(range_ *BookmarkRange) BookmarkConfiguration {
	_init_.Initialize()

	if err := validateBookmarkConfiguration_PauseParameters(range_); err != nil {
		panic(err)
	}
	var returns BookmarkConfiguration

	_jsii_.StaticInvoke(
		"cdk-extensions.glue.BookmarkConfiguration",
		"pause",
		[]interface{}{range_},
		&returns,
	)

	return returns
}

