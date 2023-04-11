package networkmanager

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"cdk-extensions.networkmanager.GlobalNetwork",
		reflect.TypeOf((*GlobalNetwork)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "description", GoGetter: "Description"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "globalNetworkArn", GoGetter: "GlobalNetworkArn"},
			_jsii_.MemberProperty{JsiiProperty: "globalNetworkId", GoGetter: "GlobalNetworkId"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "registerTransitGateway", GoMethod: "RegisterTransitGateway"},
			_jsii_.MemberProperty{JsiiProperty: "resource", GoGetter: "Resource"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_GlobalNetwork{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkResource)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IGlobalNetwork)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdk-extensions.networkmanager.GlobalNetworkProps",
		reflect.TypeOf((*GlobalNetworkProps)(nil)).Elem(),
	)
	_jsii_.RegisterInterface(
		"cdk-extensions.networkmanager.IGlobalNetwork",
		reflect.TypeOf((*IGlobalNetwork)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "globalNetworkArn", GoGetter: "GlobalNetworkArn"},
			_jsii_.MemberProperty{JsiiProperty: "globalNetworkId", GoGetter: "GlobalNetworkId"},
		},
		func() interface{} {
			return &jsiiProxy_IGlobalNetwork{}
		},
	)
	_jsii_.RegisterInterface(
		"cdk-extensions.networkmanager.ISite",
		reflect.TypeOf((*ISite)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "siteArn", GoGetter: "SiteArn"},
			_jsii_.MemberProperty{JsiiProperty: "siteId", GoGetter: "SiteId"},
		},
		func() interface{} {
			return &jsiiProxy_ISite{}
		},
	)
	_jsii_.RegisterInterface(
		"cdk-extensions.networkmanager.ITransitGatewayRegistration",
		reflect.TypeOf((*ITransitGatewayRegistration)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_ITransitGatewayRegistration{}
		},
	)
	_jsii_.RegisterStruct(
		"cdk-extensions.networkmanager.RegisterTransitGatewayProps",
		reflect.TypeOf((*RegisterTransitGatewayProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk-extensions.networkmanager.Site",
		reflect.TypeOf((*Site)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "address", GoGetter: "Address"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "description", GoGetter: "Description"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "globalNetwork", GoGetter: "GlobalNetwork"},
			_jsii_.MemberProperty{JsiiProperty: "latitude", GoGetter: "Latitude"},
			_jsii_.MemberProperty{JsiiProperty: "longitude", GoGetter: "Longitude"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "resource", GoGetter: "Resource"},
			_jsii_.MemberProperty{JsiiProperty: "siteArn", GoGetter: "SiteArn"},
			_jsii_.MemberProperty{JsiiProperty: "siteId", GoGetter: "SiteId"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_Site{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkResource)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_ISite)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdk-extensions.networkmanager.SiteProps",
		reflect.TypeOf((*SiteProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk-extensions.networkmanager.TransitGatewayRegistration",
		reflect.TypeOf((*TransitGatewayRegistration)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "globalNetwork", GoGetter: "GlobalNetwork"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "resource", GoGetter: "Resource"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "transitGateway", GoGetter: "TransitGateway"},
		},
		func() interface{} {
			j := jsiiProxy_TransitGatewayRegistration{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkResource)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_ITransitGatewayRegistration)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdk-extensions.networkmanager.TransitGatewayRegistrationProps",
		reflect.TypeOf((*TransitGatewayRegistrationProps)(nil)).Elem(),
	)
}
