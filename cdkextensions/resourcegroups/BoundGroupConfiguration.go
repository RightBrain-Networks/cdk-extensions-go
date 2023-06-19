package resourcegroups

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsresourcegroups"
)

type BoundGroupConfiguration struct {
	Configuration *[]*awsresourcegroups.CfnGroup_ConfigurationItemProperty `field:"optional" json:"configuration" yaml:"configuration"`
	Query *awsresourcegroups.CfnGroup_ResourceQueryProperty `field:"optional" json:"query" yaml:"query"`
}

