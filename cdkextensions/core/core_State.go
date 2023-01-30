package core

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/vibe-io/cdk-extensions-go/cdkextensions/jsii"

	"github.com/aws/constructs-go/constructs/v10"
)

// Provides a means of storing and retrieving arbitrary data that can be associated with a construct.
type State interface {
	// Gets the value of a key from state.
	//
	// Returns: The value of the requested key or `defaultValue` if the requested
	// key was not found.
	Get(key *string, defaultValue interface{}) interface{}
	// Sets the value of a key in state.
	//
	// Returns: The previous value for the key that was set (if one exists).
	Set(key *string, value interface{}) interface{}
}

// The jsii proxy struct for State
type jsiiProxy_State struct {
	_ byte // padding
}

// Gets an object that allows for interacting with the stored state of a construct.
//
// Returns: An object that provides a means for interacting with the stored
// state of the construct.
func State_Of(scope constructs.IConstruct) State {
	_init_.Initialize()

	if err := validateState_OfParameters(scope); err != nil {
		panic(err)
	}
	var returns State

	_jsii_.StaticInvoke(
		"cdk-extensions.core.State",
		"of",
		[]interface{}{scope},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_State) Get(key *string, defaultValue interface{}) interface{} {
	if err := s.validateGetParameters(key); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		s,
		"get",
		[]interface{}{key, defaultValue},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_State) Set(key *string, value interface{}) interface{} {
	if err := s.validateSetParameters(key, value); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		s,
		"set",
		[]interface{}{key, value},
		&returns,
	)

	return returns
}

