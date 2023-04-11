package kinesisfirehose

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/vibe-io/cdk-extensions-go/cdkextensions/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/constructs-go/constructs/v10"
)

type DeliveryStreamDestination interface {
	Role() awsiam.IRole
	Bind(scope constructs.IConstruct) *DeliveryStreamDestinationConfiguration
}

// The jsii proxy struct for DeliveryStreamDestination
type jsiiProxy_DeliveryStreamDestination struct {
	_ byte // padding
}

func (j *jsiiProxy_DeliveryStreamDestination) Role() awsiam.IRole {
	var returns awsiam.IRole
	_jsii_.Get(
		j,
		"role",
		&returns,
	)
	return returns
}


func NewDeliveryStreamDestination_Override(d DeliveryStreamDestination) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk-extensions.kinesis_firehose.DeliveryStreamDestination",
		nil, // no parameters
		d,
	)
}

func (d *jsiiProxy_DeliveryStreamDestination) Bind(scope constructs.IConstruct) *DeliveryStreamDestinationConfiguration {
	if err := d.validateBindParameters(scope); err != nil {
		panic(err)
	}
	var returns *DeliveryStreamDestinationConfiguration

	_jsii_.Invoke(
		d,
		"bind",
		[]interface{}{scope},
		&returns,
	)

	return returns
}

