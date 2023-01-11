package rds

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"cdk-extensions.rds.DatabaseProxyEndpoint",
		reflect.TypeOf((*DatabaseProxyEndpoint)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "access", GoGetter: "Access"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "connections", GoGetter: "Connections"},
			_jsii_.MemberProperty{JsiiProperty: "databaseProxy", GoGetter: "DatabaseProxy"},
			_jsii_.MemberProperty{JsiiProperty: "databaseProxyEndpointArn", GoGetter: "DatabaseProxyEndpointArn"},
			_jsii_.MemberProperty{JsiiProperty: "databaseProxyEndpointHost", GoGetter: "DatabaseProxyEndpointHost"},
			_jsii_.MemberProperty{JsiiProperty: "databaseProxyEndpointIsDefault", GoGetter: "DatabaseProxyEndpointIsDefault"},
			_jsii_.MemberProperty{JsiiProperty: "databaseProxyEndpointVpcId", GoGetter: "DatabaseProxyEndpointVpcId"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "resource", GoGetter: "Resource"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "vpc", GoGetter: "Vpc"},
			_jsii_.MemberProperty{JsiiProperty: "vpcSubnets", GoGetter: "VpcSubnets"},
		},
		func() interface{} {
			j := jsiiProxy_DatabaseProxyEndpoint{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkResource)
			_jsii_.InitJsiiProxy(&j.Type__awsec2IConnectable)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"cdk-extensions.rds.DatabaseProxyEndpointAccess",
		reflect.TypeOf((*DatabaseProxyEndpointAccess)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "role", GoGetter: "Role"},
		},
		func() interface{} {
			return &jsiiProxy_DatabaseProxyEndpointAccess{}
		},
	)
	_jsii_.RegisterStruct(
		"cdk-extensions.rds.DatabaseProxyEndpointProps",
		reflect.TypeOf((*DatabaseProxyEndpointProps)(nil)).Elem(),
	)
}
