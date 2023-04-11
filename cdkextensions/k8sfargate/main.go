package k8sfargate

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"cdk-extensions.k8s_fargate.Prometheus",
		reflect.TypeOf((*Prometheus)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "cluster", GoGetter: "Cluster"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "fargateProfile", GoGetter: "FargateProfile"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "namespace", GoGetter: "Namespace"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "queueConfiguration", GoGetter: "QueueConfiguration"},
			_jsii_.MemberProperty{JsiiProperty: "serviceAccount", GoGetter: "ServiceAccount"},
			_jsii_.MemberProperty{JsiiProperty: "serviceAccountName", GoGetter: "ServiceAccountName"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "workspace", GoGetter: "Workspace"},
		},
		func() interface{} {
			j := jsiiProxy_Prometheus{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdk-extensions.k8s_fargate.PrometheusOptions",
		reflect.TypeOf((*PrometheusOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk-extensions.k8s_fargate.PrometheusProps",
		reflect.TypeOf((*PrometheusProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk-extensions.k8s_fargate.QueueConfiguration",
		reflect.TypeOf((*QueueConfiguration)(nil)).Elem(),
	)
}
