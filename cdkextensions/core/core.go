package core

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"cdk-extensions.core.DataSize",
		reflect.TypeOf((*DataSize)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "toBytes", GoMethod: "ToBytes"},
			_jsii_.MemberMethod{JsiiMethod: "toGibibytes", GoMethod: "ToGibibytes"},
			_jsii_.MemberMethod{JsiiMethod: "toGigabytes", GoMethod: "ToGigabytes"},
			_jsii_.MemberMethod{JsiiMethod: "toKibibytes", GoMethod: "ToKibibytes"},
			_jsii_.MemberMethod{JsiiMethod: "toKilobytes", GoMethod: "ToKilobytes"},
			_jsii_.MemberMethod{JsiiMethod: "toMebibytes", GoMethod: "ToMebibytes"},
			_jsii_.MemberMethod{JsiiMethod: "toMegabytes", GoMethod: "ToMegabytes"},
			_jsii_.MemberMethod{JsiiMethod: "toPebibytes", GoMethod: "ToPebibytes"},
			_jsii_.MemberMethod{JsiiMethod: "toPetabytes", GoMethod: "ToPetabytes"},
			_jsii_.MemberMethod{JsiiMethod: "toTebibytes", GoMethod: "ToTebibytes"},
			_jsii_.MemberMethod{JsiiMethod: "toTerabytes", GoMethod: "ToTerabytes"},
		},
		func() interface{} {
			return &jsiiProxy_DataSize{}
		},
	)
	_jsii_.RegisterClass(
		"cdk-extensions.core.State",
		reflect.TypeOf((*State)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "get", GoMethod: "Get"},
			_jsii_.MemberMethod{JsiiMethod: "set", GoMethod: "Set"},
		},
		func() interface{} {
			return &jsiiProxy_State{}
		},
	)
}
