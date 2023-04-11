package k8saws


// Options for fetching tags/labels from provider secrets.
type MetadataPolicy string

const (
	// Fetch tags/labels from provider secrets.
	MetadataPolicy_FETCH MetadataPolicy = "FETCH"
	// Do not fetch tags/labels from provider secrets.
	MetadataPolicy_NONE MetadataPolicy = "NONE"
)

