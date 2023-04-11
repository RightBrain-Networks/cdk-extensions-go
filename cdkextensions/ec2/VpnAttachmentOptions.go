package ec2


type VpnAttachmentOptions struct {
	RemoteEndpoint IRemoteVpnEndpoint `field:"required" json:"remoteEndpoint" yaml:"remoteEndpoint"`
	ConnectionType VpnConnectionType `field:"optional" json:"connectionType" yaml:"connectionType"`
	StaticRoutesOnly *bool `field:"optional" json:"staticRoutesOnly" yaml:"staticRoutesOnly"`
	TunnelConfigurations *[]*TunnelOptions `field:"optional" json:"tunnelConfigurations" yaml:"tunnelConfigurations"`
}

