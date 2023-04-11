package ec2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/vibe-io/cdk-extensions-go/cdkextensions/jsii"
)

// Represents a VPN protocol that can be used to establish a connection.
type VpnConnectionType interface {
	// The name of the VPN protocol.
	Name() *string
}

// The jsii proxy struct for VpnConnectionType
type jsiiProxy_VpnConnectionType struct {
	_ byte // padding
}

func (j *jsiiProxy_VpnConnectionType) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}


// An escape hatch method that allows defining custom VPN protocols.
//
// Returns: A VpnConnectionType object representing the specified protocol.
func VpnConnectionType_Of(name *string) VpnConnectionType {
	_init_.Initialize()

	if err := validateVpnConnectionType_OfParameters(name); err != nil {
		panic(err)
	}
	var returns VpnConnectionType

	_jsii_.StaticInvoke(
		"cdk-extensions.ec2.VpnConnectionType",
		"of",
		[]interface{}{name},
		&returns,
	)

	return returns
}

func VpnConnectionType_IPSEC_1() VpnConnectionType {
	_init_.Initialize()
	var returns VpnConnectionType
	_jsii_.StaticGet(
		"cdk-extensions.ec2.VpnConnectionType",
		"IPSEC_1",
		&returns,
	)
	return returns
}

