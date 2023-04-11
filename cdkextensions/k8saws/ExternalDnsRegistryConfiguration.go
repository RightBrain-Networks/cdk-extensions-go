package k8saws

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
)

type ExternalDnsRegistryConfiguration struct {
	RegistryType *string `field:"required" json:"registryType" yaml:"registryType"`
	Permissions *[]awsiam.PolicyStatement `field:"optional" json:"permissions" yaml:"permissions"`
	Properties *map[string]interface{} `field:"optional" json:"properties" yaml:"properties"`
}

