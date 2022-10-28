package k8saws

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awseks"
)

// Configuration for the ExternalSecret resource.
type ExternalSecretProps struct {
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
	//    CloudFormation during deployment.
	// - a concrete value implies a specific physical name
	// - `PhysicalName.GENERATE_IF_NEEDED` is a marker that indicates that a physical will only be generated
	//    by the CDK if it is needed for cross-environment references. Otherwise, it will be allocated by CloudFormation.
	PhysicalName *string `field:"optional" json:"physicalName" yaml:"physicalName"`
	// The AWS region this resource belongs to.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// The EKS cluster where the secret should be created.
	Cluster awseks.ICluster `field:"required" json:"cluster" yaml:"cluster"`
	// The Kubernetes secret store resource that provides details and permissions to use for importing secrets from the provider.
	SecretStore ISecretStore `field:"required" json:"secretStore" yaml:"secretStore"`
	// The name to use for the Kubernetes secret resource when it is synchronized into the cluster.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The name where the synchronized secret should be created.
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
	// The frequency at which synchronization should occur.
	RefreshInterval awscdk.Duration `field:"optional" json:"refreshInterval" yaml:"refreshInterval"`
	// The secrets to synchronize into this Kubernetes secret.
	//
	// If multiple secrets are provided their fields will be merged.
	Secrets *[]ISecretReference `field:"optional" json:"secrets" yaml:"secrets"`
}

