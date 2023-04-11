package k8saws


// Controls the operations ExternalDNS will perform on the records it manages.
type ExternalDnsSyncPolicy string

const (
	// Full sync mode.
	//
	// Records will be created, updated, and deleted based on the
	// statis of their backing resources on the Kubernetes cluster.
	ExternalDnsSyncPolicy_SYNC ExternalDnsSyncPolicy = "SYNC"
	// Only allow create and update operations.
	//
	// Records will have their values
	// set based on the status of their backing Kubernetes resources, however if
	// those resources are removed the DNS records will be retained, set to their
	// last configured value.
	ExternalDnsSyncPolicy_UPSERT_ONLY ExternalDnsSyncPolicy = "UPSERT_ONLY"
)

