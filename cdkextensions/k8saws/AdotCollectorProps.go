package k8saws

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awseks"
)

// Condifuration for the AdorCollector resource.
type AdotCollectorProps struct {
	// The AWS account ID this resource belongs to.
	// Default: - the resource is in the same account as the stack it belongs to.
	//
	Account *string `field:"optional" json:"account" yaml:"account"`
	// ARN to deduce region and account from.
	//
	// The ARN is parsed and the account and region are taken from the ARN.
	// This should be used for imported resources.
	//
	// Cannot be supplied together with either `account` or `region`.
	// Default: - take environment from `account`, `region` parameters, or use Stack environment.
	//
	EnvironmentFromArn *string `field:"optional" json:"environmentFromArn" yaml:"environmentFromArn"`
	// The value passed in by users to the physical name prop of the resource.
	//
	// - `undefined` implies that a physical name will be allocated by
	//   CloudFormation during deployment.
	// - a concrete value implies a specific physical name
	// - `PhysicalName.GENERATE_IF_NEEDED` is a marker that indicates that a physical will only be generated
	//   by the CDK if it is needed for cross-environment references. Otherwise, it will be allocated by CloudFormation.
	// Default: - The physical name will be allocated by CloudFormation at deployment time.
	//
	PhysicalName *string `field:"optional" json:"physicalName" yaml:"physicalName"`
	// The AWS region this resource belongs to.
	// Default: - the resource is in the same region as the stack it belongs to.
	//
	Region *string `field:"optional" json:"region" yaml:"region"`
	// The EKS cluster where the ADOT Collector will be deployed.
	Cluster awseks.ICluster `field:"required" json:"cluster" yaml:"cluster"`
	// Flag wich sets whether the deploy of the ADOT collector should include creating the Kubernetes namespace the service will be deployed to.
	// Default: true.
	//
	CreateNamespace *bool `field:"optional" json:"createNamespace" yaml:"createNamespace"`
	// The Kubernetes namespace where resources related to the ADOT collector will be created.
	// Default: {@link AdotCollector.DEFAULT_NAMESPACE }
	//
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
}

