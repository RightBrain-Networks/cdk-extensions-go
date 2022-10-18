package sso

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/vibe-io/cdk-extensions-go/cdkextensions/jsii"

	"github.com/aws/constructs-go/constructs/v10"
)

type Group interface {
}

// The jsii proxy struct for Group
type jsiiProxy_Group struct {
	_ byte // padding
}

func NewGroup() Group {
	_init_.Initialize()

	j := jsiiProxy_Group{}

	_jsii_.Create(
		"cdk-extensions.sso.Group",
		nil, // no parameters
		&j,
	)

	return &j
}

func NewGroup_Override(g Group) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk-extensions.sso.Group",
		nil, // no parameters
		g,
	)
}

func Group_FromGroupId(scope constructs.IConstruct, id *string, groupId *string) IGroup {
	_init_.Initialize()

	if err := validateGroup_FromGroupIdParameters(scope, id, groupId); err != nil {
		panic(err)
	}
	var returns IGroup

	_jsii_.StaticInvoke(
		"cdk-extensions.sso.Group",
		"fromGroupId",
		[]interface{}{scope, id, groupId},
		&returns,
	)

	return returns
}

