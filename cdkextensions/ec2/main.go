package ec2

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterStruct(
		"cdk-extensions.ec2.AddChildPoolOptions",
		reflect.TypeOf((*AddChildPoolOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk-extensions.ec2.AddCidrToPoolOptions",
		reflect.TypeOf((*AddCidrToPoolOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk-extensions.ec2.AddCidrToPoolResult",
		reflect.TypeOf((*AddCidrToPoolResult)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk-extensions.ec2.AddressConfiguration",
		reflect.TypeOf((*AddressConfiguration)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "defaultNetmaskLength", GoGetter: "DefaultNetmaskLength"},
			_jsii_.MemberProperty{JsiiProperty: "family", GoGetter: "Family"},
			_jsii_.MemberProperty{JsiiProperty: "maxNetmaskLength", GoGetter: "MaxNetmaskLength"},
			_jsii_.MemberProperty{JsiiProperty: "minNetmaskLength", GoGetter: "MinNetmaskLength"},
			_jsii_.MemberProperty{JsiiProperty: "publiclyAdvertisable", GoGetter: "PubliclyAdvertisable"},
		},
		func() interface{} {
			return &jsiiProxy_AddressConfiguration{}
		},
	)
	_jsii_.RegisterStruct(
		"cdk-extensions.ec2.AddressConfigurationProps",
		reflect.TypeOf((*AddressConfigurationProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk-extensions.ec2.AllocateCidrFromPoolOptions",
		reflect.TypeOf((*AllocateCidrFromPoolOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk-extensions.ec2.CidrProvider",
		reflect.TypeOf((*CidrProvider)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_CidrProvider{}
		},
	)
	_jsii_.RegisterClass(
		"cdk-extensions.ec2.CustomerGateway",
		reflect.TypeOf((*CustomerGateway)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "bgpAsn", GoGetter: "BgpAsn"},
			_jsii_.MemberProperty{JsiiProperty: "connectionType", GoGetter: "ConnectionType"},
			_jsii_.MemberProperty{JsiiProperty: "customerGatewayAsn", GoGetter: "CustomerGatewayAsn"},
			_jsii_.MemberProperty{JsiiProperty: "customerGatewayId", GoGetter: "CustomerGatewayId"},
			_jsii_.MemberProperty{JsiiProperty: "customerGatewayIp", GoGetter: "CustomerGatewayIp"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "ipAddress", GoGetter: "IpAddress"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "resource", GoGetter: "Resource"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_CustomerGateway{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkResource)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_ICustomerGateway)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdk-extensions.ec2.CustomerGatewayAttributes",
		reflect.TypeOf((*CustomerGatewayAttributes)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk-extensions.ec2.CustomerGatewayConfigurationRemoteVpnEndpoint",
		reflect.TypeOf((*CustomerGatewayConfigurationRemoteVpnEndpoint)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
			_jsii_.MemberProperty{JsiiProperty: "configuration", GoGetter: "Configuration"},
			_jsii_.MemberProperty{JsiiProperty: "customerGateway", GoGetter: "CustomerGateway"},
		},
		func() interface{} {
			j := jsiiProxy_CustomerGatewayConfigurationRemoteVpnEndpoint{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IRemoteVpnEndpoint)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdk-extensions.ec2.CustomerGatewayProps",
		reflect.TypeOf((*CustomerGatewayProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk-extensions.ec2.CustomerGatewayRemoteVpnEndpoint",
		reflect.TypeOf((*CustomerGatewayRemoteVpnEndpoint)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
			_jsii_.MemberProperty{JsiiProperty: "customerGateway", GoGetter: "CustomerGateway"},
		},
		func() interface{} {
			j := jsiiProxy_CustomerGatewayRemoteVpnEndpoint{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IRemoteVpnEndpoint)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"cdk-extensions.ec2.FlowLog",
		reflect.TypeOf((*FlowLog)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "bucket", GoGetter: "Bucket"},
			_jsii_.MemberProperty{JsiiProperty: "destination", GoGetter: "Destination"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "flowLogId", GoGetter: "FlowLogId"},
			_jsii_.MemberProperty{JsiiProperty: "flowLogName", GoGetter: "FlowLogName"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "iamRole", GoGetter: "IamRole"},
			_jsii_.MemberProperty{JsiiProperty: "keyPrefix", GoGetter: "KeyPrefix"},
			_jsii_.MemberProperty{JsiiProperty: "logFormat", GoGetter: "LogFormat"},
			_jsii_.MemberProperty{JsiiProperty: "logGroup", GoGetter: "LogGroup"},
			_jsii_.MemberProperty{JsiiProperty: "maxAggregationInterval", GoGetter: "MaxAggregationInterval"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "resource", GoGetter: "Resource"},
			_jsii_.MemberProperty{JsiiProperty: "resourceType", GoGetter: "ResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "trafficType", GoGetter: "TrafficType"},
		},
		func() interface{} {
			j := jsiiProxy_FlowLog{}
			_jsii_.InitJsiiProxy(&j.Type__awsec2FlowLog)
			return &j
		},
	)
	_jsii_.RegisterEnum(
		"cdk-extensions.ec2.FlowLogDataType",
		reflect.TypeOf((*FlowLogDataType)(nil)).Elem(),
		map[string]interface{}{
			"INT_32": FlowLogDataType_INT_32,
			"INT_64": FlowLogDataType_INT_64,
			"STRING": FlowLogDataType_STRING,
		},
	)
	_jsii_.RegisterClass(
		"cdk-extensions.ec2.FlowLogDestination",
		reflect.TypeOf((*FlowLogDestination)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
		},
		func() interface{} {
			j := jsiiProxy_FlowLogDestination{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_ILogDestination)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdk-extensions.ec2.FlowLogDestinationConfig",
		reflect.TypeOf((*FlowLogDestinationConfig)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk-extensions.ec2.FlowLogField",
		reflect.TypeOf((*FlowLogField)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "type", GoGetter: "Type"},
		},
		func() interface{} {
			return &jsiiProxy_FlowLogField{}
		},
	)
	_jsii_.RegisterEnum(
		"cdk-extensions.ec2.FlowLogFileFormat",
		reflect.TypeOf((*FlowLogFileFormat)(nil)).Elem(),
		map[string]interface{}{
			"PARQUET": FlowLogFileFormat_PARQUET,
			"PLAIN_TEXT": FlowLogFileFormat_PLAIN_TEXT,
		},
	)
	_jsii_.RegisterClass(
		"cdk-extensions.ec2.FlowLogFormat",
		reflect.TypeOf((*FlowLogFormat)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addField", GoMethod: "AddField"},
			_jsii_.MemberProperty{JsiiProperty: "fields", GoGetter: "Fields"},
			_jsii_.MemberProperty{JsiiProperty: "template", GoGetter: "Template"},
		},
		func() interface{} {
			return &jsiiProxy_FlowLogFormat{}
		},
	)
	_jsii_.RegisterStruct(
		"cdk-extensions.ec2.FlowLogProps",
		reflect.TypeOf((*FlowLogProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk-extensions.ec2.FlowLogS3Options",
		reflect.TypeOf((*FlowLogS3Options)(nil)).Elem(),
	)
	_jsii_.RegisterInterface(
		"cdk-extensions.ec2.IAddressConfiguration",
		reflect.TypeOf((*IAddressConfiguration)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "family", GoGetter: "Family"},
			_jsii_.MemberProperty{JsiiProperty: "publiclyAdvertisable", GoGetter: "PubliclyAdvertisable"},
		},
		func() interface{} {
			return &jsiiProxy_IAddressConfiguration{}
		},
	)
	_jsii_.RegisterInterface(
		"cdk-extensions.ec2.ICidrProvider",
		reflect.TypeOf((*ICidrProvider)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "ipamOptions", GoGetter: "IpamOptions"},
			_jsii_.MemberProperty{JsiiProperty: "ipamPool", GoGetter: "IpamPool"},
			_jsii_.MemberProperty{JsiiProperty: "netmask", GoGetter: "Netmask"},
		},
		func() interface{} {
			return &jsiiProxy_ICidrProvider{}
		},
	)
	_jsii_.RegisterInterface(
		"cdk-extensions.ec2.ICustomerGateway",
		reflect.TypeOf((*ICustomerGateway)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "customerGatewayAsn", GoGetter: "CustomerGatewayAsn"},
			_jsii_.MemberProperty{JsiiProperty: "customerGatewayId", GoGetter: "CustomerGatewayId"},
			_jsii_.MemberProperty{JsiiProperty: "customerGatewayIp", GoGetter: "CustomerGatewayIp"},
		},
		func() interface{} {
			return &jsiiProxy_ICustomerGateway{}
		},
	)
	_jsii_.RegisterInterface(
		"cdk-extensions.ec2.IIpam",
		reflect.TypeOf((*IIpam)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addScope", GoMethod: "AddScope"},
			_jsii_.MemberMethod{JsiiMethod: "associateResourceDiscovery", GoMethod: "AssociateResourceDiscovery"},
			_jsii_.MemberProperty{JsiiProperty: "ipamArn", GoGetter: "IpamArn"},
			_jsii_.MemberProperty{JsiiProperty: "ipamId", GoGetter: "IpamId"},
			_jsii_.MemberProperty{JsiiProperty: "ipamPrivateDefaultScopeId", GoGetter: "IpamPrivateDefaultScopeId"},
			_jsii_.MemberProperty{JsiiProperty: "ipamPublicDefaultScopeId", GoGetter: "IpamPublicDefaultScopeId"},
			_jsii_.MemberProperty{JsiiProperty: "ipamScopeCount", GoGetter: "IpamScopeCount"},
		},
		func() interface{} {
			return &jsiiProxy_IIpam{}
		},
	)
	_jsii_.RegisterInterface(
		"cdk-extensions.ec2.IIpamAllocation",
		reflect.TypeOf((*IIpamAllocation)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "ipamAllocationCidr", GoGetter: "IpamAllocationCidr"},
			_jsii_.MemberProperty{JsiiProperty: "ipamAllocationId", GoGetter: "IpamAllocationId"},
		},
		func() interface{} {
			return &jsiiProxy_IIpamAllocation{}
		},
	)
	_jsii_.RegisterInterface(
		"cdk-extensions.ec2.IIpamAllocationConfiguration",
		reflect.TypeOf((*IIpamAllocationConfiguration)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
		},
		func() interface{} {
			return &jsiiProxy_IIpamAllocationConfiguration{}
		},
	)
	_jsii_.RegisterInterface(
		"cdk-extensions.ec2.IIpamPool",
		reflect.TypeOf((*IIpamPool)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addChildPool", GoMethod: "AddChildPool"},
			_jsii_.MemberMethod{JsiiMethod: "addCidrToPool", GoMethod: "AddCidrToPool"},
			_jsii_.MemberMethod{JsiiMethod: "allocateCidrFromPool", GoMethod: "AllocateCidrFromPool"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "consumer", GoGetter: "Consumer"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "ipamPoolArn", GoGetter: "IpamPoolArn"},
			_jsii_.MemberProperty{JsiiProperty: "ipamPoolDepth", GoGetter: "IpamPoolDepth"},
			_jsii_.MemberProperty{JsiiProperty: "ipamPoolId", GoGetter: "IpamPoolId"},
			_jsii_.MemberProperty{JsiiProperty: "ipamPoolIpamArn", GoGetter: "IpamPoolIpamArn"},
			_jsii_.MemberProperty{JsiiProperty: "ipamPoolScopeArn", GoGetter: "IpamPoolScopeArn"},
			_jsii_.MemberProperty{JsiiProperty: "ipamPoolScopeType", GoGetter: "IpamPoolScopeType"},
			_jsii_.MemberProperty{JsiiProperty: "ipamPoolState", GoGetter: "IpamPoolState"},
			_jsii_.MemberProperty{JsiiProperty: "ipamPoolStateMessage", GoGetter: "IpamPoolStateMessage"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
		},
		func() interface{} {
			j := jsiiProxy_IIpamPool{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkIResource)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"cdk-extensions.ec2.IIpamPoolCidr",
		reflect.TypeOf((*IIpamPoolCidr)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "ipamPoolCidrId", GoGetter: "IpamPoolCidrId"},
			_jsii_.MemberProperty{JsiiProperty: "ipamPoolCidrState", GoGetter: "IpamPoolCidrState"},
		},
		func() interface{} {
			return &jsiiProxy_IIpamPoolCidr{}
		},
	)
	_jsii_.RegisterInterface(
		"cdk-extensions.ec2.IIpamPoolCidrConfiguration",
		reflect.TypeOf((*IIpamPoolCidrConfiguration)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
			_jsii_.MemberProperty{JsiiProperty: "inline", GoGetter: "Inline"},
		},
		func() interface{} {
			return &jsiiProxy_IIpamPoolCidrConfiguration{}
		},
	)
	_jsii_.RegisterInterface(
		"cdk-extensions.ec2.IIpamResourceDiscovery",
		reflect.TypeOf((*IIpamResourceDiscovery)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addIpam", GoMethod: "AddIpam"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberMethod{JsiiMethod: "associateIpam", GoMethod: "AssociateIpam"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "ipamResourceDiscoveryArn", GoGetter: "IpamResourceDiscoveryArn"},
			_jsii_.MemberProperty{JsiiProperty: "ipamResourceDiscoveryId", GoGetter: "IpamResourceDiscoveryId"},
			_jsii_.MemberProperty{JsiiProperty: "ipamResourceDiscoveryIsDefault", GoGetter: "IpamResourceDiscoveryIsDefault"},
			_jsii_.MemberProperty{JsiiProperty: "ipamResourceDiscoveryOwnerId", GoGetter: "IpamResourceDiscoveryOwnerId"},
			_jsii_.MemberProperty{JsiiProperty: "ipamResourceDiscoveryRegion", GoGetter: "IpamResourceDiscoveryRegion"},
			_jsii_.MemberProperty{JsiiProperty: "ipamResourceDiscoveryState", GoGetter: "IpamResourceDiscoveryState"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
		},
		func() interface{} {
			j := jsiiProxy_IIpamResourceDiscovery{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkIResource)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"cdk-extensions.ec2.IIpamResourceDiscoveryAssociation",
		reflect.TypeOf((*IIpamResourceDiscoveryAssociation)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "ipamResourceDiscoveryAssociationArn", GoGetter: "IpamResourceDiscoveryAssociationArn"},
			_jsii_.MemberProperty{JsiiProperty: "ipamResourceDiscoveryAssociationId", GoGetter: "IpamResourceDiscoveryAssociationId"},
			_jsii_.MemberProperty{JsiiProperty: "ipamResourceDiscoveryAssociationIpamArn", GoGetter: "IpamResourceDiscoveryAssociationIpamArn"},
			_jsii_.MemberProperty{JsiiProperty: "ipamResourceDiscoveryAssociationIpamRegion", GoGetter: "IpamResourceDiscoveryAssociationIpamRegion"},
			_jsii_.MemberProperty{JsiiProperty: "ipamResourceDiscoveryAssociationIsDefault", GoGetter: "IpamResourceDiscoveryAssociationIsDefault"},
			_jsii_.MemberProperty{JsiiProperty: "ipamResourceDiscoveryAssociationOwnerId", GoGetter: "IpamResourceDiscoveryAssociationOwnerId"},
			_jsii_.MemberProperty{JsiiProperty: "ipamResourceDiscoveryAssociationResourceDiscoveryId", GoGetter: "IpamResourceDiscoveryAssociationResourceDiscoveryId"},
			_jsii_.MemberProperty{JsiiProperty: "ipamResourceDiscoveryAssociationResourceDiscoveryStatus", GoGetter: "IpamResourceDiscoveryAssociationResourceDiscoveryStatus"},
			_jsii_.MemberProperty{JsiiProperty: "ipamResourceDiscoveryAssociationState", GoGetter: "IpamResourceDiscoveryAssociationState"},
		},
		func() interface{} {
			return &jsiiProxy_IIpamResourceDiscoveryAssociation{}
		},
	)
	_jsii_.RegisterInterface(
		"cdk-extensions.ec2.IIpamScope",
		reflect.TypeOf((*IIpamScope)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addPool", GoMethod: "AddPool"},
			_jsii_.MemberProperty{JsiiProperty: "ipamScopeArn", GoGetter: "IpamScopeArn"},
			_jsii_.MemberProperty{JsiiProperty: "ipamScopeId", GoGetter: "IpamScopeId"},
			_jsii_.MemberProperty{JsiiProperty: "ipamScopeIpamArn", GoGetter: "IpamScopeIpamArn"},
			_jsii_.MemberProperty{JsiiProperty: "ipamScopeIsDefault", GoGetter: "IpamScopeIsDefault"},
			_jsii_.MemberProperty{JsiiProperty: "ipamScopePoolCount", GoGetter: "IpamScopePoolCount"},
			_jsii_.MemberProperty{JsiiProperty: "ipamScopeType", GoGetter: "IpamScopeType"},
		},
		func() interface{} {
			return &jsiiProxy_IIpamScope{}
		},
	)
	_jsii_.RegisterInterface(
		"cdk-extensions.ec2.ILocalVpnEndpoint",
		reflect.TypeOf((*ILocalVpnEndpoint)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
		},
		func() interface{} {
			return &jsiiProxy_ILocalVpnEndpoint{}
		},
	)
	_jsii_.RegisterInterface(
		"cdk-extensions.ec2.ILogDestination",
		reflect.TypeOf((*ILogDestination)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
		},
		func() interface{} {
			return &jsiiProxy_ILogDestination{}
		},
	)
	_jsii_.RegisterInterface(
		"cdk-extensions.ec2.IRemoteVpnEndpoint",
		reflect.TypeOf((*IRemoteVpnEndpoint)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
		},
		func() interface{} {
			return &jsiiProxy_IRemoteVpnEndpoint{}
		},
	)
	_jsii_.RegisterInterface(
		"cdk-extensions.ec2.ITransitGateway",
		reflect.TypeOf((*ITransitGateway)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addRouteTable", GoMethod: "AddRouteTable"},
			_jsii_.MemberMethod{JsiiMethod: "addVpn", GoMethod: "AddVpn"},
			_jsii_.MemberMethod{JsiiMethod: "attachVpc", GoMethod: "AttachVpc"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "transitGatewayArn", GoGetter: "TransitGatewayArn"},
			_jsii_.MemberProperty{JsiiProperty: "transitGatewayId", GoGetter: "TransitGatewayId"},
		},
		func() interface{} {
			j := jsiiProxy_ITransitGateway{}
			_jsii_.InitJsiiProxy(&j.Type__constructsIConstruct)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"cdk-extensions.ec2.ITransitGatewayAttachment",
		reflect.TypeOf((*ITransitGatewayAttachment)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addRoute", GoMethod: "AddRoute"},
			_jsii_.MemberProperty{JsiiProperty: "transitGatewayAttachmentArn", GoGetter: "TransitGatewayAttachmentArn"},
			_jsii_.MemberProperty{JsiiProperty: "transitGatewayAttachmentId", GoGetter: "TransitGatewayAttachmentId"},
		},
		func() interface{} {
			return &jsiiProxy_ITransitGatewayAttachment{}
		},
	)
	_jsii_.RegisterInterface(
		"cdk-extensions.ec2.ITransitGatewayPeeringAttachment",
		reflect.TypeOf((*ITransitGatewayPeeringAttachment)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addRoute", GoMethod: "AddRoute"},
			_jsii_.MemberProperty{JsiiProperty: "transitGatewayAttachmentArn", GoGetter: "TransitGatewayAttachmentArn"},
			_jsii_.MemberProperty{JsiiProperty: "transitGatewayAttachmentCreationTime", GoGetter: "TransitGatewayAttachmentCreationTime"},
			_jsii_.MemberProperty{JsiiProperty: "transitGatewayAttachmentId", GoGetter: "TransitGatewayAttachmentId"},
			_jsii_.MemberProperty{JsiiProperty: "transitGatewayAttachmentState", GoGetter: "TransitGatewayAttachmentState"},
			_jsii_.MemberProperty{JsiiProperty: "transitGatewayAttachmentStatus", GoGetter: "TransitGatewayAttachmentStatus"},
			_jsii_.MemberProperty{JsiiProperty: "transitGatewayAttachmentStatusCode", GoGetter: "TransitGatewayAttachmentStatusCode"},
			_jsii_.MemberProperty{JsiiProperty: "transitGatewayAttachmentStatusMessage", GoGetter: "TransitGatewayAttachmentStatusMessage"},
		},
		func() interface{} {
			j := jsiiProxy_ITransitGatewayPeeringAttachment{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_ITransitGatewayAttachment)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"cdk-extensions.ec2.ITransitGatewayRoute",
		reflect.TypeOf((*ITransitGatewayRoute)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "transitGatewayRouteId", GoGetter: "TransitGatewayRouteId"},
		},
		func() interface{} {
			return &jsiiProxy_ITransitGatewayRoute{}
		},
	)
	_jsii_.RegisterInterface(
		"cdk-extensions.ec2.ITransitGatewayRouteTable",
		reflect.TypeOf((*ITransitGatewayRouteTable)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addRoute", GoMethod: "AddRoute"},
			_jsii_.MemberProperty{JsiiProperty: "transitGatewayRouteTableArn", GoGetter: "TransitGatewayRouteTableArn"},
			_jsii_.MemberProperty{JsiiProperty: "transitGatewayRouteTableId", GoGetter: "TransitGatewayRouteTableId"},
		},
		func() interface{} {
			return &jsiiProxy_ITransitGatewayRouteTable{}
		},
	)
	_jsii_.RegisterClass(
		"cdk-extensions.ec2.Ipam",
		reflect.TypeOf((*Ipam)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addRegion", GoMethod: "AddRegion"},
			_jsii_.MemberMethod{JsiiMethod: "addScope", GoMethod: "AddScope"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberMethod{JsiiMethod: "associateResourceDiscovery", GoMethod: "AssociateResourceDiscovery"},
			_jsii_.MemberProperty{JsiiProperty: "defaultPrivateScope", GoGetter: "DefaultPrivateScope"},
			_jsii_.MemberProperty{JsiiProperty: "defaultPublicScope", GoGetter: "DefaultPublicScope"},
			_jsii_.MemberProperty{JsiiProperty: "description", GoGetter: "Description"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "ipamArn", GoGetter: "IpamArn"},
			_jsii_.MemberProperty{JsiiProperty: "ipamId", GoGetter: "IpamId"},
			_jsii_.MemberProperty{JsiiProperty: "ipamPrivateDefaultScopeId", GoGetter: "IpamPrivateDefaultScopeId"},
			_jsii_.MemberProperty{JsiiProperty: "ipamPublicDefaultScopeId", GoGetter: "IpamPublicDefaultScopeId"},
			_jsii_.MemberProperty{JsiiProperty: "ipamScopeCount", GoGetter: "IpamScopeCount"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "regions", GoGetter: "Regions"},
			_jsii_.MemberProperty{JsiiProperty: "resource", GoGetter: "Resource"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_Ipam{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkResource)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IIpam)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"cdk-extensions.ec2.IpamAllocation",
		reflect.TypeOf((*IpamAllocation)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "allocation", GoGetter: "Allocation"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "description", GoGetter: "Description"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "ipamAllocationCidr", GoGetter: "IpamAllocationCidr"},
			_jsii_.MemberProperty{JsiiProperty: "ipamAllocationId", GoGetter: "IpamAllocationId"},
			_jsii_.MemberProperty{JsiiProperty: "ipamPool", GoGetter: "IpamPool"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "resource", GoGetter: "Resource"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_IpamAllocation{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkResource)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IIpamAllocation)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"cdk-extensions.ec2.IpamAllocationConfiguration",
		reflect.TypeOf((*IpamAllocationConfiguration)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_IpamAllocationConfiguration{}
		},
	)
	_jsii_.RegisterStruct(
		"cdk-extensions.ec2.IpamAllocationOptions",
		reflect.TypeOf((*IpamAllocationOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk-extensions.ec2.IpamAllocationProps",
		reflect.TypeOf((*IpamAllocationProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk-extensions.ec2.IpamAttributes",
		reflect.TypeOf((*IpamAttributes)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk-extensions.ec2.IpamConsumer",
		reflect.TypeOf((*IpamConsumer)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
		},
		func() interface{} {
			return &jsiiProxy_IpamConsumer{}
		},
	)
	_jsii_.RegisterClass(
		"cdk-extensions.ec2.IpamPool",
		reflect.TypeOf((*IpamPool)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addChildPool", GoMethod: "AddChildPool"},
			_jsii_.MemberMethod{JsiiMethod: "addCidrToPool", GoMethod: "AddCidrToPool"},
			_jsii_.MemberProperty{JsiiProperty: "addressConfiguration", GoGetter: "AddressConfiguration"},
			_jsii_.MemberMethod{JsiiMethod: "addTagRestriction", GoMethod: "AddTagRestriction"},
			_jsii_.MemberMethod{JsiiMethod: "allocateCidrFromPool", GoMethod: "AllocateCidrFromPool"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "autoImport", GoGetter: "AutoImport"},
			_jsii_.MemberProperty{JsiiProperty: "consumer", GoGetter: "Consumer"},
			_jsii_.MemberProperty{JsiiProperty: "description", GoGetter: "Description"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "ipamPoolArn", GoGetter: "IpamPoolArn"},
			_jsii_.MemberProperty{JsiiProperty: "ipamPoolDepth", GoGetter: "IpamPoolDepth"},
			_jsii_.MemberProperty{JsiiProperty: "ipamPoolId", GoGetter: "IpamPoolId"},
			_jsii_.MemberProperty{JsiiProperty: "ipamPoolIpamArn", GoGetter: "IpamPoolIpamArn"},
			_jsii_.MemberProperty{JsiiProperty: "ipamPoolScopeArn", GoGetter: "IpamPoolScopeArn"},
			_jsii_.MemberProperty{JsiiProperty: "ipamPoolScopeType", GoGetter: "IpamPoolScopeType"},
			_jsii_.MemberProperty{JsiiProperty: "ipamPoolState", GoGetter: "IpamPoolState"},
			_jsii_.MemberProperty{JsiiProperty: "ipamPoolStateMessage", GoGetter: "IpamPoolStateMessage"},
			_jsii_.MemberProperty{JsiiProperty: "ipamScope", GoGetter: "IpamScope"},
			_jsii_.MemberProperty{JsiiProperty: "locale", GoGetter: "Locale"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "parentPool", GoGetter: "ParentPool"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "publicIpSource", GoGetter: "PublicIpSource"},
			_jsii_.MemberProperty{JsiiProperty: "resource", GoGetter: "Resource"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_IpamPool{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkResource)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IIpamPool)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"cdk-extensions.ec2.IpamPoolCidr",
		reflect.TypeOf((*IpamPoolCidr)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "configuration", GoGetter: "Configuration"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "ipamPool", GoGetter: "IpamPool"},
			_jsii_.MemberProperty{JsiiProperty: "ipamPoolCidrId", GoGetter: "IpamPoolCidrId"},
			_jsii_.MemberProperty{JsiiProperty: "ipamPoolCidrState", GoGetter: "IpamPoolCidrState"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "resource", GoGetter: "Resource"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_IpamPoolCidr{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkResource)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IIpamPoolCidr)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"cdk-extensions.ec2.IpamPoolCidrConfiguration",
		reflect.TypeOf((*IpamPoolCidrConfiguration)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_IpamPoolCidrConfiguration{}
		},
	)
	_jsii_.RegisterStruct(
		"cdk-extensions.ec2.IpamPoolCidrProps",
		reflect.TypeOf((*IpamPoolCidrProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk-extensions.ec2.IpamPoolOptions",
		reflect.TypeOf((*IpamPoolOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk-extensions.ec2.IpamPoolProps",
		reflect.TypeOf((*IpamPoolProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk-extensions.ec2.IpamProps",
		reflect.TypeOf((*IpamProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk-extensions.ec2.IpamResourceDiscovery",
		reflect.TypeOf((*IpamResourceDiscovery)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addIpam", GoMethod: "AddIpam"},
			_jsii_.MemberMethod{JsiiMethod: "addRegion", GoMethod: "AddRegion"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberMethod{JsiiMethod: "associateIpam", GoMethod: "AssociateIpam"},
			_jsii_.MemberProperty{JsiiProperty: "description", GoGetter: "Description"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "ipamResourceDiscoveryArn", GoGetter: "IpamResourceDiscoveryArn"},
			_jsii_.MemberProperty{JsiiProperty: "ipamResourceDiscoveryId", GoGetter: "IpamResourceDiscoveryId"},
			_jsii_.MemberProperty{JsiiProperty: "ipamResourceDiscoveryIsDefault", GoGetter: "IpamResourceDiscoveryIsDefault"},
			_jsii_.MemberProperty{JsiiProperty: "ipamResourceDiscoveryOwnerId", GoGetter: "IpamResourceDiscoveryOwnerId"},
			_jsii_.MemberProperty{JsiiProperty: "ipamResourceDiscoveryRegion", GoGetter: "IpamResourceDiscoveryRegion"},
			_jsii_.MemberProperty{JsiiProperty: "ipamResourceDiscoveryState", GoGetter: "IpamResourceDiscoveryState"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "regions", GoGetter: "Regions"},
			_jsii_.MemberProperty{JsiiProperty: "resource", GoGetter: "Resource"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_IpamResourceDiscovery{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkResource)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IIpamResourceDiscovery)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"cdk-extensions.ec2.IpamResourceDiscoveryAssociation",
		reflect.TypeOf((*IpamResourceDiscoveryAssociation)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "ipam", GoGetter: "Ipam"},
			_jsii_.MemberProperty{JsiiProperty: "ipamResourceDiscovery", GoGetter: "IpamResourceDiscovery"},
			_jsii_.MemberProperty{JsiiProperty: "ipamResourceDiscoveryAssociationArn", GoGetter: "IpamResourceDiscoveryAssociationArn"},
			_jsii_.MemberProperty{JsiiProperty: "ipamResourceDiscoveryAssociationId", GoGetter: "IpamResourceDiscoveryAssociationId"},
			_jsii_.MemberProperty{JsiiProperty: "ipamResourceDiscoveryAssociationIpamArn", GoGetter: "IpamResourceDiscoveryAssociationIpamArn"},
			_jsii_.MemberProperty{JsiiProperty: "ipamResourceDiscoveryAssociationIpamRegion", GoGetter: "IpamResourceDiscoveryAssociationIpamRegion"},
			_jsii_.MemberProperty{JsiiProperty: "ipamResourceDiscoveryAssociationIsDefault", GoGetter: "IpamResourceDiscoveryAssociationIsDefault"},
			_jsii_.MemberProperty{JsiiProperty: "ipamResourceDiscoveryAssociationOwnerId", GoGetter: "IpamResourceDiscoveryAssociationOwnerId"},
			_jsii_.MemberProperty{JsiiProperty: "ipamResourceDiscoveryAssociationResourceDiscoveryId", GoGetter: "IpamResourceDiscoveryAssociationResourceDiscoveryId"},
			_jsii_.MemberProperty{JsiiProperty: "ipamResourceDiscoveryAssociationResourceDiscoveryStatus", GoGetter: "IpamResourceDiscoveryAssociationResourceDiscoveryStatus"},
			_jsii_.MemberProperty{JsiiProperty: "ipamResourceDiscoveryAssociationState", GoGetter: "IpamResourceDiscoveryAssociationState"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "resource", GoGetter: "Resource"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_IpamResourceDiscoveryAssociation{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkResource)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IIpamResourceDiscoveryAssociation)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdk-extensions.ec2.IpamResourceDiscoveryAssociationProps",
		reflect.TypeOf((*IpamResourceDiscoveryAssociationProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk-extensions.ec2.IpamResourceDiscoveryAttributes",
		reflect.TypeOf((*IpamResourceDiscoveryAttributes)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk-extensions.ec2.IpamResourceDiscoveryProps",
		reflect.TypeOf((*IpamResourceDiscoveryProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk-extensions.ec2.IpamScope",
		reflect.TypeOf((*IpamScope)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addPool", GoMethod: "AddPool"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "description", GoGetter: "Description"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "ipam", GoGetter: "Ipam"},
			_jsii_.MemberProperty{JsiiProperty: "ipamScopeArn", GoGetter: "IpamScopeArn"},
			_jsii_.MemberProperty{JsiiProperty: "ipamScopeId", GoGetter: "IpamScopeId"},
			_jsii_.MemberProperty{JsiiProperty: "ipamScopeIpamArn", GoGetter: "IpamScopeIpamArn"},
			_jsii_.MemberProperty{JsiiProperty: "ipamScopeIsDefault", GoGetter: "IpamScopeIsDefault"},
			_jsii_.MemberProperty{JsiiProperty: "ipamScopePoolCount", GoGetter: "IpamScopePoolCount"},
			_jsii_.MemberProperty{JsiiProperty: "ipamScopeType", GoGetter: "IpamScopeType"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "resource", GoGetter: "Resource"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_IpamScope{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkResource)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IIpamScope)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdk-extensions.ec2.IpamScopeAttributes",
		reflect.TypeOf((*IpamScopeAttributes)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk-extensions.ec2.IpamScopeOptions",
		reflect.TypeOf((*IpamScopeOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk-extensions.ec2.IpamScopeProps",
		reflect.TypeOf((*IpamScopeProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk-extensions.ec2.Ipv4ConfigurationOptions",
		reflect.TypeOf((*Ipv4ConfigurationOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk-extensions.ec2.Ipv6ConfigurationOptions",
		reflect.TypeOf((*Ipv6ConfigurationOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk-extensions.ec2.LocalVpnEndpointConfiguration",
		reflect.TypeOf((*LocalVpnEndpointConfiguration)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk-extensions.ec2.NatProvider",
		reflect.TypeOf((*NatProvider)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_NatProvider{}
		},
	)
	_jsii_.RegisterStruct(
		"cdk-extensions.ec2.NetmaskLengthOptions",
		reflect.TypeOf((*NetmaskLengthOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk-extensions.ec2.PublicIpSource",
		reflect.TypeOf((*PublicIpSource)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
		},
		func() interface{} {
			return &jsiiProxy_PublicIpSource{}
		},
	)
	_jsii_.RegisterStruct(
		"cdk-extensions.ec2.RemoteVpnEndpointConfiguration",
		reflect.TypeOf((*RemoteVpnEndpointConfiguration)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk-extensions.ec2.ResolvedIpamAllocationConfiguration",
		reflect.TypeOf((*ResolvedIpamAllocationConfiguration)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk-extensions.ec2.ResolvedIpamPoolCidrConfiguration",
		reflect.TypeOf((*ResolvedIpamPoolCidrConfiguration)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk-extensions.ec2.SharingOptions",
		reflect.TypeOf((*SharingOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk-extensions.ec2.TieredSubnets",
		reflect.TypeOf((*TieredSubnets)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "allocateSubnetsCidr", GoMethod: "AllocateSubnetsCidr"},
			_jsii_.MemberMethod{JsiiMethod: "allocateVpcCidr", GoMethod: "AllocateVpcCidr"},
			_jsii_.MemberProperty{JsiiProperty: "provider", GoGetter: "Provider"},
			_jsii_.MemberProperty{JsiiProperty: "tierMask", GoGetter: "TierMask"},
		},
		func() interface{} {
			j := jsiiProxy_TieredSubnets{}
			_jsii_.InitJsiiProxy(&j.Type__awsec2IIpAddresses)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdk-extensions.ec2.TieredSubnetsOptions",
		reflect.TypeOf((*TieredSubnetsOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk-extensions.ec2.TransitGateway",
		reflect.TypeOf((*TransitGateway)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addCidrBlock", GoMethod: "AddCidrBlock"},
			_jsii_.MemberMethod{JsiiMethod: "addRouteTable", GoMethod: "AddRouteTable"},
			_jsii_.MemberMethod{JsiiMethod: "addVpn", GoMethod: "AddVpn"},
			_jsii_.MemberProperty{JsiiProperty: "amazonSideAsn", GoGetter: "AmazonSideAsn"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberMethod{JsiiMethod: "attachPeer", GoMethod: "AttachPeer"},
			_jsii_.MemberMethod{JsiiMethod: "attachVpc", GoMethod: "AttachVpc"},
			_jsii_.MemberProperty{JsiiProperty: "autoAcceptSharedAttachments", GoGetter: "AutoAcceptSharedAttachments"},
			_jsii_.MemberProperty{JsiiProperty: "defaultRouteTable", GoGetter: "DefaultRouteTable"},
			_jsii_.MemberProperty{JsiiProperty: "defaultRouteTableAssociation", GoGetter: "DefaultRouteTableAssociation"},
			_jsii_.MemberProperty{JsiiProperty: "defaultRouteTableId", GoGetter: "DefaultRouteTableId"},
			_jsii_.MemberProperty{JsiiProperty: "defaultRouteTablePropagation", GoGetter: "DefaultRouteTablePropagation"},
			_jsii_.MemberProperty{JsiiProperty: "description", GoGetter: "Description"},
			_jsii_.MemberProperty{JsiiProperty: "dnsSupport", GoGetter: "DnsSupport"},
			_jsii_.MemberMethod{JsiiMethod: "enableSharing", GoMethod: "EnableSharing"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "multicastSupport", GoGetter: "MulticastSupport"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "resource", GoGetter: "Resource"},
			_jsii_.MemberProperty{JsiiProperty: "resourceShare", GoGetter: "ResourceShare"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "transitGatewayArn", GoGetter: "TransitGatewayArn"},
			_jsii_.MemberProperty{JsiiProperty: "transitGatewayId", GoGetter: "TransitGatewayId"},
			_jsii_.MemberProperty{JsiiProperty: "vpnEcmpSupport", GoGetter: "VpnEcmpSupport"},
		},
		func() interface{} {
			j := jsiiProxy_TransitGateway{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkResource)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_ITransitGateway)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"cdk-extensions.ec2.TransitGatewayAttachment",
		reflect.TypeOf((*TransitGatewayAttachment)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addRoute", GoMethod: "AddRoute"},
			_jsii_.MemberProperty{JsiiProperty: "applianceModeSupport", GoGetter: "ApplianceModeSupport"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "dnsSupport", GoGetter: "DnsSupport"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "ipv6Support", GoGetter: "Ipv6Support"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "resource", GoGetter: "Resource"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "subnets", GoGetter: "Subnets"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "transitGateway", GoGetter: "TransitGateway"},
			_jsii_.MemberProperty{JsiiProperty: "transitGatewayAttachmentArn", GoGetter: "TransitGatewayAttachmentArn"},
			_jsii_.MemberProperty{JsiiProperty: "transitGatewayAttachmentId", GoGetter: "TransitGatewayAttachmentId"},
			_jsii_.MemberMethod{JsiiMethod: "translateBoolean", GoMethod: "TranslateBoolean"},
			_jsii_.MemberProperty{JsiiProperty: "vpc", GoGetter: "Vpc"},
		},
		func() interface{} {
			j := jsiiProxy_TransitGatewayAttachment{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_TransitGatewayAttachmentResource)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"cdk-extensions.ec2.TransitGatewayAttachmentBase",
		reflect.TypeOf((*TransitGatewayAttachmentBase)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addRoute", GoMethod: "AddRoute"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "transitGatewayAttachmentArn", GoGetter: "TransitGatewayAttachmentArn"},
			_jsii_.MemberProperty{JsiiProperty: "transitGatewayAttachmentId", GoGetter: "TransitGatewayAttachmentId"},
		},
		func() interface{} {
			j := jsiiProxy_TransitGatewayAttachmentBase{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkResource)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_ITransitGatewayAttachment)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdk-extensions.ec2.TransitGatewayAttachmentProps",
		reflect.TypeOf((*TransitGatewayAttachmentProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk-extensions.ec2.TransitGatewayAttachmentResource",
		reflect.TypeOf((*TransitGatewayAttachmentResource)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addRoute", GoMethod: "AddRoute"},
			_jsii_.MemberProperty{JsiiProperty: "applianceModeSupport", GoGetter: "ApplianceModeSupport"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "dnsSupport", GoGetter: "DnsSupport"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "ipv6Support", GoGetter: "Ipv6Support"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "subnets", GoGetter: "Subnets"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "transitGateway", GoGetter: "TransitGateway"},
			_jsii_.MemberProperty{JsiiProperty: "transitGatewayAttachmentArn", GoGetter: "TransitGatewayAttachmentArn"},
			_jsii_.MemberProperty{JsiiProperty: "transitGatewayAttachmentId", GoGetter: "TransitGatewayAttachmentId"},
			_jsii_.MemberMethod{JsiiMethod: "translateBoolean", GoMethod: "TranslateBoolean"},
			_jsii_.MemberProperty{JsiiProperty: "vpc", GoGetter: "Vpc"},
		},
		func() interface{} {
			j := jsiiProxy_TransitGatewayAttachmentResource{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_TransitGatewayAttachmentBase)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdk-extensions.ec2.TransitGatewayAttachmentResourceProps",
		reflect.TypeOf((*TransitGatewayAttachmentResourceProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk-extensions.ec2.TransitGatewayLocalVpnEndpoint",
		reflect.TypeOf((*TransitGatewayLocalVpnEndpoint)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
			_jsii_.MemberProperty{JsiiProperty: "transitGateway", GoGetter: "TransitGateway"},
		},
		func() interface{} {
			j := jsiiProxy_TransitGatewayLocalVpnEndpoint{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_ILocalVpnEndpoint)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"cdk-extensions.ec2.TransitGatewayNatProvider",
		reflect.TypeOf((*TransitGatewayNatProvider)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "configuredGateways", GoGetter: "ConfiguredGateways"},
			_jsii_.MemberMethod{JsiiMethod: "configureNat", GoMethod: "ConfigureNat"},
			_jsii_.MemberMethod{JsiiMethod: "configureSubnet", GoMethod: "ConfigureSubnet"},
			_jsii_.MemberProperty{JsiiProperty: "transitGateway", GoGetter: "TransitGateway"},
			_jsii_.MemberProperty{JsiiProperty: "transitGatewayAttachment", GoGetter: "TransitGatewayAttachment"},
		},
		func() interface{} {
			j := jsiiProxy_TransitGatewayNatProvider{}
			_jsii_.InitJsiiProxy(&j.Type__awsec2NatProvider)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdk-extensions.ec2.TransitGatewayNatProviderOptions",
		reflect.TypeOf((*TransitGatewayNatProviderOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk-extensions.ec2.TransitGatewayPeeringAttachment",
		reflect.TypeOf((*TransitGatewayPeeringAttachment)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addRoute", GoMethod: "AddRoute"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "localTransitGateway", GoGetter: "LocalTransitGateway"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "peerAccountId", GoGetter: "PeerAccountId"},
			_jsii_.MemberProperty{JsiiProperty: "peerRegion", GoGetter: "PeerRegion"},
			_jsii_.MemberProperty{JsiiProperty: "peerTransitGateway", GoGetter: "PeerTransitGateway"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "resource", GoGetter: "Resource"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "transitGatewayAttachmentArn", GoGetter: "TransitGatewayAttachmentArn"},
			_jsii_.MemberProperty{JsiiProperty: "transitGatewayAttachmentCreationTime", GoGetter: "TransitGatewayAttachmentCreationTime"},
			_jsii_.MemberProperty{JsiiProperty: "transitGatewayAttachmentId", GoGetter: "TransitGatewayAttachmentId"},
			_jsii_.MemberProperty{JsiiProperty: "transitGatewayAttachmentState", GoGetter: "TransitGatewayAttachmentState"},
			_jsii_.MemberProperty{JsiiProperty: "transitGatewayAttachmentStatus", GoGetter: "TransitGatewayAttachmentStatus"},
			_jsii_.MemberProperty{JsiiProperty: "transitGatewayAttachmentStatusCode", GoGetter: "TransitGatewayAttachmentStatusCode"},
			_jsii_.MemberProperty{JsiiProperty: "transitGatewayAttachmentStatusMessage", GoGetter: "TransitGatewayAttachmentStatusMessage"},
		},
		func() interface{} {
			j := jsiiProxy_TransitGatewayPeeringAttachment{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_TransitGatewayAttachmentBase)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_ITransitGatewayPeeringAttachment)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdk-extensions.ec2.TransitGatewayPeeringAttachmentImportAttributes",
		reflect.TypeOf((*TransitGatewayPeeringAttachmentImportAttributes)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk-extensions.ec2.TransitGatewayPeeringAttachmentOptions",
		reflect.TypeOf((*TransitGatewayPeeringAttachmentOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk-extensions.ec2.TransitGatewayPeeringAttachmentProps",
		reflect.TypeOf((*TransitGatewayPeeringAttachmentProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk-extensions.ec2.TransitGatewayProps",
		reflect.TypeOf((*TransitGatewayProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk-extensions.ec2.TransitGatewayRoute",
		reflect.TypeOf((*TransitGatewayRoute)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "attachment", GoGetter: "Attachment"},
			_jsii_.MemberProperty{JsiiProperty: "blackhole", GoGetter: "Blackhole"},
			_jsii_.MemberProperty{JsiiProperty: "cidr", GoGetter: "Cidr"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "resource", GoGetter: "Resource"},
			_jsii_.MemberProperty{JsiiProperty: "routeTable", GoGetter: "RouteTable"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "transitGatewayRouteId", GoGetter: "TransitGatewayRouteId"},
		},
		func() interface{} {
			j := jsiiProxy_TransitGatewayRoute{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdk-extensions.ec2.TransitGatewayRouteOptions",
		reflect.TypeOf((*TransitGatewayRouteOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk-extensions.ec2.TransitGatewayRouteProps",
		reflect.TypeOf((*TransitGatewayRouteProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk-extensions.ec2.TransitGatewayRouteTable",
		reflect.TypeOf((*TransitGatewayRouteTable)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addRoute", GoMethod: "AddRoute"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "resource", GoGetter: "Resource"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "transitGateway", GoGetter: "TransitGateway"},
			_jsii_.MemberProperty{JsiiProperty: "transitGatewayRouteTableArn", GoGetter: "TransitGatewayRouteTableArn"},
			_jsii_.MemberProperty{JsiiProperty: "transitGatewayRouteTableId", GoGetter: "TransitGatewayRouteTableId"},
		},
		func() interface{} {
			j := jsiiProxy_TransitGatewayRouteTable{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkResource)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_ITransitGatewayRouteTable)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdk-extensions.ec2.TransitGatewayRouteTableOptions",
		reflect.TypeOf((*TransitGatewayRouteTableOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk-extensions.ec2.TransitGatewayRouteTableProps",
		reflect.TypeOf((*TransitGatewayRouteTableProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk-extensions.ec2.TunnelOptions",
		reflect.TypeOf((*TunnelOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk-extensions.ec2.VpcAttachmentOptions",
		reflect.TypeOf((*VpcAttachmentOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk-extensions.ec2.VpnAttachmentOptions",
		reflect.TypeOf((*VpnAttachmentOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk-extensions.ec2.VpnConnection",
		reflect.TypeOf((*VpnConnection)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addTunnelConfiguration", GoMethod: "AddTunnelConfiguration"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "connectionType", GoGetter: "ConnectionType"},
			_jsii_.MemberProperty{JsiiProperty: "customerGatewayAsn", GoGetter: "CustomerGatewayAsn"},
			_jsii_.MemberProperty{JsiiProperty: "customerGatewayId", GoGetter: "CustomerGatewayId"},
			_jsii_.MemberProperty{JsiiProperty: "customerGatewayIp", GoGetter: "CustomerGatewayIp"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "localEndpoint", GoGetter: "LocalEndpoint"},
			_jsii_.MemberMethod{JsiiMethod: "metric", GoMethod: "Metric"},
			_jsii_.MemberMethod{JsiiMethod: "metricTunnelDataIn", GoMethod: "MetricTunnelDataIn"},
			_jsii_.MemberMethod{JsiiMethod: "metricTunnelDataOut", GoMethod: "MetricTunnelDataOut"},
			_jsii_.MemberMethod{JsiiMethod: "metricTunnelState", GoMethod: "MetricTunnelState"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "remoteEndpoint", GoGetter: "RemoteEndpoint"},
			_jsii_.MemberProperty{JsiiProperty: "resource", GoGetter: "Resource"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "staticRoutesOnly", GoGetter: "StaticRoutesOnly"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "tunnelConfigurations", GoGetter: "TunnelConfigurations"},
			_jsii_.MemberProperty{JsiiProperty: "vpnId", GoGetter: "VpnId"},
		},
		func() interface{} {
			j := jsiiProxy_VpnConnection{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkResource)
			_jsii_.InitJsiiProxy(&j.Type__awsec2IVpnConnection)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"cdk-extensions.ec2.VpnConnectionLocalEndpoint",
		reflect.TypeOf((*VpnConnectionLocalEndpoint)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_VpnConnectionLocalEndpoint{}
		},
	)
	_jsii_.RegisterStruct(
		"cdk-extensions.ec2.VpnConnectionProps",
		reflect.TypeOf((*VpnConnectionProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk-extensions.ec2.VpnConnectionRemoteEndpoint",
		reflect.TypeOf((*VpnConnectionRemoteEndpoint)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_VpnConnectionRemoteEndpoint{}
		},
	)
	_jsii_.RegisterClass(
		"cdk-extensions.ec2.VpnConnectionType",
		reflect.TypeOf((*VpnConnectionType)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
		},
		func() interface{} {
			return &jsiiProxy_VpnConnectionType{}
		},
	)
	_jsii_.RegisterClass(
		"cdk-extensions.ec2.VpnGatewayLocalVpnEndpoint",
		reflect.TypeOf((*VpnGatewayLocalVpnEndpoint)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
			_jsii_.MemberProperty{JsiiProperty: "vpnGateway", GoGetter: "VpnGateway"},
		},
		func() interface{} {
			j := jsiiProxy_VpnGatewayLocalVpnEndpoint{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_ILocalVpnEndpoint)
			return &j
		},
	)
}
