package rds

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/vibe-io/cdk-extensions-go/cdkextensions/jsii"
)

type DatabaseProxyEndpointAccess interface {
	Role() *string
}

// The jsii proxy struct for DatabaseProxyEndpointAccess
type jsiiProxy_DatabaseProxyEndpointAccess struct {
	_ byte // padding
}

func (j *jsiiProxy_DatabaseProxyEndpointAccess) Role() *string {
	var returns *string
	_jsii_.Get(
		j,
		"role",
		&returns,
	)
	return returns
}


func DatabaseProxyEndpointAccess_Of(role *string) DatabaseProxyEndpointAccess {
	_init_.Initialize()

	if err := validateDatabaseProxyEndpointAccess_OfParameters(role); err != nil {
		panic(err)
	}
	var returns DatabaseProxyEndpointAccess

	_jsii_.StaticInvoke(
		"cdk-extensions.rds.DatabaseProxyEndpointAccess",
		"of",
		[]interface{}{role},
		&returns,
	)

	return returns
}

func DatabaseProxyEndpointAccess_READ_ONLY() DatabaseProxyEndpointAccess {
	_init_.Initialize()
	var returns DatabaseProxyEndpointAccess
	_jsii_.StaticGet(
		"cdk-extensions.rds.DatabaseProxyEndpointAccess",
		"READ_ONLY",
		&returns,
	)
	return returns
}

func DatabaseProxyEndpointAccess_READ_WRITE() DatabaseProxyEndpointAccess {
	_init_.Initialize()
	var returns DatabaseProxyEndpointAccess
	_jsii_.StaticGet(
		"cdk-extensions.rds.DatabaseProxyEndpointAccess",
		"READ_WRITE",
		&returns,
	)
	return returns
}

