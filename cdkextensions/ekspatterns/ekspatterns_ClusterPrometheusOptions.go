package ekspatterns

import (
	"github.com/vibe-io/cdk-extensions-go/cdkextensions/aps"
	"github.com/vibe-io/cdk-extensions-go/cdkextensions/k8sfargate"
)

type ClusterPrometheusOptions struct {
	// The Kubernetes namespace where the service should be deployed.
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
	// {@inheritdoc QueueConfiguration}.
	QueueConfiguration *k8sfargate.QueueConfiguration `field:"optional" json:"queueConfiguration" yaml:"queueConfiguration"`
	// Name of the Kubernetes service account that should be created and used by Prometheus.
	ServiceAccountName *string `field:"optional" json:"serviceAccountName" yaml:"serviceAccountName"`
	Enabled *bool `field:"optional" json:"enabled" yaml:"enabled"`
	Workspace aps.IWorkspace `field:"optional" json:"workspace" yaml:"workspace"`
}

