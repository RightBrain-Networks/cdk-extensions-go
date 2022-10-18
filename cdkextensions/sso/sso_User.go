package sso

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/vibe-io/cdk-extensions-go/cdkextensions/jsii"

	"github.com/aws/constructs-go/constructs/v10"
)

type User interface {
}

// The jsii proxy struct for User
type jsiiProxy_User struct {
	_ byte // padding
}

func NewUser() User {
	_init_.Initialize()

	j := jsiiProxy_User{}

	_jsii_.Create(
		"cdk-extensions.sso.User",
		nil, // no parameters
		&j,
	)

	return &j
}

func NewUser_Override(u User) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk-extensions.sso.User",
		nil, // no parameters
		u,
	)
}

func User_FromUserId(scope constructs.IConstruct, id *string, userId *string) IUser {
	_init_.Initialize()

	if err := validateUser_FromUserIdParameters(scope, id, userId); err != nil {
		panic(err)
	}
	var returns IUser

	_jsii_.StaticInvoke(
		"cdk-extensions.sso.User",
		"fromUserId",
		[]interface{}{scope, id, userId},
		&returns,
	)

	return returns
}

