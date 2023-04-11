package kinesisfirehose

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudwatch"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/vibe-io/cdk-extensions-go/cdkextensions/kinesisfirehose/internal"
)

type IDeliveryStream interface {
	awsec2.IConnectable
	awsiam.IGrantable
	awscdk.IResource
	Grant(grantee awsiam.IGrantable, actions ...*string) awsiam.Grant
	GrantPutRecords(grantee awsiam.IGrantable) awsiam.Grant
	Metric(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricBackupToS3Bytes(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricBackupToS3DataFreshness(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricBackupToS3Records(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricIncomingBytes(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricIncomingRecords(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	DeliveryStreamArn() *string
	DeliveryStreamName() *string
}

// The jsii proxy for IDeliveryStream
type jsiiProxy_IDeliveryStream struct {
	internal.Type__awsec2IConnectable
	internal.Type__awsiamIGrantable
	internal.Type__awscdkIResource
}

func (i *jsiiProxy_IDeliveryStream) Grant(grantee awsiam.IGrantable, actions ...*string) awsiam.Grant {
	if err := i.validateGrantParameters(grantee); err != nil {
		panic(err)
	}
	args := []interface{}{grantee}
	for _, a := range actions {
		args = append(args, a)
	}

	var returns awsiam.Grant

	_jsii_.Invoke(
		i,
		"grant",
		args,
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IDeliveryStream) GrantPutRecords(grantee awsiam.IGrantable) awsiam.Grant {
	if err := i.validateGrantPutRecordsParameters(grantee); err != nil {
		panic(err)
	}
	var returns awsiam.Grant

	_jsii_.Invoke(
		i,
		"grantPutRecords",
		[]interface{}{grantee},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IDeliveryStream) Metric(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := i.validateMetricParameters(metricName, props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metric",
		[]interface{}{metricName, props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IDeliveryStream) MetricBackupToS3Bytes(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := i.validateMetricBackupToS3BytesParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricBackupToS3Bytes",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IDeliveryStream) MetricBackupToS3DataFreshness(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := i.validateMetricBackupToS3DataFreshnessParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricBackupToS3DataFreshness",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IDeliveryStream) MetricBackupToS3Records(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := i.validateMetricBackupToS3RecordsParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricBackupToS3Records",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IDeliveryStream) MetricIncomingBytes(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := i.validateMetricIncomingBytesParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricIncomingBytes",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IDeliveryStream) MetricIncomingRecords(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := i.validateMetricIncomingRecordsParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricIncomingRecords",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IDeliveryStream) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	if err := i.validateApplyRemovalPolicyParameters(policy); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (j *jsiiProxy_IDeliveryStream) DeliveryStreamArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deliveryStreamArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IDeliveryStream) DeliveryStreamName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deliveryStreamName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IDeliveryStream) Connections() awsec2.Connections {
	var returns awsec2.Connections
	_jsii_.Get(
		j,
		"connections",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IDeliveryStream) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IDeliveryStream) GrantPrincipal() awsiam.IPrincipal {
	var returns awsiam.IPrincipal
	_jsii_.Get(
		j,
		"grantPrincipal",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IDeliveryStream) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IDeliveryStream) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

