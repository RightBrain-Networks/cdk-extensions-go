package alerting

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/vibe-io/cdk-extensions-go/cdkextensions/jsii"

	"github.com/aws/constructs-go/constructs/v10"
)

type IssueParser interface {
}

// The jsii proxy struct for IssueParser
type jsiiProxy_IssueParser struct {
	_ byte // padding
}

func NewIssueParser() IssueParser {
	_init_.Initialize()

	j := jsiiProxy_IssueParser{}

	_jsii_.Create(
		"cdk-extensions.alerting.IssueParser",
		nil, // no parameters
		&j,
	)

	return &j
}

func NewIssueParser_Override(i IssueParser) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk-extensions.alerting.IssueParser",
		nil, // no parameters
		i,
	)
}

func IssueParser_ConfigComplianceChange(scope constructs.IConstruct, id *string, props *ConfigComplianceChangeProps) ConfigComplianceChange {
	_init_.Initialize()

	if err := validateIssueParser_ConfigComplianceChangeParameters(scope, id, props); err != nil {
		panic(err)
	}
	var returns ConfigComplianceChange

	_jsii_.StaticInvoke(
		"cdk-extensions.alerting.IssueParser",
		"configComplianceChange",
		[]interface{}{scope, id, props},
		&returns,
	)

	return returns
}

func IssueParser_EcrScanFinding(scope constructs.IConstruct, id *string, props *EcrScanFindingProps) EcrScanFinding {
	_init_.Initialize()

	if err := validateIssueParser_EcrScanFindingParameters(scope, id, props); err != nil {
		panic(err)
	}
	var returns EcrScanFinding

	_jsii_.StaticInvoke(
		"cdk-extensions.alerting.IssueParser",
		"ecrScanFinding",
		[]interface{}{scope, id, props},
		&returns,
	)

	return returns
}

func IssueParser_GuardDutyFinding(scope constructs.IConstruct, id *string, props *GuardDutyFindingProps) GuardDutyFinding {
	_init_.Initialize()

	if err := validateIssueParser_GuardDutyFindingParameters(scope, id, props); err != nil {
		panic(err)
	}
	var returns GuardDutyFinding

	_jsii_.StaticInvoke(
		"cdk-extensions.alerting.IssueParser",
		"guardDutyFinding",
		[]interface{}{scope, id, props},
		&returns,
	)

	return returns
}

func IssueParser_InspectorFinding(scope constructs.IConstruct, id *string, props *InspectorFindingProps) InspectorFinding {
	_init_.Initialize()

	if err := validateIssueParser_InspectorFindingParameters(scope, id, props); err != nil {
		panic(err)
	}
	var returns InspectorFinding

	_jsii_.StaticInvoke(
		"cdk-extensions.alerting.IssueParser",
		"inspectorFinding",
		[]interface{}{scope, id, props},
		&returns,
	)

	return returns
}

func IssueParser_OpenSearchEvent(scope constructs.IConstruct, id *string, props *OpenSearchEventProps) OpenSearchEvent {
	_init_.Initialize()

	if err := validateIssueParser_OpenSearchEventParameters(scope, id, props); err != nil {
		panic(err)
	}
	var returns OpenSearchEvent

	_jsii_.StaticInvoke(
		"cdk-extensions.alerting.IssueParser",
		"openSearchEvent",
		[]interface{}{scope, id, props},
		&returns,
	)

	return returns
}

func IssueParser_SecurityHubFinding(scope constructs.IConstruct, id *string, props *SecurityHubFindingProps) SecurityHubFinding {
	_init_.Initialize()

	if err := validateIssueParser_SecurityHubFindingParameters(scope, id, props); err != nil {
		panic(err)
	}
	var returns SecurityHubFinding

	_jsii_.StaticInvoke(
		"cdk-extensions.alerting.IssueParser",
		"securityHubFinding",
		[]interface{}{scope, id, props},
		&returns,
	)

	return returns
}

