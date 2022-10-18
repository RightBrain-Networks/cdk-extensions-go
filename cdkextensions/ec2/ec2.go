package ec2

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"cdk-extensions.ec2.FlowLog",
		reflect.TypeOf((*FlowLog)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "destination", GoGetter: "Destination"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "flowLogArn", GoGetter: "FlowLogArn"},
			_jsii_.MemberProperty{JsiiProperty: "flowLogId", GoGetter: "FlowLogId"},
			_jsii_.MemberProperty{JsiiProperty: "format", GoGetter: "Format"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "maxAggregationInterval", GoGetter: "MaxAggregationInterval"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "resource", GoGetter: "Resource"},
			_jsii_.MemberProperty{JsiiProperty: "resourceType", GoGetter: "ResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "trafficType", GoGetter: "TrafficType"},
		},
		func() interface{} {
			j := jsiiProxy_FlowLog{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkResource)
			_jsii_.InitJsiiProxy(&j.Type__awsec2IFlowLog)
			return &j
		},
	)
	_jsii_.RegisterEnum(
		"cdk-extensions.ec2.FlowLogAggregationInterval",
		reflect.TypeOf((*FlowLogAggregationInterval)(nil)).Elem(),
		map[string]interface{}{
			"ONE_MINUTE": FlowLogAggregationInterval_ONE_MINUTE,
			"TEN_MINUTES": FlowLogAggregationInterval_TEN_MINUTES,
		},
	)
	_jsii_.RegisterEnum(
		"cdk-extensions.ec2.FlowLogDataType",
		reflect.TypeOf((*FlowLogDataType)(nil)).Elem(),
		map[string]interface{}{
			"INT_32": FlowLogDataType_INT_32,
			"INT_64": FlowLogDataType_INT_64,
			"STRING": FlowLogDataType_STRING,
		},
	)
	_jsii_.RegisterClass(
		"cdk-extensions.ec2.FlowLogDestination",
		reflect.TypeOf((*FlowLogDestination)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
		},
		func() interface{} {
			j := jsiiProxy_FlowLogDestination{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_ILogDestination)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdk-extensions.ec2.FlowLogDestinationConfig",
		reflect.TypeOf((*FlowLogDestinationConfig)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk-extensions.ec2.FlowLogField",
		reflect.TypeOf((*FlowLogField)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "type", GoGetter: "Type"},
		},
		func() interface{} {
			return &jsiiProxy_FlowLogField{}
		},
	)
	_jsii_.RegisterEnum(
		"cdk-extensions.ec2.FlowLogFileFormat",
		reflect.TypeOf((*FlowLogFileFormat)(nil)).Elem(),
		map[string]interface{}{
			"PARQUET": FlowLogFileFormat_PARQUET,
			"PLAIN_TEXT": FlowLogFileFormat_PLAIN_TEXT,
		},
	)
	_jsii_.RegisterClass(
		"cdk-extensions.ec2.FlowLogFormat",
		reflect.TypeOf((*FlowLogFormat)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addField", GoMethod: "AddField"},
			_jsii_.MemberProperty{JsiiProperty: "fields", GoGetter: "Fields"},
			_jsii_.MemberProperty{JsiiProperty: "template", GoGetter: "Template"},
		},
		func() interface{} {
			return &jsiiProxy_FlowLogFormat{}
		},
	)
	_jsii_.RegisterStruct(
		"cdk-extensions.ec2.FlowLogProps",
		reflect.TypeOf((*FlowLogProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk-extensions.ec2.FlowLogS3Options",
		reflect.TypeOf((*FlowLogS3Options)(nil)).Elem(),
	)
	_jsii_.RegisterInterface(
		"cdk-extensions.ec2.ILogDestination",
		reflect.TypeOf((*ILogDestination)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
		},
		func() interface{} {
			return &jsiiProxy_ILogDestination{}
		},
	)
}
