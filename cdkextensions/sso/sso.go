package sso

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"cdk-extensions.sso.AccessControlAttribute",
		reflect.TypeOf((*AccessControlAttribute)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addSource", GoMethod: "AddSource"},
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "sources", GoGetter: "Sources"},
		},
		func() interface{} {
			return &jsiiProxy_AccessControlAttribute{}
		},
	)
	_jsii_.RegisterStruct(
		"cdk-extensions.sso.AccessControlAttributeOptions",
		reflect.TypeOf((*AccessControlAttributeOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk-extensions.sso.Assignment",
		reflect.TypeOf((*Assignment)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "instance", GoGetter: "Instance"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "permissionSet", GoGetter: "PermissionSet"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "principal", GoGetter: "Principal"},
			_jsii_.MemberProperty{JsiiProperty: "resource", GoGetter: "Resource"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "target", GoGetter: "Target"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_Assignment{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdk-extensions.sso.AssignmentProps",
		reflect.TypeOf((*AssignmentProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk-extensions.sso.AssignmentTarget",
		reflect.TypeOf((*AssignmentTarget)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "targetId", GoGetter: "TargetId"},
			_jsii_.MemberProperty{JsiiProperty: "targetType", GoGetter: "TargetType"},
		},
		func() interface{} {
			return &jsiiProxy_AssignmentTarget{}
		},
	)
	_jsii_.RegisterClass(
		"cdk-extensions.sso.AssignmentTargetType",
		reflect.TypeOf((*AssignmentTargetType)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
		},
		func() interface{} {
			return &jsiiProxy_AssignmentTargetType{}
		},
	)
	_jsii_.RegisterClass(
		"cdk-extensions.sso.Group",
		reflect.TypeOf((*Group)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_Group{}
		},
	)
	_jsii_.RegisterClass(
		"cdk-extensions.sso.GroupBase",
		reflect.TypeOf((*GroupBase)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "groupId", GoGetter: "GroupId"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "principalId", GoGetter: "PrincipalId"},
			_jsii_.MemberProperty{JsiiProperty: "principalType", GoGetter: "PrincipalType"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_GroupBase{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkResource)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IGroup)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IIdentityCenterPrincipal)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"cdk-extensions.sso.IGroup",
		reflect.TypeOf((*IGroup)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "groupId", GoGetter: "GroupId"},
		},
		func() interface{} {
			return &jsiiProxy_IGroup{}
		},
	)
	_jsii_.RegisterInterface(
		"cdk-extensions.sso.IIdentityCenterPrincipal",
		reflect.TypeOf((*IIdentityCenterPrincipal)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "principalId", GoGetter: "PrincipalId"},
			_jsii_.MemberProperty{JsiiProperty: "principalType", GoGetter: "PrincipalType"},
		},
		func() interface{} {
			return &jsiiProxy_IIdentityCenterPrincipal{}
		},
	)
	_jsii_.RegisterInterface(
		"cdk-extensions.sso.IInstance",
		reflect.TypeOf((*IInstance)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "instanceArn", GoGetter: "InstanceArn"},
			_jsii_.MemberProperty{JsiiProperty: "instanceId", GoGetter: "InstanceId"},
		},
		func() interface{} {
			return &jsiiProxy_IInstance{}
		},
	)
	_jsii_.RegisterInterface(
		"cdk-extensions.sso.IPermissionSet",
		reflect.TypeOf((*IPermissionSet)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "permissionSetArn", GoGetter: "PermissionSetArn"},
		},
		func() interface{} {
			return &jsiiProxy_IPermissionSet{}
		},
	)
	_jsii_.RegisterInterface(
		"cdk-extensions.sso.IUser",
		reflect.TypeOf((*IUser)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "userId", GoGetter: "UserId"},
		},
		func() interface{} {
			return &jsiiProxy_IUser{}
		},
	)
	_jsii_.RegisterClass(
		"cdk-extensions.sso.IdentityCenterPrincipalType",
		reflect.TypeOf((*IdentityCenterPrincipalType)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
		},
		func() interface{} {
			return &jsiiProxy_IdentityCenterPrincipalType{}
		},
	)
	_jsii_.RegisterClass(
		"cdk-extensions.sso.Instance",
		reflect.TypeOf((*Instance)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_Instance{}
		},
	)
	_jsii_.RegisterClass(
		"cdk-extensions.sso.InstanceAccessControlAttributeConfiguration",
		reflect.TypeOf((*InstanceAccessControlAttributeConfiguration)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addAttribute", GoMethod: "AddAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "attributes", GoGetter: "Attributes"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "instance", GoGetter: "Instance"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "resource", GoGetter: "Resource"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_InstanceAccessControlAttributeConfiguration{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdk-extensions.sso.InstanceAccessControlAttributeConfigurationProps",
		reflect.TypeOf((*InstanceAccessControlAttributeConfigurationProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk-extensions.sso.InstanceBase",
		reflect.TypeOf((*InstanceBase)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "instanceArn", GoGetter: "InstanceArn"},
			_jsii_.MemberProperty{JsiiProperty: "instanceId", GoGetter: "InstanceId"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_InstanceBase{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkResource)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IInstance)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"cdk-extensions.sso.ManagedPolicyPermissionsBoundary",
		reflect.TypeOf((*ManagedPolicyPermissionsBoundary)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
			_jsii_.MemberProperty{JsiiProperty: "managedPolicy", GoGetter: "ManagedPolicy"},
		},
		func() interface{} {
			j := jsiiProxy_ManagedPolicyPermissionsBoundary{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_PermissionsBoundary)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"cdk-extensions.sso.PermissionSet",
		reflect.TypeOf((*PermissionSet)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addCustomerManagedPolicy", GoMethod: "AddCustomerManagedPolicy"},
			_jsii_.MemberMethod{JsiiMethod: "addManagedPolicy", GoMethod: "AddManagedPolicy"},
			_jsii_.MemberMethod{JsiiMethod: "addToPrincipalPolicy", GoMethod: "AddToPrincipalPolicy"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "description", GoGetter: "Description"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "instance", GoGetter: "Instance"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "permissionsBoundary", GoGetter: "PermissionsBoundary"},
			_jsii_.MemberProperty{JsiiProperty: "permissionSetArn", GoGetter: "PermissionSetArn"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "relayState", GoGetter: "RelayState"},
			_jsii_.MemberProperty{JsiiProperty: "resource", GoGetter: "Resource"},
			_jsii_.MemberProperty{JsiiProperty: "sessionDuration", GoGetter: "SessionDuration"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_PermissionSet{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkResource)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IPermissionSet)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdk-extensions.sso.PermissionSetProps",
		reflect.TypeOf((*PermissionSetProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk-extensions.sso.PermissionsBoundary",
		reflect.TypeOf((*PermissionsBoundary)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
		},
		func() interface{} {
			return &jsiiProxy_PermissionsBoundary{}
		},
	)
	_jsii_.RegisterStruct(
		"cdk-extensions.sso.ReferenceOptions",
		reflect.TypeOf((*ReferenceOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk-extensions.sso.ReferencedManagedPolicy",
		reflect.TypeOf((*ReferencedManagedPolicy)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addStatements", GoMethod: "AddStatements"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberMethod{JsiiMethod: "attachToGroup", GoMethod: "AttachToGroup"},
			_jsii_.MemberMethod{JsiiMethod: "attachToRole", GoMethod: "AttachToRole"},
			_jsii_.MemberMethod{JsiiMethod: "attachToUser", GoMethod: "AttachToUser"},
			_jsii_.MemberProperty{JsiiProperty: "description", GoGetter: "Description"},
			_jsii_.MemberProperty{JsiiProperty: "document", GoGetter: "Document"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "grantPrincipal", GoGetter: "GrantPrincipal"},
			_jsii_.MemberProperty{JsiiProperty: "managedPolicyArn", GoGetter: "ManagedPolicyArn"},
			_jsii_.MemberProperty{JsiiProperty: "managedPolicyName", GoGetter: "ManagedPolicyName"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "path", GoGetter: "Path"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "referencedName", GoGetter: "ReferencedName"},
			_jsii_.MemberProperty{JsiiProperty: "referencedPath", GoGetter: "ReferencedPath"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_ReferencedManagedPolicy{}
			_jsii_.InitJsiiProxy(&j.Type__awsiamManagedPolicy)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdk-extensions.sso.ReferencedManagedPolicyProps",
		reflect.TypeOf((*ReferencedManagedPolicyProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk-extensions.sso.ReferencedPermissionsBoundary",
		reflect.TypeOf((*ReferencedPermissionsBoundary)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
			_jsii_.MemberProperty{JsiiProperty: "referencedPolicy", GoGetter: "ReferencedPolicy"},
		},
		func() interface{} {
			j := jsiiProxy_ReferencedPermissionsBoundary{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_PermissionsBoundary)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"cdk-extensions.sso.User",
		reflect.TypeOf((*User)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_User{}
		},
	)
	_jsii_.RegisterClass(
		"cdk-extensions.sso.UserBase",
		reflect.TypeOf((*UserBase)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "principalId", GoGetter: "PrincipalId"},
			_jsii_.MemberProperty{JsiiProperty: "principalType", GoGetter: "PrincipalType"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "userId", GoGetter: "UserId"},
		},
		func() interface{} {
			j := jsiiProxy_UserBase{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkResource)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IIdentityCenterPrincipal)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IUser)
			return &j
		},
	)
}
