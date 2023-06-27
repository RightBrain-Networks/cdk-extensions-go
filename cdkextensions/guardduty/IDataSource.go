package guardduty

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsguardduty"
	"github.com/aws/constructs-go/constructs/v10"
)

// Represents a data source that should be enabled on a detector.
type IDataSource interface {
	// Associates the data source with a resource configuring a GuardDuty detector.
	Bind(scope constructs.IConstruct) *awsguardduty.CfnDetector_CFNDataSourceConfigurationsProperty
}

// The jsii proxy for IDataSource
type jsiiProxy_IDataSource struct {
	_ byte // padding
}

func (i *jsiiProxy_IDataSource) Bind(scope constructs.IConstruct) *awsguardduty.CfnDetector_CFNDataSourceConfigurationsProperty {
	if err := i.validateBindParameters(scope); err != nil {
		panic(err)
	}
	var returns *awsguardduty.CfnDetector_CFNDataSourceConfigurationsProperty

	_jsii_.Invoke(
		i,
		"bind",
		[]interface{}{scope},
		&returns,
	)

	return returns
}

