package guardduty

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/vibe-io/cdk-extensions-go/cdkextensions/jsii"
)

// A data source that will be configured for a GuardDuty detector.
type DataSource interface {
}

// The jsii proxy struct for DataSource
type jsiiProxy_DataSource struct {
	_ byte // padding
}

func NewDataSource() DataSource {
	_init_.Initialize()

	j := jsiiProxy_DataSource{}

	_jsii_.Create(
		"cdk-extensions.guardduty.DataSource",
		nil, // no parameters
		&j,
	)

	return &j
}

func NewDataSource_Override(d DataSource) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk-extensions.guardduty.DataSource",
		nil, // no parameters
		d,
	)
}

// Creates a data source configuration for processing Kubernetes (EKS) events.
//
// Returns: An object representing the configured data source.
func DataSource_Kubernetes(options *KubernetesOptions) IDataSource {
	_init_.Initialize()

	if err := validateDataSource_KubernetesParameters(options); err != nil {
		panic(err)
	}
	var returns IDataSource

	_jsii_.StaticInvoke(
		"cdk-extensions.guardduty.DataSource",
		"kubernetes",
		[]interface{}{options},
		&returns,
	)

	return returns
}

// Creates a data source configuration allowing GuardDuty top perform malware scanning.
//
// Returns: An object representing the configured data source.
func DataSource_MalwareProtection(options *MalwareProtectionOptions) IDataSource {
	_init_.Initialize()

	if err := validateDataSource_MalwareProtectionParameters(options); err != nil {
		panic(err)
	}
	var returns IDataSource

	_jsii_.StaticInvoke(
		"cdk-extensions.guardduty.DataSource",
		"malwareProtection",
		[]interface{}{options},
		&returns,
	)

	return returns
}

// Creates a data source configuration for analyzing S3 data events in GuardDuty.
//
// Returns: An object representing the configured data source.
func DataSource_S3Logs(options *S3LogsOptions) IDataSource {
	_init_.Initialize()

	if err := validateDataSource_S3LogsParameters(options); err != nil {
		panic(err)
	}
	var returns IDataSource

	_jsii_.StaticInvoke(
		"cdk-extensions.guardduty.DataSource",
		"s3Logs",
		[]interface{}{options},
		&returns,
	)

	return returns
}

