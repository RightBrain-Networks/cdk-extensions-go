package ec2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/vibe-io/cdk-extensions-go/cdkextensions/jsii"
)

type AddressConfiguration interface {
	DefaultNetmaskLength() *float64
	Family() *string
	MaxNetmaskLength() *float64
	MinNetmaskLength() *float64
	PubliclyAdvertisable() *bool
}

// The jsii proxy struct for AddressConfiguration
type jsiiProxy_AddressConfiguration struct {
	_ byte // padding
}

func (j *jsiiProxy_AddressConfiguration) DefaultNetmaskLength() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"defaultNetmaskLength",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AddressConfiguration) Family() *string {
	var returns *string
	_jsii_.Get(
		j,
		"family",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AddressConfiguration) MaxNetmaskLength() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxNetmaskLength",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AddressConfiguration) MinNetmaskLength() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minNetmaskLength",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AddressConfiguration) PubliclyAdvertisable() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"publiclyAdvertisable",
		&returns,
	)
	return returns
}


func AddressConfiguration_Ipv4(options *Ipv4ConfigurationOptions) AddressConfiguration {
	_init_.Initialize()

	if err := validateAddressConfiguration_Ipv4Parameters(options); err != nil {
		panic(err)
	}
	var returns AddressConfiguration

	_jsii_.StaticInvoke(
		"cdk-extensions.ec2.AddressConfiguration",
		"ipv4",
		[]interface{}{options},
		&returns,
	)

	return returns
}

func AddressConfiguration_Ipv6(options *Ipv6ConfigurationOptions) AddressConfiguration {
	_init_.Initialize()

	if err := validateAddressConfiguration_Ipv6Parameters(options); err != nil {
		panic(err)
	}
	var returns AddressConfiguration

	_jsii_.StaticInvoke(
		"cdk-extensions.ec2.AddressConfiguration",
		"ipv6",
		[]interface{}{options},
		&returns,
	)

	return returns
}

func AddressConfiguration_Of(props *AddressConfigurationProps) AddressConfiguration {
	_init_.Initialize()

	if err := validateAddressConfiguration_OfParameters(props); err != nil {
		panic(err)
	}
	var returns AddressConfiguration

	_jsii_.StaticInvoke(
		"cdk-extensions.ec2.AddressConfiguration",
		"of",
		[]interface{}{props},
		&returns,
	)

	return returns
}

