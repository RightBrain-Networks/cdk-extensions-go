package k8saws

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awseks"
	"github.com/vibe-io/cdk-extensions-go/cdkextensions/route53"
)

// Configuration for the Echoserver resource.
type EchoserverProps struct {
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
	// The EKS Cluster where the service should be deployed.
	Cluster awseks.ICluster `field:"required" json:"cluster" yaml:"cluster"`
	// Determines the behavior of automatic DNS discovery and configuration.
	DomainDiscovery route53.DomainDiscovery `field:"optional" json:"domainDiscovery" yaml:"domainDiscovery"`
	// The subnets where the load balancer should be created.
	LoadBalancerSubnets *awsec2.SubnetSelection `field:"optional" json:"loadBalancerSubnets" yaml:"loadBalancerSubnets"`
	// The name of the Kubernetes service to be created.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The Kubernetes namespace where the service should be created.
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
	// The port which netcat should listen on.
	Port *float64 `field:"optional" json:"port" yaml:"port"`
	// The number of replicas that should exist.
	Replicas *float64 `field:"optional" json:"replicas" yaml:"replicas"`
	// The Security groups which should be applied to the service.
	SecurityGroups *[]awsec2.ISecurityGroup `field:"optional" json:"securityGroups" yaml:"securityGroups"`
	// A subdomain that should be prefixed to the beginning of all registered domains.
	Subdomain *string `field:"optional" json:"subdomain" yaml:"subdomain"`
	// The Docker tag specifying the version of echoserver to use.
	// See: [Google echoserver image repository](https://console.cloud.google.com/gcr/images/google-containers/GLOBAL/echoserver)
	//
	Tag *string `field:"optional" json:"tag" yaml:"tag"`
}

