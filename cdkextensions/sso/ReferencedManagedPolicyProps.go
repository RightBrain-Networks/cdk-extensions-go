package sso

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
)

// Represents configuration options when creating a managed policy from a class generated when adding a custom policy reference.
type ReferencedManagedPolicyProps struct {
	// A friendly description of the policy.
	//
	// Typically used to store information about the permissions defined in the
	// policy. For example, "Grants access to production DynamoDB tables."
	//
	// The policy description is immutable. After a value is assigned, it
	// cannot be changed.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The policy document that you want to use as the content for the new policy.
	Document awsiam.PolicyDocument `field:"optional" json:"document" yaml:"document"`
	// The groups to attach the policy to.
	//
	// When creating managed policies that will be referenced by IAM identity
	// center it is possible to associate them with other resources such as
	// users, groups, and roles. However, this is typically not done as IAM
	// Identity Center will handle configuring associations in the background.
	Groups *[]awsiam.IGroup `field:"optional" json:"groups" yaml:"groups"`
	// The roles to attach the policy to.
	//
	// When creating managed policies that will be referenced by IAM identity
	// center it is possible to associate them with other resources such as
	// users, groups, and roles. However, this is typically not done as IAM
	// Identity Center will handle configuring associations in the background.
	Roles *[]awsiam.IRole `field:"optional" json:"roles" yaml:"roles"`
	// Initial set of permissions to add to this policy document.
	Statements *[]awsiam.PolicyStatement `field:"optional" json:"statements" yaml:"statements"`
	// The users to attach the policy to.
	//
	// When creating managed policies that will be referenced by IAM identity
	// center it is possible to associate them with other resources such as
	// users, groups, and roles. However, this is typically not done as IAM
	// Identity Center will handle configuring associations in the background.
	Users *[]awsiam.IUser `field:"optional" json:"users" yaml:"users"`
}

