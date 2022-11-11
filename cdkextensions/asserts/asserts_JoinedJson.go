package asserts

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/vibe-io/cdk-extensions-go/cdkextensions/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2/assertions"
	"github.com/vibe-io/cdk-extensions-go/cdkextensions/asserts/internal"
)

type JoinedJson interface {
	assertions.Matcher
	// A name for the matcher.
	//
	// This is collected as part of the result and may be presented to the user.
	Name() *string
	// Test whether a target matches the provided pattern.
	//
	// Every Matcher must implement this method.
	// This method will be invoked by the assertions framework. Do not call this method directly.
	Test(actual interface{}) assertions.MatchResult
}

// The jsii proxy struct for JoinedJson
type jsiiProxy_JoinedJson struct {
	internal.Type__assertionsMatcher
}

func (j *jsiiProxy_JoinedJson) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}


func NewJoinedJson(pattern interface{}) JoinedJson {
	_init_.Initialize()

	if err := validateNewJoinedJsonParameters(pattern); err != nil {
		panic(err)
	}
	j := jsiiProxy_JoinedJson{}

	_jsii_.Create(
		"cdk-extensions.asserts.JoinedJson",
		[]interface{}{pattern},
		&j,
	)

	return &j
}

func NewJoinedJson_Override(j JoinedJson, pattern interface{}) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk-extensions.asserts.JoinedJson",
		[]interface{}{pattern},
		j,
	)
}

// Check whether the provided object is a subtype of the `IMatcher`.
func JoinedJson_IsMatcher(x interface{}) *bool {
	_init_.Initialize()

	if err := validateJoinedJson_IsMatcherParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"cdk-extensions.asserts.JoinedJson",
		"isMatcher",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JoinedJson) Test(actual interface{}) assertions.MatchResult {
	if err := j.validateTestParameters(actual); err != nil {
		panic(err)
	}
	var returns assertions.MatchResult

	_jsii_.Invoke(
		j,
		"test",
		[]interface{}{actual},
		&returns,
	)

	return returns
}

