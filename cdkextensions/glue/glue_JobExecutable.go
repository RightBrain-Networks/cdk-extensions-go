package glue

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/vibe-io/cdk-extensions-go/cdkextensions/jsii"
)

// The executable properties related to the Glue job's GlueVersion, JobType and code.
type JobExecutable interface {
	// Called during Job initialization to get JobExecutableConfig.
	Bind() *JobExecutableConfig
}

// The jsii proxy struct for JobExecutable
type jsiiProxy_JobExecutable struct {
	_ byte // padding
}

// Create a custom JobExecutable.
func JobExecutable_Of(config *JobExecutableConfig) JobExecutable {
	_init_.Initialize()

	if err := validateJobExecutable_OfParameters(config); err != nil {
		panic(err)
	}
	var returns JobExecutable

	_jsii_.StaticInvoke(
		"cdk-extensions.glue.JobExecutable",
		"of",
		[]interface{}{config},
		&returns,
	)

	return returns
}

// Create Python executable props for Apache Spark ETL job.
func JobExecutable_PythonEtl(props *PythonSparkJobExecutableProps) JobExecutable {
	_init_.Initialize()

	if err := validateJobExecutable_PythonEtlParameters(props); err != nil {
		panic(err)
	}
	var returns JobExecutable

	_jsii_.StaticInvoke(
		"cdk-extensions.glue.JobExecutable",
		"pythonEtl",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Create Python executable props for python shell jobs.
func JobExecutable_PythonShell(props *PythonShellExecutableProps) JobExecutable {
	_init_.Initialize()

	if err := validateJobExecutable_PythonShellParameters(props); err != nil {
		panic(err)
	}
	var returns JobExecutable

	_jsii_.StaticInvoke(
		"cdk-extensions.glue.JobExecutable",
		"pythonShell",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Create Python executable props for Apache Spark Streaming job.
func JobExecutable_PythonStreaming(props *PythonSparkJobExecutableProps) JobExecutable {
	_init_.Initialize()

	if err := validateJobExecutable_PythonStreamingParameters(props); err != nil {
		panic(err)
	}
	var returns JobExecutable

	_jsii_.StaticInvoke(
		"cdk-extensions.glue.JobExecutable",
		"pythonStreaming",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Create Scala executable props for Apache Spark ETL job.
func JobExecutable_ScalaEtl(props *ScalaJobExecutableProps) JobExecutable {
	_init_.Initialize()

	if err := validateJobExecutable_ScalaEtlParameters(props); err != nil {
		panic(err)
	}
	var returns JobExecutable

	_jsii_.StaticInvoke(
		"cdk-extensions.glue.JobExecutable",
		"scalaEtl",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Create Scala executable props for Apache Spark Streaming job.
func JobExecutable_ScalaStreaming(props *ScalaJobExecutableProps) JobExecutable {
	_init_.Initialize()

	if err := validateJobExecutable_ScalaStreamingParameters(props); err != nil {
		panic(err)
	}
	var returns JobExecutable

	_jsii_.StaticInvoke(
		"cdk-extensions.glue.JobExecutable",
		"scalaStreaming",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JobExecutable) Bind() *JobExecutableConfig {
	var returns *JobExecutableConfig

	_jsii_.Invoke(
		j,
		"bind",
		nil, // no parameters
		&returns,
	)

	return returns
}

