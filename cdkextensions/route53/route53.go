package route53

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"cdk-extensions.route53.Domain",
		reflect.TypeOf((*Domain)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "certificate", GoGetter: "Certificate"},
			_jsii_.MemberProperty{JsiiProperty: "fqdn", GoGetter: "Fqdn"},
			_jsii_.MemberProperty{JsiiProperty: "isPublic", GoGetter: "IsPublic"},
			_jsii_.MemberProperty{JsiiProperty: "subdomain", GoGetter: "Subdomain"},
			_jsii_.MemberMethod{JsiiMethod: "visit", GoMethod: "Visit"},
			_jsii_.MemberProperty{JsiiProperty: "zone", GoGetter: "Zone"},
			_jsii_.MemberProperty{JsiiProperty: "zoneName", GoGetter: "ZoneName"},
		},
		func() interface{} {
			return &jsiiProxy_Domain{}
		},
	)
	_jsii_.RegisterEnum(
		"cdk-extensions.route53.DomainDiscovery",
		reflect.TypeOf((*DomainDiscovery)(nil)).Elem(),
		map[string]interface{}{
			"ALL": DomainDiscovery_ALL,
			"NONE": DomainDiscovery_NONE,
			"PRIVATE": DomainDiscovery_PRIVATE,
			"PUBLIC": DomainDiscovery_PUBLIC,
		},
	)
	_jsii_.RegisterClass(
		"cdk-extensions.route53.DomainManager",
		reflect.TypeOf((*DomainManager)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_DomainManager{}
		},
	)
	_jsii_.RegisterStruct(
		"cdk-extensions.route53.DomainOptions",
		reflect.TypeOf((*DomainOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk-extensions.route53.Domains",
		reflect.TypeOf((*Domains)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "add", GoMethod: "Add"},
		},
		func() interface{} {
			return &jsiiProxy_Domains{}
		},
	)
	_jsii_.RegisterInterface(
		"cdk-extensions.route53.IDnsResolvable",
		reflect.TypeOf((*IDnsResolvable)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "domainDiscovery", GoGetter: "DomainDiscovery"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "registerDomain", GoMethod: "RegisterDomain"},
		},
		func() interface{} {
			j := jsiiProxy_IDnsResolvable{}
			_jsii_.InitJsiiProxy(&j.Type__constructsIConstruct)
			return &j
		},
	)
}
