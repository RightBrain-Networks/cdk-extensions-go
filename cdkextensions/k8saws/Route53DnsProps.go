package k8saws

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awseks"
)

// Full configuration for the Route53Dns resource.
type Route53DnsProps struct {
	// The AWS account ID this resource belongs to.
	Account *string `field:"optional" json:"account" yaml:"account"`
	// ARN to deduce region and account from.
	//
	// The ARN is parsed and the account and region are taken from the ARN.
	// This should be used for imported resources.
	//
	// Cannot be supplied together with either `account` or `region`.
	EnvironmentFromArn *string `field:"optional" json:"environmentFromArn" yaml:"environmentFromArn"`
	// The value passed in by users to the physical name prop of the resource.
	//
	// - `undefined` implies that a physical name will be allocated by
	//   CloudFormation during deployment.
	// - a concrete value implies a specific physical name
	// - `PhysicalName.GENERATE_IF_NEEDED` is a marker that indicates that a physical will only be generated
	//   by the CDK if it is needed for cross-environment references. Otherwise, it will be allocated by CloudFormation.
	PhysicalName *string `field:"optional" json:"physicalName" yaml:"physicalName"`
	// Override the default region external-dns uses when calling AWS API's.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Maximum number of retries for AWS API calls before giving up.
	ApiRetries *float64 `field:"optional" json:"apiRetries" yaml:"apiRetries"`
	// Set the maximum number of changes that will be applied in each batch.
	BatchChangeSize *float64 `field:"optional" json:"batchChangeSize" yaml:"batchChangeSize"`
	// Limits possible target zones by domain suffixes.
	DomainFilter *[]*string `field:"optional" json:"domainFilter" yaml:"domainFilter"`
	// Sets a flag determining whether the health of the backend service should be evaluated when determining DNS routing.
	EvaluateTargetHealth *bool `field:"optional" json:"evaluateTargetHealth" yaml:"evaluateTargetHealth"`
	// Sets the output format external dns will use when generating logs.
	LogFormat ExternalDnsLogFormat `field:"optional" json:"logFormat" yaml:"logFormat"`
	// Controls the verbosity of logs generated using the external-dns service.
	LogLevel ExternalDnsLogLevel `field:"optional" json:"logLevel" yaml:"logLevel"`
	// The Kubernetes namespace where the service should be deployed.
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
	// When true, alias records will be avoided and CNAME records will be used instead.
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
	RecordOwnershipRegistry IExternalDnsRegistry `field:"optional" json:"recordOwnershipRegistry" yaml:"recordOwnershipRegistry"`
	// Desired number of ExternalDNS replicas.
	ReplicaCount *float64 `field:"optional" json:"replicaCount" yaml:"replicaCount"`
	// Controls the operations ExternalDNS will perform on the records it manages.
	SyncPolicy ExternalDnsSyncPolicy `field:"optional" json:"syncPolicy" yaml:"syncPolicy"`
	// A set of tags that can be used to restrict which hosted zones external DNS will make changes to.
	ZoneTags *[]*ExternalDnsZoneTag `field:"optional" json:"zoneTags" yaml:"zoneTags"`
	// Controls the types of hosted zones external-dns will create records for.
	ZoneType ExternalDnsZoneType `field:"optional" json:"zoneType" yaml:"zoneType"`
	// The EKS cluster where external-dns should be deployed.
	Cluster awseks.ICluster `field:"required" json:"cluster" yaml:"cluster"`
}

