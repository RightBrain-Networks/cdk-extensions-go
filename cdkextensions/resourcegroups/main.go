package resourcegroups

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterStruct(
		"cdk-extensions.resourcegroups.BoundGroupConfiguration",
		reflect.TypeOf((*BoundGroupConfiguration)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk-extensions.resourcegroups.CloudFormationStack",
		reflect.TypeOf((*CloudFormationStack)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addResourceType", GoMethod: "AddResourceType"},
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
			_jsii_.MemberProperty{JsiiProperty: "reference", GoGetter: "Reference"},
			_jsii_.MemberProperty{JsiiProperty: "resourceTypes", GoGetter: "ResourceTypes"},
		},
		func() interface{} {
			j := jsiiProxy_CloudFormationStack{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IGroupConfiguration)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdk-extensions.resourcegroups.CloudFormationStackProps",
		reflect.TypeOf((*CloudFormationStackProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk-extensions.resourcegroups.Group",
		reflect.TypeOf((*Group)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "description", GoGetter: "Description"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "groupArn", GoGetter: "GroupArn"},
			_jsii_.MemberProperty{JsiiProperty: "groupName", GoGetter: "GroupName"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "resource", GoGetter: "Resource"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_Group{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkResource)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IGroup)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdk-extensions.resourcegroups.GroupAttributes",
		reflect.TypeOf((*GroupAttributes)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk-extensions.resourcegroups.GroupConfiguration",
		reflect.TypeOf((*GroupConfiguration)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_GroupConfiguration{}
		},
	)
	_jsii_.RegisterStruct(
		"cdk-extensions.resourcegroups.GroupProps",
		reflect.TypeOf((*GroupProps)(nil)).Elem(),
	)
	_jsii_.RegisterInterface(
		"cdk-extensions.resourcegroups.IGroup",
		reflect.TypeOf((*IGroup)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "groupArn", GoGetter: "GroupArn"},
			_jsii_.MemberProperty{JsiiProperty: "groupName", GoGetter: "GroupName"},
		},
		func() interface{} {
			return &jsiiProxy_IGroup{}
		},
	)
	_jsii_.RegisterInterface(
		"cdk-extensions.resourcegroups.IGroupConfiguration",
		reflect.TypeOf((*IGroupConfiguration)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
		},
		func() interface{} {
			return &jsiiProxy_IGroupConfiguration{}
		},
	)
	_jsii_.RegisterInterface(
		"cdk-extensions.resourcegroups.IStackReference",
		reflect.TypeOf((*IStackReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "stackConstruct", GoGetter: "StackConstruct"},
			_jsii_.MemberProperty{JsiiProperty: "stackId", GoGetter: "StackId"},
		},
		func() interface{} {
			return &jsiiProxy_IStackReference{}
		},
	)
	_jsii_.RegisterClass(
		"cdk-extensions.resourcegroups.StackReference",
		reflect.TypeOf((*StackReference)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_StackReference{}
		},
	)
	_jsii_.RegisterClass(
		"cdk-extensions.resourcegroups.TagFilter",
		reflect.TypeOf((*TagFilter)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addResourceType", GoMethod: "AddResourceType"},
			_jsii_.MemberMethod{JsiiMethod: "addTagFilter", GoMethod: "AddTagFilter"},
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
			_jsii_.MemberProperty{JsiiProperty: "resourceTypes", GoGetter: "ResourceTypes"},
			_jsii_.MemberProperty{JsiiProperty: "tagFilters", GoGetter: "TagFilters"},
		},
		func() interface{} {
			j := jsiiProxy_TagFilter{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IGroupConfiguration)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"cdk-extensions.resourcegroups.TagFilterProps",
		reflect.TypeOf((*TagFilterProps)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "resourceTypes", GoGetter: "ResourceTypes"},
			_jsii_.MemberProperty{JsiiProperty: "tagFilters", GoGetter: "TagFilters"},
		},
		func() interface{} {
			return &jsiiProxy_TagFilterProps{}
		},
	)
}
