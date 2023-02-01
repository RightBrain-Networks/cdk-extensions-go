package k8sfargate


// Configures the queue used to write to Amazon Managed Service for Prometheus.
// See: [Remote write configuration](https://prometheus.io/docs/prometheus/latest/configuration/configuration/#remote_write)
//
type QueueConfiguration struct {
	// Number of samples to buffer per shard before we block reading of more samples from the WAL.
	//
	// It is recommended to have enough capacity in each
	// shard to buffer several requests to keep throughput up while processing
	// occasional slow remote requests.
	Capacity *float64 `field:"optional" json:"capacity" yaml:"capacity"`
	// Maximum number of samples per send.
	MaxSamplesPerSend *float64 `field:"optional" json:"maxSamplesPerSend" yaml:"maxSamplesPerSend"`
	// Maximum number of shards, i.e. amount of concurrency.
	MaxShards *float64 `field:"optional" json:"maxShards" yaml:"maxShards"`
}

