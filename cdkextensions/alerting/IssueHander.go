package alerting

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/vibe-io/cdk-extensions-go/cdkextensions/jsii"

	"github.com/aws/constructs-go/constructs/v10"
)

type IssueHander interface {
}

// The jsii proxy struct for IssueHander
type jsiiProxy_IssueHander struct {
	_ byte // padding
}

func NewIssueHander() IssueHander {
	_init_.Initialize()

	j := jsiiProxy_IssueHander{}

	_jsii_.Create(
		"cdk-extensions.alerting.IssueHander",
		nil, // no parameters
		&j,
	)

	return &j
}

func NewIssueHander_Override(i IssueHander) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk-extensions.alerting.IssueHander",
		nil, // no parameters
		i,
	)
}

func IssueHander_Discord(scope constructs.IConstruct, id *string, props *DiscordProps) Discord {
	_init_.Initialize()

	if err := validateIssueHander_DiscordParameters(scope, id, props); err != nil {
		panic(err)
	}
	var returns Discord

	_jsii_.StaticInvoke(
		"cdk-extensions.alerting.IssueHander",
		"discord",
		[]interface{}{scope, id, props},
		&returns,
	)

	return returns
}

func IssueHander_JiraTicket(scope constructs.IConstruct, id *string, props *JiraTicketProps) JiraTicket {
	_init_.Initialize()

	if err := validateIssueHander_JiraTicketParameters(scope, id, props); err != nil {
		panic(err)
	}
	var returns JiraTicket

	_jsii_.StaticInvoke(
		"cdk-extensions.alerting.IssueHander",
		"jiraTicket",
		[]interface{}{scope, id, props},
		&returns,
	)

	return returns
}

