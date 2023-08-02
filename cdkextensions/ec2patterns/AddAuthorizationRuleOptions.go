package ec2patterns

import (
	"github.com/aws/constructs-go/constructs/v10"
)

type AddAuthorizationRuleOptions struct {
	// The IPv4 address range, in CIDR notation, of the network for which access is being authorized.
	Cidr *string `field:"required" json:"cidr" yaml:"cidr"`
	// A brief description of the authorization rule.
	// Default: - no description.
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The ID of the group to grant access to, for example, the Active Directory group or identity provider (IdP) group.
	// Default: - authorize all groups.
	//
	GroupId *string `field:"optional" json:"groupId" yaml:"groupId"`
	Scope constructs.IConstruct `field:"optional" json:"scope" yaml:"scope"`
}

