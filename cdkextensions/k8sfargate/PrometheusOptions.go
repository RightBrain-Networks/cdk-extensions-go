package k8sfargate


// Optional configurations for the Prometheus resource.
type PrometheusOptions struct {
	// The Kubernetes namespace where the service should be deployed.
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
	// {@inheritdoc QueueConfiguration}.
	QueueConfiguration *QueueConfiguration `field:"optional" json:"queueConfiguration" yaml:"queueConfiguration"`
	// Name of the Kubernetes service account that should be created and used by Prometheus.
	ServiceAccountName *string `field:"optional" json:"serviceAccountName" yaml:"serviceAccountName"`
}

