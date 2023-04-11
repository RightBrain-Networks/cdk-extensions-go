package core

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/vibe-io/cdk-extensions-go/cdkextensions/jsii"

	"github.com/aws/constructs-go/constructs/v10"
)

type SecretReference interface {
	ValueForScope(scope constructs.IConstruct) *string
}

// The jsii proxy struct for SecretReference
type jsiiProxy_SecretReference struct {
	_ byte // padding
}

func SecretReference_String(scope constructs.IConstruct, key *string, value *string) *string {
	_init_.Initialize()

	if err := validateSecretReference_StringParameters(scope, key, value); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.StaticInvoke(
		"cdk-extensions.core.SecretReference",
		"string",
		[]interface{}{scope, key, value},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SecretReference) ValueForScope(scope constructs.IConstruct) *string {
	if err := s.validateValueForScopeParameters(scope); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		s,
		"valueForScope",
		[]interface{}{scope},
		&returns,
	)

	return returns
}

