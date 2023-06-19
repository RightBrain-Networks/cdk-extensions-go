package glue


// Result of binding a `JobExecutable` into a `Job`.
type JobExecutableConfig struct {
	// Glue version.
	// See: https://docs.aws.amazon.com/glue/latest/dg/release-notes.html
	//
	GlueVersion GlueVersion `field:"required" json:"glueVersion" yaml:"glueVersion"`
	// The language of the job (Scala or Python).
	// See:  `--job-language` in https://docs.aws.amazon.com/glue/latest/dg/aws-glue-programming-etl-glue-arguments.html
	//
	Language JobLanguage `field:"required" json:"language" yaml:"language"`
	// The script that is executed by a job.
	Script Code `field:"required" json:"script" yaml:"script"`
	// Specify the type of the job whether it's an Apache Spark ETL or streaming one or if it's a Python shell job.
	Type JobType `field:"required" json:"type" yaml:"type"`
	// The Scala class that serves as the entry point for the job.
	//
	// This applies only if your the job langauage is Scala.
	// See:  `--class` in https://docs.aws.amazon.com/glue/latest/dg/aws-glue-programming-etl-glue-arguments.html
	//
	ClassName *string `field:"optional" json:"className" yaml:"className"`
	// Additional files, such as configuration files that AWS Glue copies to the working directory of your script before executing it.
	// See:  `--extra-files` in https://docs.aws.amazon.com/glue/latest/dg/aws-glue-programming-etl-glue-arguments.html
	//
	ExtraFiles *[]Code `field:"optional" json:"extraFiles" yaml:"extraFiles"`
	// Additional Java .jar files that AWS Glue adds to the Java classpath before executing your script.
	// See:  `--extra-jars` in https://docs.aws.amazon.com/glue/latest/dg/aws-glue-programming-etl-glue-arguments.html
	//
	ExtraJars *[]Code `field:"optional" json:"extraJars" yaml:"extraJars"`
	// Setting this value to true prioritizes the customer's extra JAR files in the classpath.
	// See:  `--user-jars-first` in https://docs.aws.amazon.com/glue/latest/dg/aws-glue-programming-etl-glue-arguments.html
	//
	ExtraJarsFirst *bool `field:"optional" json:"extraJarsFirst" yaml:"extraJarsFirst"`
	// Additional Python files that AWS Glue adds to the Python path before executing your script.
	// See:  `--extra-py-files` in https://docs.aws.amazon.com/glue/latest/dg/aws-glue-programming-etl-glue-arguments.html
	//
	ExtraPythonFiles *[]Code `field:"optional" json:"extraPythonFiles" yaml:"extraPythonFiles"`
	// The Python version to use.
	PythonVersion PythonVersion `field:"optional" json:"pythonVersion" yaml:"pythonVersion"`
}

