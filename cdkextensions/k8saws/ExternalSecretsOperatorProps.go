package k8saws

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awseks"
)

// Configuration for the ExternalSecretsOperator resource.
type ExternalSecretsOperatorProps struct {
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
	// The EKS cluster where the external secrets operator should be installed.
	Cluster awseks.Cluster `field:"required" json:"cluster" yaml:"cluster"`
	// Determines the behavior when the service is deployed to a namespace that doesn't already exist on the EKS cluster.
	//
	// When this flag is `true` and the namespace doesn't exist, the namespace
	// will be created automatically.
	//
	// When this flag is `false` and the namespace doesn't exist, an error will
	// occur and resource creation will fail.
	CreateNamespace *bool `field:"optional" json:"createNamespace" yaml:"createNamespace"`
	// The Kubernetes namespace where the external secrets operator service should be installed and configured.
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
}

