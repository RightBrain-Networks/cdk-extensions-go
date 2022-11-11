package asserts

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"cdk-extensions.asserts.JoinedJson",
		reflect.TypeOf((*JoinedJson)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberMethod{JsiiMethod: "test", GoMethod: "Test"},
		},
		func() interface{} {
			j := jsiiProxy_JoinedJson{}
			_jsii_.InitJsiiProxy(&j.Type__assertionsMatcher)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"cdk-extensions.asserts.Match",
		reflect.TypeOf((*Match)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			j := jsiiProxy_Match{}
			_jsii_.InitJsiiProxy(&j.Type__assertionsMatch)
			return &j
		},
	)
}
