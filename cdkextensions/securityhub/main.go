package securityhub

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"cdk-extensions.securityhub.ControlFindingGenerator",
		reflect.TypeOf((*ControlFindingGenerator)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_ControlFindingGenerator{}
		},
	)
	_jsii_.RegisterStruct(
		"cdk-extensions.securityhub.DisableControlOptions",
		reflect.TypeOf((*DisableControlOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk-extensions.securityhub.Hub",
		reflect.TypeOf((*Hub)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "autoEnableControls", GoGetter: "AutoEnableControls"},
			_jsii_.MemberProperty{JsiiProperty: "consolidatedFindings", GoGetter: "ConsolidatedFindings"},
			_jsii_.MemberProperty{JsiiProperty: "controlFindingGenerator", GoGetter: "ControlFindingGenerator"},
			_jsii_.MemberProperty{JsiiProperty: "enableDefaultStandards", GoGetter: "EnableDefaultStandards"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "hubArn", GoGetter: "HubArn"},
			_jsii_.MemberProperty{JsiiProperty: "hubName", GoGetter: "HubName"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "resource", GoGetter: "Resource"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_Hub{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkResource)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IHub)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdk-extensions.securityhub.HubAttributes",
		reflect.TypeOf((*HubAttributes)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk-extensions.securityhub.HubProps",
		reflect.TypeOf((*HubProps)(nil)).Elem(),
	)
	_jsii_.RegisterInterface(
		"cdk-extensions.securityhub.IHub",
		reflect.TypeOf((*IHub)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "hubArn", GoGetter: "HubArn"},
			_jsii_.MemberProperty{JsiiProperty: "hubName", GoGetter: "HubName"},
		},
		func() interface{} {
			return &jsiiProxy_IHub{}
		},
	)
	_jsii_.RegisterInterface(
		"cdk-extensions.securityhub.IStandard",
		reflect.TypeOf((*IStandard)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "standardArn", GoGetter: "StandardArn"},
		},
		func() interface{} {
			j := jsiiProxy_IStandard{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkIResource)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"cdk-extensions.securityhub.RuleSet",
		reflect.TypeOf((*RuleSet)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
		},
		func() interface{} {
			return &jsiiProxy_RuleSet{}
		},
	)
	_jsii_.RegisterStruct(
		"cdk-extensions.securityhub.RuleSetProps",
		reflect.TypeOf((*RuleSetProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk-extensions.securityhub.ScopedRuleSet",
		reflect.TypeOf((*ScopedRuleSet)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk-extensions.securityhub.Standard",
		reflect.TypeOf((*Standard)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberMethod{JsiiMethod: "disableControl", GoMethod: "DisableControl"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "resource", GoGetter: "Resource"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "standardArn", GoGetter: "StandardArn"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_Standard{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkResource)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IStandard)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdk-extensions.securityhub.StandardProps",
		reflect.TypeOf((*StandardProps)(nil)).Elem(),
	)
}
