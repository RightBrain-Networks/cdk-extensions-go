package guardduty

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"cdk-extensions.guardduty.DataSource",
		reflect.TypeOf((*DataSource)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_DataSource{}
		},
	)
	_jsii_.RegisterClass(
		"cdk-extensions.guardduty.Detector",
		reflect.TypeOf((*Detector)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addFeature", GoMethod: "AddFeature"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "detectorArn", GoGetter: "DetectorArn"},
			_jsii_.MemberProperty{JsiiProperty: "detectorId", GoGetter: "DetectorId"},
			_jsii_.MemberProperty{JsiiProperty: "enabled", GoGetter: "Enabled"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "features", GoGetter: "Features"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "publishingFrequency", GoGetter: "PublishingFrequency"},
			_jsii_.MemberProperty{JsiiProperty: "resource", GoGetter: "Resource"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_Detector{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkResource)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IDetector)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdk-extensions.guardduty.DetectorAttributes",
		reflect.TypeOf((*DetectorAttributes)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk-extensions.guardduty.DetectorOptions",
		reflect.TypeOf((*DetectorOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk-extensions.guardduty.DetectorProps",
		reflect.TypeOf((*DetectorProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk-extensions.guardduty.EksRuntimeMonitoringOptions",
		reflect.TypeOf((*EksRuntimeMonitoringOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk-extensions.guardduty.Feature",
		reflect.TypeOf((*Feature)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_Feature{}
		},
	)
	_jsii_.RegisterStruct(
		"cdk-extensions.guardduty.FeatureOptions",
		reflect.TypeOf((*FeatureOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk-extensions.guardduty.FindingPublishingFrequency",
		reflect.TypeOf((*FindingPublishingFrequency)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "label", GoGetter: "Label"},
		},
		func() interface{} {
			return &jsiiProxy_FindingPublishingFrequency{}
		},
	)
	_jsii_.RegisterInterface(
		"cdk-extensions.guardduty.IDataSource",
		reflect.TypeOf((*IDataSource)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
		},
		func() interface{} {
			return &jsiiProxy_IDataSource{}
		},
	)
	_jsii_.RegisterInterface(
		"cdk-extensions.guardduty.IDetector",
		reflect.TypeOf((*IDetector)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "detectorArn", GoGetter: "DetectorArn"},
			_jsii_.MemberProperty{JsiiProperty: "detectorId", GoGetter: "DetectorId"},
		},
		func() interface{} {
			return &jsiiProxy_IDetector{}
		},
	)
	_jsii_.RegisterInterface(
		"cdk-extensions.guardduty.IFeature",
		reflect.TypeOf((*IFeature)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "dataSource", GoGetter: "DataSource"},
			_jsii_.MemberProperty{JsiiProperty: "enabled", GoGetter: "Enabled"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "settings", GoGetter: "Settings"},
		},
		func() interface{} {
			return &jsiiProxy_IFeature{}
		},
	)
	_jsii_.RegisterInterface(
		"cdk-extensions.guardduty.IFeatureSetting",
		reflect.TypeOf((*IFeatureSetting)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "enabled", GoGetter: "Enabled"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
		},
		func() interface{} {
			return &jsiiProxy_IFeatureSetting{}
		},
	)
	_jsii_.RegisterStruct(
		"cdk-extensions.guardduty.KubernetesOptions",
		reflect.TypeOf((*KubernetesOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk-extensions.guardduty.MalwareProtectionOptions",
		reflect.TypeOf((*MalwareProtectionOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk-extensions.guardduty.S3LogsOptions",
		reflect.TypeOf((*S3LogsOptions)(nil)).Elem(),
	)
}
