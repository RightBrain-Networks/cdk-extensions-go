package k8saws


// Controls the types of Hosted Zones external DNS will create records for.
type ExternalDnsZoneType string

const (
	// Create DNS records for both public and private hosted zones.
	ExternalDnsZoneType_ALL ExternalDnsZoneType = "ALL"
	// Only create DNS records for private hosted zones.
	ExternalDnsZoneType_PRIVATE ExternalDnsZoneType = "PRIVATE"
	// Only create DNS records for public hosted zones.
	ExternalDnsZoneType_PUBLIC ExternalDnsZoneType = "PUBLIC"
)

