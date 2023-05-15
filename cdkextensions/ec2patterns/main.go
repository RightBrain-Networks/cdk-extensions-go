package ec2patterns

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterStruct(
		"cdk-extensions.ec2_patterns.AddAuthorizationRuleOptions",
		reflect.TypeOf((*AddAuthorizationRuleOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk-extensions.ec2_patterns.AddCidrBlockOptions",
		reflect.TypeOf((*AddCidrBlockOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk-extensions.ec2_patterns.AddClientVpnEndpointOptions",
		reflect.TypeOf((*AddClientVpnEndpointOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk-extensions.ec2_patterns.AddHubOptions",
		reflect.TypeOf((*AddHubOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk-extensions.ec2_patterns.AddIsolatedClientVpnEndpointOptions",
		reflect.TypeOf((*AddIsolatedClientVpnEndpointOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk-extensions.ec2_patterns.AddMultiSubnetRouteOptions",
		reflect.TypeOf((*AddMultiSubnetRouteOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk-extensions.ec2_patterns.AddNetworkOptions",
		reflect.TypeOf((*AddNetworkOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk-extensions.ec2_patterns.AddPoolOptions",
		reflect.TypeOf((*AddPoolOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk-extensions.ec2_patterns.AddSpokeNetworkProps",
		reflect.TypeOf((*AddSpokeNetworkProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk-extensions.ec2_patterns.AllocatePrivateNetworkOptions",
		reflect.TypeOf((*AllocatePrivateNetworkOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk-extensions.ec2_patterns.FlowLogOptions",
		reflect.TypeOf((*FlowLogOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk-extensions.ec2_patterns.FourTierNetwork",
		reflect.TypeOf((*FourTierNetwork)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addCidrBlock", GoMethod: "AddCidrBlock"},
			_jsii_.MemberMethod{JsiiMethod: "addClientVpnEndpoint", GoMethod: "AddClientVpnEndpoint"},
			_jsii_.MemberMethod{JsiiMethod: "addFlowLog", GoMethod: "AddFlowLog"},
			_jsii_.MemberMethod{JsiiMethod: "addGatewayEndpoint", GoMethod: "AddGatewayEndpoint"},
			_jsii_.MemberMethod{JsiiMethod: "addInterfaceEndpoint", GoMethod: "AddInterfaceEndpoint"},
			_jsii_.MemberMethod{JsiiMethod: "addIsolatedClientVpnEndpoint", GoMethod: "AddIsolatedClientVpnEndpoint"},
			_jsii_.MemberProperty{JsiiProperty: "addressManager", GoGetter: "AddressManager"},
			_jsii_.MemberMethod{JsiiMethod: "addVpcFlowLog", GoMethod: "AddVpcFlowLog"},
			_jsii_.MemberMethod{JsiiMethod: "addVpnConnection", GoMethod: "AddVpnConnection"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "availabilityZones", GoGetter: "AvailabilityZones"},
			_jsii_.MemberProperty{JsiiProperty: "defaultInstanceTenancy", GoGetter: "DefaultInstanceTenancy"},
			_jsii_.MemberProperty{JsiiProperty: "dnsHostnamesEnabled", GoGetter: "DnsHostnamesEnabled"},
			_jsii_.MemberProperty{JsiiProperty: "dnsSupportEnabled", GoGetter: "DnsSupportEnabled"},
			_jsii_.MemberProperty{JsiiProperty: "enableDnsHostnames", GoGetter: "EnableDnsHostnames"},
			_jsii_.MemberProperty{JsiiProperty: "enableDnsSupport", GoGetter: "EnableDnsSupport"},
			_jsii_.MemberMethod{JsiiMethod: "enableVpnGateway", GoMethod: "EnableVpnGateway"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "incompleteSubnetDefinition", GoGetter: "IncompleteSubnetDefinition"},
			_jsii_.MemberProperty{JsiiProperty: "internetConnectivityEstablished", GoGetter: "InternetConnectivityEstablished"},
			_jsii_.MemberProperty{JsiiProperty: "internetGatewayId", GoGetter: "InternetGatewayId"},
			_jsii_.MemberProperty{JsiiProperty: "ipamPool", GoGetter: "IpamPool"},
			_jsii_.MemberProperty{JsiiProperty: "isolatedSubnets", GoGetter: "IsolatedSubnets"},
			_jsii_.MemberProperty{JsiiProperty: "maxAzs", GoGetter: "MaxAzs"},
			_jsii_.MemberProperty{JsiiProperty: "netmask", GoGetter: "Netmask"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "privateSubnets", GoGetter: "PrivateSubnets"},
			_jsii_.MemberProperty{JsiiProperty: "publicSubnets", GoGetter: "PublicSubnets"},
			_jsii_.MemberMethod{JsiiMethod: "selectSubnetObjects", GoMethod: "SelectSubnetObjects"},
			_jsii_.MemberMethod{JsiiMethod: "selectSubnets", GoMethod: "SelectSubnets"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "vpcArn", GoGetter: "VpcArn"},
			_jsii_.MemberProperty{JsiiProperty: "vpcCidrBlock", GoGetter: "VpcCidrBlock"},
			_jsii_.MemberProperty{JsiiProperty: "vpcCidrBlockAssociations", GoGetter: "VpcCidrBlockAssociations"},
			_jsii_.MemberProperty{JsiiProperty: "vpcDefaultNetworkAcl", GoGetter: "VpcDefaultNetworkAcl"},
			_jsii_.MemberProperty{JsiiProperty: "vpcDefaultSecurityGroup", GoGetter: "VpcDefaultSecurityGroup"},
			_jsii_.MemberProperty{JsiiProperty: "vpcId", GoGetter: "VpcId"},
			_jsii_.MemberProperty{JsiiProperty: "vpcIpv6CidrBlocks", GoGetter: "VpcIpv6CidrBlocks"},
			_jsii_.MemberProperty{JsiiProperty: "vpcName", GoGetter: "VpcName"},
			_jsii_.MemberProperty{JsiiProperty: "vpnGatewayId", GoGetter: "VpnGatewayId"},
		},
		func() interface{} {
			j := jsiiProxy_FourTierNetwork{}
			_jsii_.InitJsiiProxy(&j.Type__awsec2Vpc)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"cdk-extensions.ec2_patterns.FourTierNetworkHub",
		reflect.TypeOf((*FourTierNetworkHub)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addCidrBlock", GoMethod: "AddCidrBlock"},
			_jsii_.MemberMethod{JsiiMethod: "addClientVpnEndpoint", GoMethod: "AddClientVpnEndpoint"},
			_jsii_.MemberMethod{JsiiMethod: "addFlowLog", GoMethod: "AddFlowLog"},
			_jsii_.MemberMethod{JsiiMethod: "addGatewayEndpoint", GoMethod: "AddGatewayEndpoint"},
			_jsii_.MemberMethod{JsiiMethod: "addInterfaceEndpoint", GoMethod: "AddInterfaceEndpoint"},
			_jsii_.MemberMethod{JsiiMethod: "addIsolatedClientVpnEndpoint", GoMethod: "AddIsolatedClientVpnEndpoint"},
			_jsii_.MemberProperty{JsiiProperty: "addressManager", GoGetter: "AddressManager"},
			_jsii_.MemberMethod{JsiiMethod: "addSpoke", GoMethod: "AddSpoke"},
			_jsii_.MemberMethod{JsiiMethod: "addVpcFlowLog", GoMethod: "AddVpcFlowLog"},
			_jsii_.MemberMethod{JsiiMethod: "addVpnConnection", GoMethod: "AddVpnConnection"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "availabilityZones", GoGetter: "AvailabilityZones"},
			_jsii_.MemberProperty{JsiiProperty: "defaultInstanceTenancy", GoGetter: "DefaultInstanceTenancy"},
			_jsii_.MemberProperty{JsiiProperty: "defaultTransitGatewayRouteTable", GoGetter: "DefaultTransitGatewayRouteTable"},
			_jsii_.MemberProperty{JsiiProperty: "dnsHostnamesEnabled", GoGetter: "DnsHostnamesEnabled"},
			_jsii_.MemberProperty{JsiiProperty: "dnsSupportEnabled", GoGetter: "DnsSupportEnabled"},
			_jsii_.MemberProperty{JsiiProperty: "enableDnsHostnames", GoGetter: "EnableDnsHostnames"},
			_jsii_.MemberProperty{JsiiProperty: "enableDnsSupport", GoGetter: "EnableDnsSupport"},
			_jsii_.MemberMethod{JsiiMethod: "enableTransitGateway", GoMethod: "EnableTransitGateway"},
			_jsii_.MemberMethod{JsiiMethod: "enableVpnGateway", GoMethod: "EnableVpnGateway"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "globalNetwork", GoGetter: "GlobalNetwork"},
			_jsii_.MemberProperty{JsiiProperty: "incompleteSubnetDefinition", GoGetter: "IncompleteSubnetDefinition"},
			_jsii_.MemberProperty{JsiiProperty: "internetConnectivityEstablished", GoGetter: "InternetConnectivityEstablished"},
			_jsii_.MemberProperty{JsiiProperty: "internetGatewayId", GoGetter: "InternetGatewayId"},
			_jsii_.MemberProperty{JsiiProperty: "ipamPool", GoGetter: "IpamPool"},
			_jsii_.MemberProperty{JsiiProperty: "isolatedSubnets", GoGetter: "IsolatedSubnets"},
			_jsii_.MemberProperty{JsiiProperty: "maxAzs", GoGetter: "MaxAzs"},
			_jsii_.MemberProperty{JsiiProperty: "netmask", GoGetter: "Netmask"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "privateSubnets", GoGetter: "PrivateSubnets"},
			_jsii_.MemberProperty{JsiiProperty: "publicSubnets", GoGetter: "PublicSubnets"},
			_jsii_.MemberMethod{JsiiMethod: "selectSubnetObjects", GoMethod: "SelectSubnetObjects"},
			_jsii_.MemberMethod{JsiiMethod: "selectSubnets", GoMethod: "SelectSubnets"},
			_jsii_.MemberProperty{JsiiProperty: "sharing", GoGetter: "Sharing"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "transitGateway", GoGetter: "TransitGateway"},
			_jsii_.MemberProperty{JsiiProperty: "vpcArn", GoGetter: "VpcArn"},
			_jsii_.MemberProperty{JsiiProperty: "vpcCidrBlock", GoGetter: "VpcCidrBlock"},
			_jsii_.MemberProperty{JsiiProperty: "vpcCidrBlockAssociations", GoGetter: "VpcCidrBlockAssociations"},
			_jsii_.MemberProperty{JsiiProperty: "vpcDefaultNetworkAcl", GoGetter: "VpcDefaultNetworkAcl"},
			_jsii_.MemberProperty{JsiiProperty: "vpcDefaultSecurityGroup", GoGetter: "VpcDefaultSecurityGroup"},
			_jsii_.MemberProperty{JsiiProperty: "vpcId", GoGetter: "VpcId"},
			_jsii_.MemberProperty{JsiiProperty: "vpcIpv6CidrBlocks", GoGetter: "VpcIpv6CidrBlocks"},
			_jsii_.MemberProperty{JsiiProperty: "vpcName", GoGetter: "VpcName"},
			_jsii_.MemberProperty{JsiiProperty: "vpnGatewayId", GoGetter: "VpnGatewayId"},
		},
		func() interface{} {
			j := jsiiProxy_FourTierNetworkHub{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_FourTierNetwork)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdk-extensions.ec2_patterns.FourTierNetworkHubProps",
		reflect.TypeOf((*FourTierNetworkHubProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk-extensions.ec2_patterns.FourTierNetworkProps",
		reflect.TypeOf((*FourTierNetworkProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk-extensions.ec2_patterns.FourTierNetworkShareProperties",
		reflect.TypeOf((*FourTierNetworkShareProperties)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk-extensions.ec2_patterns.FourTierNetworkSpoke",
		reflect.TypeOf((*FourTierNetworkSpoke)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addCidrBlock", GoMethod: "AddCidrBlock"},
			_jsii_.MemberMethod{JsiiMethod: "addClientVpnEndpoint", GoMethod: "AddClientVpnEndpoint"},
			_jsii_.MemberMethod{JsiiMethod: "addFlowLog", GoMethod: "AddFlowLog"},
			_jsii_.MemberMethod{JsiiMethod: "addGatewayEndpoint", GoMethod: "AddGatewayEndpoint"},
			_jsii_.MemberMethod{JsiiMethod: "addInterfaceEndpoint", GoMethod: "AddInterfaceEndpoint"},
			_jsii_.MemberMethod{JsiiMethod: "addIsolatedClientVpnEndpoint", GoMethod: "AddIsolatedClientVpnEndpoint"},
			_jsii_.MemberProperty{JsiiProperty: "addressManager", GoGetter: "AddressManager"},
			_jsii_.MemberMethod{JsiiMethod: "addVpcFlowLog", GoMethod: "AddVpcFlowLog"},
			_jsii_.MemberMethod{JsiiMethod: "addVpnConnection", GoMethod: "AddVpnConnection"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "availabilityZones", GoGetter: "AvailabilityZones"},
			_jsii_.MemberProperty{JsiiProperty: "defaultInstanceTenancy", GoGetter: "DefaultInstanceTenancy"},
			_jsii_.MemberProperty{JsiiProperty: "dnsHostnamesEnabled", GoGetter: "DnsHostnamesEnabled"},
			_jsii_.MemberProperty{JsiiProperty: "dnsSupportEnabled", GoGetter: "DnsSupportEnabled"},
			_jsii_.MemberProperty{JsiiProperty: "enableDnsHostnames", GoGetter: "EnableDnsHostnames"},
			_jsii_.MemberProperty{JsiiProperty: "enableDnsSupport", GoGetter: "EnableDnsSupport"},
			_jsii_.MemberMethod{JsiiMethod: "enableVpnGateway", GoMethod: "EnableVpnGateway"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "incompleteSubnetDefinition", GoGetter: "IncompleteSubnetDefinition"},
			_jsii_.MemberProperty{JsiiProperty: "internetConnectivityEstablished", GoGetter: "InternetConnectivityEstablished"},
			_jsii_.MemberProperty{JsiiProperty: "internetGatewayId", GoGetter: "InternetGatewayId"},
			_jsii_.MemberProperty{JsiiProperty: "ipamPool", GoGetter: "IpamPool"},
			_jsii_.MemberProperty{JsiiProperty: "isolatedSubnets", GoGetter: "IsolatedSubnets"},
			_jsii_.MemberProperty{JsiiProperty: "maxAzs", GoGetter: "MaxAzs"},
			_jsii_.MemberProperty{JsiiProperty: "netmask", GoGetter: "Netmask"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "privateSubnets", GoGetter: "PrivateSubnets"},
			_jsii_.MemberProperty{JsiiProperty: "publicSubnets", GoGetter: "PublicSubnets"},
			_jsii_.MemberMethod{JsiiMethod: "selectSubnetObjects", GoMethod: "SelectSubnetObjects"},
			_jsii_.MemberMethod{JsiiMethod: "selectSubnets", GoMethod: "SelectSubnets"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "transitGateway", GoGetter: "TransitGateway"},
			_jsii_.MemberProperty{JsiiProperty: "transitGatewayAttachment", GoGetter: "TransitGatewayAttachment"},
			_jsii_.MemberProperty{JsiiProperty: "vpcArn", GoGetter: "VpcArn"},
			_jsii_.MemberProperty{JsiiProperty: "vpcCidrBlock", GoGetter: "VpcCidrBlock"},
			_jsii_.MemberProperty{JsiiProperty: "vpcCidrBlockAssociations", GoGetter: "VpcCidrBlockAssociations"},
			_jsii_.MemberProperty{JsiiProperty: "vpcDefaultNetworkAcl", GoGetter: "VpcDefaultNetworkAcl"},
			_jsii_.MemberProperty{JsiiProperty: "vpcDefaultSecurityGroup", GoGetter: "VpcDefaultSecurityGroup"},
			_jsii_.MemberProperty{JsiiProperty: "vpcId", GoGetter: "VpcId"},
			_jsii_.MemberProperty{JsiiProperty: "vpcIpv6CidrBlocks", GoGetter: "VpcIpv6CidrBlocks"},
			_jsii_.MemberProperty{JsiiProperty: "vpcName", GoGetter: "VpcName"},
			_jsii_.MemberProperty{JsiiProperty: "vpnGatewayId", GoGetter: "VpnGatewayId"},
		},
		func() interface{} {
			j := jsiiProxy_FourTierNetworkSpoke{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_FourTierNetwork)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdk-extensions.ec2_patterns.FourTierNetworkSpokeProps",
		reflect.TypeOf((*FourTierNetworkSpokeProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk-extensions.ec2_patterns.GetClientVpnConfigurationOptions",
		reflect.TypeOf((*GetClientVpnConfigurationOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk-extensions.ec2_patterns.GetClientVpnConfigurationResult",
		reflect.TypeOf((*GetClientVpnConfigurationResult)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk-extensions.ec2_patterns.GetVpcConfigurationOptions",
		reflect.TypeOf((*GetVpcConfigurationOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk-extensions.ec2_patterns.IpAddressManager",
		reflect.TypeOf((*IpAddressManager)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addRegion", GoMethod: "AddRegion"},
			_jsii_.MemberProperty{JsiiProperty: "allowExternalPricipals", GoGetter: "AllowExternalPricipals"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "clientVpnAllocationMask", GoGetter: "ClientVpnAllocationMask"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getClientVpnConfiguration", GoMethod: "GetClientVpnConfiguration"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getVpcConfiguration", GoMethod: "GetVpcConfiguration"},
			_jsii_.MemberProperty{JsiiProperty: "ipam", GoGetter: "Ipam"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "privateVpcPoolForEnvironment", GoMethod: "PrivateVpcPoolForEnvironment"},
			_jsii_.MemberMethod{JsiiMethod: "privateVpcPoolForRegion", GoMethod: "PrivateVpcPoolForRegion"},
			_jsii_.MemberMethod{JsiiMethod: "privateVpnPoolForEnvironment", GoMethod: "PrivateVpnPoolForEnvironment"},
			_jsii_.MemberMethod{JsiiMethod: "privateVpnPoolForRegion", GoMethod: "PrivateVpnPoolForRegion"},
			_jsii_.MemberMethod{JsiiMethod: "registerAccount", GoMethod: "RegisterAccount"},
			_jsii_.MemberMethod{JsiiMethod: "registerCidr", GoMethod: "RegisterCidr"},
			_jsii_.MemberProperty{JsiiProperty: "resourceShare", GoGetter: "ResourceShare"},
			_jsii_.MemberProperty{JsiiProperty: "sharingEnabled", GoGetter: "SharingEnabled"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "vpcAllocationMask", GoGetter: "VpcAllocationMask"},
		},
		func() interface{} {
			j := jsiiProxy_IpAddressManager{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdk-extensions.ec2_patterns.IpAddressManagerProps",
		reflect.TypeOf((*IpAddressManagerProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk-extensions.ec2_patterns.IpAddressManagerSharingProps",
		reflect.TypeOf((*IpAddressManagerSharingProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk-extensions.ec2_patterns.NetworkController",
		reflect.TypeOf((*NetworkController)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addClientVpnEndpoint", GoMethod: "AddClientVpnEndpoint"},
			_jsii_.MemberMethod{JsiiMethod: "addHub", GoMethod: "AddHub"},
			_jsii_.MemberProperty{JsiiProperty: "addressManager", GoGetter: "AddressManager"},
			_jsii_.MemberMethod{JsiiMethod: "addSpoke", GoMethod: "AddSpoke"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "defaultClientVpnNetmask", GoGetter: "DefaultClientVpnNetmask"},
			_jsii_.MemberProperty{JsiiProperty: "defaultNetmask", GoGetter: "DefaultNetmask"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "flowLogBucket", GoGetter: "FlowLogBucket"},
			_jsii_.MemberProperty{JsiiProperty: "flowLogFormat", GoGetter: "FlowLogFormat"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "globalNetwork", GoGetter: "GlobalNetwork"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "registerAccount", GoMethod: "RegisterAccount"},
			_jsii_.MemberMethod{JsiiMethod: "registerCidr", GoMethod: "RegisterCidr"},
			_jsii_.MemberProperty{JsiiProperty: "registeredAccounts", GoGetter: "RegisteredAccounts"},
			_jsii_.MemberProperty{JsiiProperty: "registeredRegions", GoGetter: "RegisteredRegions"},
			_jsii_.MemberMethod{JsiiMethod: "registerRegion", GoMethod: "RegisterRegion"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_NetworkController{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdk-extensions.ec2_patterns.NetworkControllerProps",
		reflect.TypeOf((*NetworkControllerProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk-extensions.ec2_patterns.NetworkIsolatedClientVpnEndpoint",
		reflect.TypeOf((*NetworkIsolatedClientVpnEndpoint)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addAuthorizationRule", GoMethod: "AddAuthorizationRule"},
			_jsii_.MemberMethod{JsiiMethod: "addMultiSubnetRoute", GoMethod: "AddMultiSubnetRoute"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "authorizeAllUsersToVpcCidr", GoGetter: "AuthorizeAllUsersToVpcCidr"},
			_jsii_.MemberProperty{JsiiProperty: "clientCertificate", GoGetter: "ClientCertificate"},
			_jsii_.MemberProperty{JsiiProperty: "clientConnectionHandler", GoGetter: "ClientConnectionHandler"},
			_jsii_.MemberProperty{JsiiProperty: "clientLoginBanner", GoGetter: "ClientLoginBanner"},
			_jsii_.MemberProperty{JsiiProperty: "clientVpnEndpoint", GoGetter: "ClientVpnEndpoint"},
			_jsii_.MemberProperty{JsiiProperty: "connections", GoGetter: "Connections"},
			_jsii_.MemberProperty{JsiiProperty: "description", GoGetter: "Description"},
			_jsii_.MemberProperty{JsiiProperty: "dnsServers", GoGetter: "DnsServers"},
			_jsii_.MemberProperty{JsiiProperty: "endpointId", GoGetter: "EndpointId"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "logging", GoGetter: "Logging"},
			_jsii_.MemberProperty{JsiiProperty: "logGroup", GoGetter: "LogGroup"},
			_jsii_.MemberProperty{JsiiProperty: "logStream", GoGetter: "LogStream"},
			_jsii_.MemberProperty{JsiiProperty: "maxAzs", GoGetter: "MaxAzs"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "port", GoGetter: "Port"},
			_jsii_.MemberProperty{JsiiProperty: "securityGroups", GoGetter: "SecurityGroups"},
			_jsii_.MemberProperty{JsiiProperty: "selfServicePortal", GoGetter: "SelfServicePortal"},
			_jsii_.MemberProperty{JsiiProperty: "serverCertificate", GoGetter: "ServerCertificate"},
			_jsii_.MemberProperty{JsiiProperty: "splitTunnel", GoGetter: "SplitTunnel"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "subnets", GoGetter: "Subnets"},
			_jsii_.MemberProperty{JsiiProperty: "targetNetworksAssociated", GoGetter: "TargetNetworksAssociated"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "transitGateway", GoGetter: "TransitGateway"},
			_jsii_.MemberProperty{JsiiProperty: "transportProtocol", GoGetter: "TransportProtocol"},
			_jsii_.MemberProperty{JsiiProperty: "userBasedAuthentication", GoGetter: "UserBasedAuthentication"},
			_jsii_.MemberProperty{JsiiProperty: "vpc", GoGetter: "Vpc"},
			_jsii_.MemberProperty{JsiiProperty: "vpcCidrBlock", GoGetter: "VpcCidrBlock"},
			_jsii_.MemberProperty{JsiiProperty: "vpnCidr", GoGetter: "VpnCidr"},
		},
		func() interface{} {
			j := jsiiProxy_NetworkIsolatedClientVpnEndpoint{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkResource)
			_jsii_.InitJsiiProxy(&j.Type__awsec2IClientVpnEndpoint)
			_jsii_.InitJsiiProxy(&j.Type__awsec2IConnectable)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdk-extensions.ec2_patterns.NetworkIsolatedClientVpnEndpointProps",
		reflect.TypeOf((*NetworkIsolatedClientVpnEndpointProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk-extensions.ec2_patterns.TransitGatewayHubConfiguration",
		reflect.TypeOf((*TransitGatewayHubConfiguration)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk-extensions.ec2_patterns.TransitGatewaySpokeConfiguration",
		reflect.TypeOf((*TransitGatewaySpokeConfiguration)(nil)).Elem(),
	)
}
