package sso


// Configuration for Assignment resource.
type AssignmentProps struct {
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
	// The IAM Identity Center {@link aws-sso.IInstance | instance } under which the operation will be executed.
	Instance IInstance `field:"required" json:"instance" yaml:"instance"`
	// The permission set which governs the access being assigned.
	//
	// The
	// permission set grants the {@link principal} permissions on
	// {@link target}.
	PermissionSet IPermissionSet `field:"required" json:"permissionSet" yaml:"permissionSet"`
	// The IAM Identity Center principal you wish to grant permissions to.
	Principal IIdentityCenterPrincipal `field:"required" json:"principal" yaml:"principal"`
	// The resource you wish to grant the {@link principal} entity access to using the permissions defined in the {@link permissionSet}.
	//
	// For example,
	// an AWS account.
	Target AssignmentTarget `field:"required" json:"target" yaml:"target"`
}

