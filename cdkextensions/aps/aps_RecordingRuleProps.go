package aps


// Options needed to configure a Prometheus recording rule inside an APS rules configuration.
type RecordingRuleProps struct {
	// The PromQL expression to evaluate.
	//
	// Every evaluation cycle this is
	// evaluated at the current time, and the result recorded as a new set of
	// time series with the metric name as given by `record`.
	// See: [Querying prometheus](https://prometheus.io/docs/prometheus/latest/querying/basics/)
	//
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// The name of the time series to output to.
	//
	// Must be a valid metric name.
	Record *string `field:"required" json:"record" yaml:"record"`
	// Labels to add or overwrite before storing the result.
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
}

