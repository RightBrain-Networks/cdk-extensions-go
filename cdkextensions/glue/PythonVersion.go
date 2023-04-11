package glue


type PythonVersion string

const (
	// Python 3 (the exact version depends on GlueVersion and JobCommand used).
	PythonVersion_THREE PythonVersion = "THREE"
	// Python 2 (the exact version depends on GlueVersion and JobCommand used).
	PythonVersion_TWO PythonVersion = "TWO"
)

