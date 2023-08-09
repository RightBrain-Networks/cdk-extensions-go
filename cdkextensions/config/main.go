package config

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterStruct(
		"cdk-extensions.config.AutomationDocumentRemediationProps",
		reflect.TypeOf((*AutomationDocumentRemediationProps)(nil)).Elem(),
	)
	_jsii_.RegisterInterface(
		"cdk-extensions.config.IRemediationConfiguration",
		reflect.TypeOf((*IRemediationConfiguration)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "remediationConfigurationArn", GoGetter: "RemediationConfigurationArn"},
			_jsii_.MemberProperty{JsiiProperty: "remediationConfigurationName", GoGetter: "RemediationConfigurationName"},
		},
		func() interface{} {
			return &jsiiProxy_IRemediationConfiguration{}
		},
	)
	_jsii_.RegisterInterface(
		"cdk-extensions.config.IRemediationTarget",
		reflect.TypeOf((*IRemediationTarget)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
		},
		func() interface{} {
			return &jsiiProxy_IRemediationTarget{}
		},
	)
	_jsii_.RegisterClass(
		"cdk-extensions.config.RemediationConfiguration",
		reflect.TypeOf((*RemediationConfiguration)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addParameter", GoMethod: "AddParameter"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "automatic", GoGetter: "Automatic"},
			_jsii_.MemberProperty{JsiiProperty: "configRule", GoGetter: "ConfigRule"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "maximumAutomaticAttempts", GoGetter: "MaximumAutomaticAttempts"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "remediationConfigurationArn", GoGetter: "RemediationConfigurationArn"},
			_jsii_.MemberProperty{JsiiProperty: "remediationConfigurationName", GoGetter: "RemediationConfigurationName"},
			_jsii_.MemberMethod{JsiiMethod: "renderParameters", GoMethod: "RenderParameters"},
			_jsii_.MemberProperty{JsiiProperty: "resource", GoGetter: "Resource"},
			_jsii_.MemberProperty{JsiiProperty: "resourceType", GoGetter: "ResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "retryInterval", GoGetter: "RetryInterval"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_RemediationConfiguration{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkResource)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IRemediationConfiguration)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdk-extensions.config.RemediationConfigurationAttributes",
		reflect.TypeOf((*RemediationConfigurationAttributes)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk-extensions.config.RemediationConfigurationProps",
		reflect.TypeOf((*RemediationConfigurationProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk-extensions.config.RemediationTarget",
		reflect.TypeOf((*RemediationTarget)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_RemediationTarget{}
		},
	)
	_jsii_.RegisterStruct(
		"cdk-extensions.config.RemediationTargetConfiguration",
		reflect.TypeOf((*RemediationTargetConfiguration)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk-extensions.config.RemediationTargetType",
		reflect.TypeOf((*RemediationTargetType)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_RemediationTargetType{}
		},
	)
}
