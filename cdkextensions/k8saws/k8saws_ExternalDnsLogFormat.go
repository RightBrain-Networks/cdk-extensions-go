package k8saws


// The format external dns should use to output logs.
type ExternalDnsLogFormat string

const (
	// Output logs will be written as JSON objects.
	ExternalDnsLogFormat_JSON ExternalDnsLogFormat = "JSON"
	// Output logs will be written in plain text.
	ExternalDnsLogFormat_TEXT ExternalDnsLogFormat = "TEXT"
)

