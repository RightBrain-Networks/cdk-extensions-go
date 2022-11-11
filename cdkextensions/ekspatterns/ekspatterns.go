package ekspatterns

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"cdk-extensions.eks_patterns.AwsIntegratedFargateCluster",
		reflect.TypeOf((*AwsIntegratedFargateCluster)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "adotCollector", GoGetter: "AdotCollector"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "externalSecrets", GoGetter: "ExternalSecrets"},
			_jsii_.MemberProperty{JsiiProperty: "fargateLogger", GoGetter: "FargateLogger"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "registerSecretsManagerSecret", GoMethod: "RegisterSecretsManagerSecret"},
			_jsii_.MemberMethod{JsiiMethod: "registerSsmParameterSecret", GoMethod: "RegisterSsmParameterSecret"},
			_jsii_.MemberProperty{JsiiProperty: "resource", GoGetter: "Resource"},
			_jsii_.MemberProperty{JsiiProperty: "route53Dns", GoGetter: "Route53Dns"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_AwsIntegratedFargateCluster{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdk-extensions.eks_patterns.AwsIntegratedFargateClusterProps",
		reflect.TypeOf((*AwsIntegratedFargateClusterProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk-extensions.eks_patterns.ClusterFargateLoggingOptions",
		reflect.TypeOf((*ClusterFargateLoggingOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk-extensions.eks_patterns.ClusterRoute53DnsOptions",
		reflect.TypeOf((*ClusterRoute53DnsOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk-extensions.eks_patterns.ContainerInsightsOptions",
		reflect.TypeOf((*ContainerInsightsOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk-extensions.eks_patterns.ExternalSecretsOptions",
		reflect.TypeOf((*ExternalSecretsOptions)(nil)).Elem(),
	)
}
