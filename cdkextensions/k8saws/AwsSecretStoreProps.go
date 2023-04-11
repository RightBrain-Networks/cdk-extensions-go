package k8saws

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awseks"
)

// Configuration options for adding a new secret store resource.
type AwsSecretStoreProps struct {
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
	// The AWS region this resource belongs to.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// The EKS cluster where the secret store should be created.
	Cluster awseks.ICluster `field:"required" json:"cluster" yaml:"cluster"`
	// The name of the service provider backing the secret store.
	Service *string `field:"required" json:"service" yaml:"service"`
	// A human friendly name for the secret store.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The Kubernetes namespace where the secret store should be created.
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
}

