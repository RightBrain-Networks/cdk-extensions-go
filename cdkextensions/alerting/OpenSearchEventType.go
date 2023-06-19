package alerting

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/vibe-io/cdk-extensions-go/cdkextensions/jsii"
)

// Represents a type of event that can be generated in response to circumstances happening on an AWS OpenSearch service cluster.
// See: [Monitoring OpenSearch Service events with Amazon EventBridge](https://docs.aws.amazon.com/opensearch-service/latest/developerguide/monitoring-events.html)
//
type OpenSearchEventType interface {
	DetailType() *string
	EventName() *string
}

// The jsii proxy struct for OpenSearchEventType
type jsiiProxy_OpenSearchEventType struct {
	_ byte // padding
}

func (j *jsiiProxy_OpenSearchEventType) DetailType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"detailType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpenSearchEventType) EventName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"eventName",
		&returns,
	)
	return returns
}


func OpenSearchEventType_Of(props *OpenSearchEventTypeProps) OpenSearchEventType {
	_init_.Initialize()

	if err := validateOpenSearchEventType_OfParameters(props); err != nil {
		panic(err)
	}
	var returns OpenSearchEventType

	_jsii_.StaticInvoke(
		"cdk-extensions.alerting.OpenSearchEventType",
		"of",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func OpenSearchEventType_All() *[]OpenSearchEventType {
	_init_.Initialize()
	var returns *[]OpenSearchEventType
	_jsii_.StaticGet(
		"cdk-extensions.alerting.OpenSearchEventType",
		"all",
		&returns,
	)
	return returns
}

func OpenSearchEventType_AUTO_TUNE() OpenSearchEventType {
	_init_.Initialize()
	var returns OpenSearchEventType
	_jsii_.StaticGet(
		"cdk-extensions.alerting.OpenSearchEventType",
		"AUTO_TUNE",
		&returns,
	)
	return returns
}

func OpenSearchEventType_CLUSTER_RECOVERY() OpenSearchEventType {
	_init_.Initialize()
	var returns OpenSearchEventType
	_jsii_.StaticGet(
		"cdk-extensions.alerting.OpenSearchEventType",
		"CLUSTER_RECOVERY",
		&returns,
	)
	return returns
}

func OpenSearchEventType_CUSTOM_INDEX_ROUTING() OpenSearchEventType {
	_init_.Initialize()
	var returns OpenSearchEventType
	_jsii_.StaticGet(
		"cdk-extensions.alerting.OpenSearchEventType",
		"CUSTOM_INDEX_ROUTING",
		&returns,
	)
	return returns
}

func OpenSearchEventType_DISK_THROUGHPUT_THROTTLE() OpenSearchEventType {
	_init_.Initialize()
	var returns OpenSearchEventType
	_jsii_.StaticGet(
		"cdk-extensions.alerting.OpenSearchEventType",
		"DISK_THROUGHPUT_THROTTLE",
		&returns,
	)
	return returns
}

func OpenSearchEventType_DOMAIN_UPDATE() OpenSearchEventType {
	_init_.Initialize()
	var returns OpenSearchEventType
	_jsii_.StaticGet(
		"cdk-extensions.alerting.OpenSearchEventType",
		"DOMAIN_UPDATE",
		&returns,
	)
	return returns
}

func OpenSearchEventType_EBS_BURST_BALANCE() OpenSearchEventType {
	_init_.Initialize()
	var returns OpenSearchEventType
	_jsii_.StaticGet(
		"cdk-extensions.alerting.OpenSearchEventType",
		"EBS_BURST_BALANCE",
		&returns,
	)
	return returns
}

func OpenSearchEventType_FAILED_SHARD_LOCK() OpenSearchEventType {
	_init_.Initialize()
	var returns OpenSearchEventType
	_jsii_.StaticGet(
		"cdk-extensions.alerting.OpenSearchEventType",
		"FAILED_SHARD_LOCK",
		&returns,
	)
	return returns
}

func OpenSearchEventType_HIGH_JVM_USAGE() OpenSearchEventType {
	_init_.Initialize()
	var returns OpenSearchEventType
	_jsii_.StaticGet(
		"cdk-extensions.alerting.OpenSearchEventType",
		"HIGH_JVM_USAGE",
		&returns,
	)
	return returns
}

func OpenSearchEventType_HIGH_SHARED_COUNT() OpenSearchEventType {
	_init_.Initialize()
	var returns OpenSearchEventType
	_jsii_.StaticGet(
		"cdk-extensions.alerting.OpenSearchEventType",
		"HIGH_SHARED_COUNT",
		&returns,
	)
	return returns
}

func OpenSearchEventType_INSUFFICIENT_GARBAGE_COLLECTION() OpenSearchEventType {
	_init_.Initialize()
	var returns OpenSearchEventType
	_jsii_.StaticGet(
		"cdk-extensions.alerting.OpenSearchEventType",
		"INSUFFICIENT_GARBAGE_COLLECTION",
		&returns,
	)
	return returns
}

func OpenSearchEventType_KMS_KEY_INACCESSIBLE() OpenSearchEventType {
	_init_.Initialize()
	var returns OpenSearchEventType
	_jsii_.StaticGet(
		"cdk-extensions.alerting.OpenSearchEventType",
		"KMS_KEY_INACCESSIBLE",
		&returns,
	)
	return returns
}

func OpenSearchEventType_LARGE_SHARD_SIZE() OpenSearchEventType {
	_init_.Initialize()
	var returns OpenSearchEventType
	_jsii_.StaticGet(
		"cdk-extensions.alerting.OpenSearchEventType",
		"LARGE_SHARD_SIZE",
		&returns,
	)
	return returns
}

func OpenSearchEventType_LOW_DISK_SPACE() OpenSearchEventType {
	_init_.Initialize()
	var returns OpenSearchEventType
	_jsii_.StaticGet(
		"cdk-extensions.alerting.OpenSearchEventType",
		"LOW_DISK_SPACE",
		&returns,
	)
	return returns
}

func OpenSearchEventType_LOW_DISK_WATERMARK_BREACH() OpenSearchEventType {
	_init_.Initialize()
	var returns OpenSearchEventType
	_jsii_.StaticGet(
		"cdk-extensions.alerting.OpenSearchEventType",
		"LOW_DISK_WATERMARK_BREACH",
		&returns,
	)
	return returns
}

func OpenSearchEventType_NODE_RETIREMENT() OpenSearchEventType {
	_init_.Initialize()
	var returns OpenSearchEventType
	_jsii_.StaticGet(
		"cdk-extensions.alerting.OpenSearchEventType",
		"NODE_RETIREMENT",
		&returns,
	)
	return returns
}

func OpenSearchEventType_SERVICE_SOFTWARE_UPDATE() OpenSearchEventType {
	_init_.Initialize()
	var returns OpenSearchEventType
	_jsii_.StaticGet(
		"cdk-extensions.alerting.OpenSearchEventType",
		"SERVICE_SOFTWARE_UPDATE",
		&returns,
	)
	return returns
}

func OpenSearchEventType_VPC_ENDPOINT_CREATE() OpenSearchEventType {
	_init_.Initialize()
	var returns OpenSearchEventType
	_jsii_.StaticGet(
		"cdk-extensions.alerting.OpenSearchEventType",
		"VPC_ENDPOINT_CREATE",
		&returns,
	)
	return returns
}

func OpenSearchEventType_VPC_ENDPOINT_DELETE() OpenSearchEventType {
	_init_.Initialize()
	var returns OpenSearchEventType
	_jsii_.StaticGet(
		"cdk-extensions.alerting.OpenSearchEventType",
		"VPC_ENDPOINT_DELETE",
		&returns,
	)
	return returns
}

func OpenSearchEventType_VPC_ENDPOINT_UPDATE() OpenSearchEventType {
	_init_.Initialize()
	var returns OpenSearchEventType
	_jsii_.StaticGet(
		"cdk-extensions.alerting.OpenSearchEventType",
		"VPC_ENDPOINT_UPDATE",
		&returns,
	)
	return returns
}

