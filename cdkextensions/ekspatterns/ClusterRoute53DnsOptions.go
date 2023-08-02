package ekspatterns

import (
	"github.com/vibe-io/cdk-extensions-go/cdkextensions/k8saws"
)

type ClusterRoute53DnsOptions struct {
	// Maximum number of retries for AWS API calls before giving up.
	// Default: 3.
	//
	ApiRetries *float64 `field:"optional" json:"apiRetries" yaml:"apiRetries"`
	// Set the maximum number of changes that will be applied in each batch.
	// Default: 1000.
	//
	BatchChangeSize *float64 `field:"optional" json:"batchChangeSize" yaml:"batchChangeSize"`
	// Limits possible target zones by domain suffixes.
	DomainFilter *[]*string `field:"optional" json:"domainFilter" yaml:"domainFilter"`
	// Sets a flag determining whether the health of the backend service should be evaluated when determining DNS routing.
	EvaluateTargetHealth *bool `field:"optional" json:"evaluateTargetHealth" yaml:"evaluateTargetHealth"`
	// Sets the output format external dns will use when generating logs.
	// Default: {@link ExternalDnsLogLevel.JSON }
	//
	LogFormat k8saws.ExternalDnsLogFormat `field:"optional" json:"logFormat" yaml:"logFormat"`
	// Controls the verbosity of logs generated using the external-dns service.
	// Default: {@link ExternalDnsLogLevel.INFO }
	//
	LogLevel k8saws.ExternalDnsLogLevel `field:"optional" json:"logLevel" yaml:"logLevel"`
	// The Kubernetes namespace where the service should be deployed.
	// Default: 'kube-system'.
	//
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
	// When true, alias records will be avoided and CNAME records will be used instead.
	// Default: false.
	//
	PreferCname *bool `field:"optional" json:"preferCname" yaml:"preferCname"`
	// Registry specifying how ExternalDNS should track record ownership.
	//
	// Without a registry to track record ownership, External has no way to know
	// which records it owns and manages and which are owned and managed by a
	// different service.
	//
	// This can cause conflicts if there are multiple instances of External DNS
	// running or if there are other services managing DNS records in similar
	// zones as the different services could try to make conflicting changes due
	// to lacking a shared state.
	// Default: A TXT registry configured with defaults.
	//
	RecordOwnershipRegistry k8saws.IExternalDnsRegistry `field:"optional" json:"recordOwnershipRegistry" yaml:"recordOwnershipRegistry"`
	// Override the default region external-dns uses when calling AWS API's.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Desired number of ExternalDNS replicas.
	// Default: 1.
	//
	ReplicaCount *float64 `field:"optional" json:"replicaCount" yaml:"replicaCount"`
	// Controls the operations ExternalDNS will perform on the records it manages.
	// Default: {@link ExternalDnsSyncPolicy.SYNC }
	//
	SyncPolicy k8saws.ExternalDnsSyncPolicy `field:"optional" json:"syncPolicy" yaml:"syncPolicy"`
	// A set of tags that can be used to restrict which hosted zones external DNS will make changes to.
	ZoneTags *[]*k8saws.ExternalDnsZoneTag `field:"optional" json:"zoneTags" yaml:"zoneTags"`
	// Controls the types of hosted zones external-dns will create records for.
	// Default: ExternalDnsZoneType.ALL
	//
	ZoneType k8saws.ExternalDnsZoneType `field:"optional" json:"zoneType" yaml:"zoneType"`
	Enabled *bool `field:"optional" json:"enabled" yaml:"enabled"`
}

