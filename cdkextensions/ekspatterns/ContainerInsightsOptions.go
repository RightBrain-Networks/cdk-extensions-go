package ekspatterns


// Configuration options for enabling CloudWatch monitoring on the cluster.
type ContainerInsightsOptions struct {
	// Flag that controls whether CloudWatch Monitoring should be enabled or not.
	// Default: true.
	//
	Enabled *bool `field:"optional" json:"enabled" yaml:"enabled"`
	// The Kubernetes namespace where resources related to the the configuration of Container Insights will be created.
	// Default: {@link AdotCollector.DEFAULT_NAMESPACE }
	//
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
}

