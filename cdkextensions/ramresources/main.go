package ramresources

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"cdk-extensions.ram_resources.SharedResource",
		reflect.TypeOf((*SharedResource)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "share", GoMethod: "Share"},
		},
		func() interface{} {
			j := jsiiProxy_SharedResource{}
			_jsii_.InitJsiiProxy(&j.Type__ramISharable)
			return &j
		},
	)
}
