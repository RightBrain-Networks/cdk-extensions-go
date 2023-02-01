package k8sfargate

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awseks"
	"github.com/vibe-io/cdk-extensions-go/cdkextensions/aps"
)

type PrometheusProps struct {
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
	// The Kubernetes namespace where the service should be deployed.
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
	// {@inheritdoc QueueConfiguration}.
	QueueConfiguration *QueueConfiguration `field:"optional" json:"queueConfiguration" yaml:"queueConfiguration"`
	// Name of the Kubernetes service account that should be created and used by Prometheus.
	ServiceAccountName *string `field:"optional" json:"serviceAccountName" yaml:"serviceAccountName"`
	// The EKS cluster where Prometheus should be deployed.
	Cluster awseks.Cluster `field:"required" json:"cluster" yaml:"cluster"`
	// The Amazon Managed Service for Prometheus workspace where the Prometheus server should sned its data.
	Workspace aps.IWorkspace `field:"required" json:"workspace" yaml:"workspace"`
}

