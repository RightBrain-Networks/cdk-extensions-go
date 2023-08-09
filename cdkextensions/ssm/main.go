package ssm

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"cdk-extensions.ssm.AutomationDocument",
		reflect.TypeOf((*AutomationDocument)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addRequirement", GoMethod: "AddRequirement"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberMethod{JsiiMethod: "arnForAutomationDefinitionVersion", GoMethod: "ArnForAutomationDefinitionVersion"},
			_jsii_.MemberProperty{JsiiProperty: "automationDefinitionArn", GoGetter: "AutomationDefinitionArn"},
			_jsii_.MemberProperty{JsiiProperty: "content", GoGetter: "Content"},
			_jsii_.MemberProperty{JsiiProperty: "documentArn", GoGetter: "DocumentArn"},
			_jsii_.MemberProperty{JsiiProperty: "documentName", GoGetter: "DocumentName"},
			_jsii_.MemberProperty{JsiiProperty: "documentType", GoGetter: "DocumentType"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "requires", GoGetter: "Requires"},
			_jsii_.MemberProperty{JsiiProperty: "resource", GoGetter: "Resource"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "targetType", GoGetter: "TargetType"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "updateMethod", GoGetter: "UpdateMethod"},
			_jsii_.MemberMethod{JsiiMethod: "validate", GoMethod: "Validate"},
			_jsii_.MemberProperty{JsiiProperty: "versionName", GoGetter: "VersionName"},
		},
		func() interface{} {
			j := jsiiProxy_AutomationDocument{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_DocumentBase)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IAutomationDocument)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"cdk-extensions.ssm.AutomationSchemaVersion",
		reflect.TypeOf((*AutomationSchemaVersion)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "version", GoGetter: "Version"},
		},
		func() interface{} {
			return &jsiiProxy_AutomationSchemaVersion{}
		},
	)
	_jsii_.RegisterClass(
		"cdk-extensions.ssm.Document",
		reflect.TypeOf((*Document)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addRequirement", GoMethod: "AddRequirement"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "content", GoGetter: "Content"},
			_jsii_.MemberProperty{JsiiProperty: "documentArn", GoGetter: "DocumentArn"},
			_jsii_.MemberProperty{JsiiProperty: "documentName", GoGetter: "DocumentName"},
			_jsii_.MemberProperty{JsiiProperty: "documentType", GoGetter: "DocumentType"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "requires", GoGetter: "Requires"},
			_jsii_.MemberProperty{JsiiProperty: "resource", GoGetter: "Resource"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "targetType", GoGetter: "TargetType"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "updateMethod", GoGetter: "UpdateMethod"},
			_jsii_.MemberMethod{JsiiMethod: "validate", GoMethod: "Validate"},
			_jsii_.MemberProperty{JsiiProperty: "versionName", GoGetter: "VersionName"},
		},
		func() interface{} {
			j := jsiiProxy_Document{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_DocumentBase)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"cdk-extensions.ssm.DocumentBase",
		reflect.TypeOf((*DocumentBase)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addRequirement", GoMethod: "AddRequirement"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "content", GoGetter: "Content"},
			_jsii_.MemberProperty{JsiiProperty: "documentArn", GoGetter: "DocumentArn"},
			_jsii_.MemberProperty{JsiiProperty: "documentName", GoGetter: "DocumentName"},
			_jsii_.MemberProperty{JsiiProperty: "documentType", GoGetter: "DocumentType"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "requires", GoGetter: "Requires"},
			_jsii_.MemberProperty{JsiiProperty: "resource", GoGetter: "Resource"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "targetType", GoGetter: "TargetType"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "updateMethod", GoGetter: "UpdateMethod"},
			_jsii_.MemberMethod{JsiiMethod: "validate", GoMethod: "Validate"},
			_jsii_.MemberProperty{JsiiProperty: "versionName", GoGetter: "VersionName"},
		},
		func() interface{} {
			j := jsiiProxy_DocumentBase{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkResource)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IDocument)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdk-extensions.ssm.DocumentBaseProps",
		reflect.TypeOf((*DocumentBaseProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk-extensions.ssm.DocumentContent",
		reflect.TypeOf((*DocumentContent)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_DocumentContent{}
		},
	)
	_jsii_.RegisterStruct(
		"cdk-extensions.ssm.DocumentContentResult",
		reflect.TypeOf((*DocumentContentResult)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk-extensions.ssm.DocumentFormat",
		reflect.TypeOf((*DocumentFormat)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_DocumentFormat{}
		},
	)
	_jsii_.RegisterStruct(
		"cdk-extensions.ssm.DocumentProps",
		reflect.TypeOf((*DocumentProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk-extensions.ssm.DocumentRequirement",
		reflect.TypeOf((*DocumentRequirement)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk-extensions.ssm.DocumentType",
		reflect.TypeOf((*DocumentType)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
		},
		func() interface{} {
			return &jsiiProxy_DocumentType{}
		},
	)
	_jsii_.RegisterClass(
		"cdk-extensions.ssm.DocumentUpdateMethod",
		reflect.TypeOf((*DocumentUpdateMethod)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_DocumentUpdateMethod{}
		},
	)
	_jsii_.RegisterInterface(
		"cdk-extensions.ssm.IAutomationDocument",
		reflect.TypeOf((*IAutomationDocument)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberMethod{JsiiMethod: "arnForAutomationDefinitionVersion", GoMethod: "ArnForAutomationDefinitionVersion"},
			_jsii_.MemberProperty{JsiiProperty: "automationDefinitionArn", GoGetter: "AutomationDefinitionArn"},
			_jsii_.MemberProperty{JsiiProperty: "documentArn", GoGetter: "DocumentArn"},
			_jsii_.MemberProperty{JsiiProperty: "documentName", GoGetter: "DocumentName"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
		},
		func() interface{} {
			j := jsiiProxy_IAutomationDocument{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IDocument)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"cdk-extensions.ssm.IDocument",
		reflect.TypeOf((*IDocument)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "documentArn", GoGetter: "DocumentArn"},
			_jsii_.MemberProperty{JsiiProperty: "documentName", GoGetter: "DocumentName"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
		},
		func() interface{} {
			j := jsiiProxy_IDocument{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkIResource)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"cdk-extensions.ssm.IDocumentContent",
		reflect.TypeOf((*IDocumentContent)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
		},
		func() interface{} {
			return &jsiiProxy_IDocumentContent{}
		},
	)
	_jsii_.RegisterStruct(
		"cdk-extensions.ssm.ObjectContentProps",
		reflect.TypeOf((*ObjectContentProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk-extensions.ssm.StringContentProps",
		reflect.TypeOf((*StringContentProps)(nil)).Elem(),
	)
}
