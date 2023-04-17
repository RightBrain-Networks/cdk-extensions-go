package ec2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

type IPublicIpamScope interface {
	IIpamScope
	AddAwsProvidedIpv6Pool(id *string, options *AddAwsProvidedIpv6PoolOptions) IIpamPool
	AddByoipIpv4Pool(id *string, options *AddByoipIpv4PoolOptions) IIpamPool
	AddByoipIpv6Pool(id *string, options *AddByoipIpv6PoolOptions) IIpamPool
}

// The jsii proxy for IPublicIpamScope
type jsiiProxy_IPublicIpamScope struct {
	jsiiProxy_IIpamScope
}

func (i *jsiiProxy_IPublicIpamScope) AddAwsProvidedIpv6Pool(id *string, options *AddAwsProvidedIpv6PoolOptions) IIpamPool {
	if err := i.validateAddAwsProvidedIpv6PoolParameters(id, options); err != nil {
		panic(err)
	}
	var returns IIpamPool

	_jsii_.Invoke(
		i,
		"addAwsProvidedIpv6Pool",
		[]interface{}{id, options},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IPublicIpamScope) AddByoipIpv4Pool(id *string, options *AddByoipIpv4PoolOptions) IIpamPool {
	if err := i.validateAddByoipIpv4PoolParameters(id, options); err != nil {
		panic(err)
	}
	var returns IIpamPool

	_jsii_.Invoke(
		i,
		"addByoipIpv4Pool",
		[]interface{}{id, options},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IPublicIpamScope) AddByoipIpv6Pool(id *string, options *AddByoipIpv6PoolOptions) IIpamPool {
	if err := i.validateAddByoipIpv6PoolParameters(id, options); err != nil {
		panic(err)
	}
	var returns IIpamPool

	_jsii_.Invoke(
		i,
		"addByoipIpv6Pool",
		[]interface{}{id, options},
		&returns,
	)

	return returns
}

