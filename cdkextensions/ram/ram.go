package ram

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterInterface(
		"cdk-extensions.ram.ISharedPrincipal",
		reflect.TypeOf((*ISharedPrincipal)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
		},
		func() interface{} {
			return &jsiiProxy_ISharedPrincipal{}
		},
	)
	_jsii_.RegisterInterface(
		"cdk-extensions.ram.ISharedResource",
		reflect.TypeOf((*ISharedResource)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
		},
		func() interface{} {
			return &jsiiProxy_ISharedResource{}
		},
	)
	_jsii_.RegisterClass(
		"cdk-extensions.ram.ResourceShare",
		reflect.TypeOf((*ResourceShare)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addPrincipal", GoMethod: "AddPrincipal"},
			_jsii_.MemberMethod{JsiiMethod: "addResource", GoMethod: "AddResource"},
			_jsii_.MemberProperty{JsiiProperty: "allowExternalPrincipals", GoGetter: "AllowExternalPrincipals"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "autoDiscovery", GoGetter: "AutoDiscovery"},
			_jsii_.MemberMethod{JsiiMethod: "enableAutoDiscovery", GoMethod: "EnableAutoDiscovery"},
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
		},
		func() interface{} {
			j := jsiiProxy_ResourceShare{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdk-extensions.ram.ResourceShareProps",
		reflect.TypeOf((*ResourceShareProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk-extensions.ram.SharedPrincipal",
		reflect.TypeOf((*SharedPrincipal)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
		},
		func() interface{} {
			j := jsiiProxy_SharedPrincipal{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_ISharedPrincipal)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"cdk-extensions.ram.SharedResource",
		reflect.TypeOf((*SharedResource)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
		},
		func() interface{} {
			j := jsiiProxy_SharedResource{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_ISharedResource)
			return &j
		},
	)
}
