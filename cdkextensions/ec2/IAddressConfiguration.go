package ec2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

type IAddressConfiguration interface {
	Family() *string
	PubliclyAdvertisable() *bool
}

// The jsii proxy for IAddressConfiguration
type jsiiProxy_IAddressConfiguration struct {
	_ byte // padding
}

func (j *jsiiProxy_IAddressConfiguration) Family() *string {
	var returns *string
	_jsii_.Get(
		j,
		"family",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IAddressConfiguration) PubliclyAdvertisable() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"publiclyAdvertisable",
		&returns,
	)
	return returns
}

