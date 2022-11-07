package sso

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
)

// Configuration for PermissionSet resource.
type PermissionSetProps struct {
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
	// The ARN of the IAM Identity Center instance under which the operation will be executed.
	Instance IInstance `field:"required" json:"instance" yaml:"instance"`
	// A user friendly description providing details about the permission set.
	// See: [AWS::SSO::PermissionSet](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sso-permissionset.html#cfn-sso-permissionset-description)
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Adds inline policy documents that will be embedded in the permission set.
	// See: [AWS::SSO::PermissionSet](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sso-permissionset.html#cfn-sso-permissionset-inlinepolicy)
	//
	InlinePolicies *map[string]awsiam.PolicyDocument `field:"optional" json:"inlinePolicies" yaml:"inlinePolicies"`
	// A list of the IAM managed policies that you want to attach to the permission set.
	//
	// Managed policies specified here must be AWS managed.
	// To reference custom managed policies use the {@link PermissionSet.addCustomerManagedPolicy}
	// method.
	// See: [AWS::SSO::PermissionSet](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sso-permissionset.html#cfn-sso-permissionset-managedpolicies)
	//
	ManagedPolicies *[]awsiam.IManagedPolicy `field:"optional" json:"managedPolicies" yaml:"managedPolicies"`
	// The name of the permission set.
	// See: [AWS::SSO::PermissionSet](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sso-permissionset.html#cfn-sso-permissionset-name)
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Specifies the configuration of the AWS managed or customer managed policy that you want to set as a permissions boundary.
	//
	// Specify either
	// CustomerManagedPolicyReference to use the name and path of a customer
	// managed policy, or ManagedPolicyArn to use the ARN of an AWS managed
	// policy. A permissions boundary represents the maximum permissions that
	// any policy can grant your role. For more information, see [Permissions
	// boundaries](https://docs.aws.amazon.com/IAM/latest/UserGuide/access_policies_boundaries.html) for IAM entities in the AWS Identity and Access Management
	// User Guide.
	// See: [AWS::SSO::PermissionSet](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sso-permissionset.html#cfn-sso-permissionset-permissionsboundary)
	//
	PermissionsBoundary PermissionsBoundary `field:"optional" json:"permissionsBoundary" yaml:"permissionsBoundary"`
	// Used to redirect users within the application during the federation authentication process.
	//
	// For example, you can redirect users to a
	// specific page that is most applicable to their job after singing in to
	// an AWS account.
	// See: [AWS::SSO::PermissionSet](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sso-permissionset.html#cfn-sso-permissionset-relaystatetype)
	//
	RelayState *string `field:"optional" json:"relayState" yaml:"relayState"`
	// The length of time that the application user sessions are valid for.
	// See: [AWS::SSO::PermissionSet](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sso-permissionset.html#cfn-sso-permissionset-sessionduration)
	//
	SessionDuration awscdk.Duration `field:"optional" json:"sessionDuration" yaml:"sessionDuration"`
}

