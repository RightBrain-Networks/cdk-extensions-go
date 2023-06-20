package stepfunctions

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"cdk-extensions.stepfunctions.SfnFn",
		reflect.TypeOf((*SfnFn)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_SfnFn{}
		},
	)
	_jsii_.RegisterClass(
		"cdk-extensions.stepfunctions.StepFunctionValidation",
		reflect.TypeOf((*StepFunctionValidation)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_StepFunctionValidation{}
		},
	)
	_jsii_.RegisterClass(
		"cdk-extensions.stepfunctions.StringReplace",
		reflect.TypeOf((*StringReplace)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "endStates", GoGetter: "EndStates"},
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberMethod{JsiiMethod: "next", GoMethod: "Next"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "startState", GoGetter: "StartState"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_StringReplace{}
			_jsii_.InitJsiiProxy(&j.Type__constructsConstruct)
			_jsii_.InitJsiiProxy(&j.Type__awsstepfunctionsIChainable)
			_jsii_.InitJsiiProxy(&j.Type__awsstepfunctionsINextable)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdk-extensions.stepfunctions.StringReplaceProps",
		reflect.TypeOf((*StringReplaceProps)(nil)).Elem(),
	)
}
