package glue

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/vibe-io/cdk-extensions-go/cdkextensions/jsii"
)

// The job type.
//
// If you need to use a JobType that doesn't exist as a static member, you
// can instantiate a `JobType` object, e.g: `JobType.of('other name')`.
type JobType interface {
	// The name of this JobType, as expected by Job resource.
	Name() *string
}

// The jsii proxy struct for JobType
type jsiiProxy_JobType struct {
	_ byte // padding
}

func (j *jsiiProxy_JobType) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}


// Custom type name.
func JobType_Of(name *string) JobType {
	_init_.Initialize()

	if err := validateJobType_OfParameters(name); err != nil {
		panic(err)
	}
	var returns JobType

	_jsii_.StaticInvoke(
		"cdk-extensions.glue.JobType",
		"of",
		[]interface{}{name},
		&returns,
	)

	return returns
}

func JobType_ETL() JobType {
	_init_.Initialize()
	var returns JobType
	_jsii_.StaticGet(
		"cdk-extensions.glue.JobType",
		"ETL",
		&returns,
	)
	return returns
}

func JobType_PYTHON_SHELL() JobType {
	_init_.Initialize()
	var returns JobType
	_jsii_.StaticGet(
		"cdk-extensions.glue.JobType",
		"PYTHON_SHELL",
		&returns,
	)
	return returns
}

func JobType_STREAMING() JobType {
	_init_.Initialize()
	var returns JobType
	_jsii_.StaticGet(
		"cdk-extensions.glue.JobType",
		"STREAMING",
		&returns,
	)
	return returns
}

