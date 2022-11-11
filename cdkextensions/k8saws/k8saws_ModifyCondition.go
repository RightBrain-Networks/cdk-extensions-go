package k8saws

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/vibe-io/cdk-extensions-go/cdkextensions/jsii"
)

type ModifyCondition interface {
	// Collection of arguments that apply to the condition.
	Args() *[]*string
	// The name of the condition being evaluated.
	Condition() *string
	// Gets a string representation of the arguments of this condition for use in a Fluent Bit plugin field.
	//
	// Returns: A fluent bit value string.
	ToString() *string
}

// The jsii proxy struct for ModifyCondition
type jsiiProxy_ModifyCondition struct {
	_ byte // padding
}

func (j *jsiiProxy_ModifyCondition) Args() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"args",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ModifyCondition) Condition() *string {
	var returns *string
	_jsii_.Get(
		j,
		"condition",
		&returns,
	)
	return returns
}


// Condition that returns true if any key matches a specified regular expression.
//
// Returns: A ModifyCondition object representing the condition.
func ModifyCondition_AKeyMatches(regex *string) ModifyCondition {
	_init_.Initialize()

	if err := validateModifyCondition_AKeyMatchesParameters(regex); err != nil {
		panic(err)
	}
	var returns ModifyCondition

	_jsii_.StaticInvoke(
		"cdk-extensions.k8s_aws.ModifyCondition",
		"aKeyMatches",
		[]interface{}{regex},
		&returns,
	)

	return returns
}

// Condition that returns true if a specified key does not exist.
//
// Returns: A ModifyCondition object representing the condition.
func ModifyCondition_KeyDoesNotExists(key *string) ModifyCondition {
	_init_.Initialize()

	if err := validateModifyCondition_KeyDoesNotExistsParameters(key); err != nil {
		panic(err)
	}
	var returns ModifyCondition

	_jsii_.StaticInvoke(
		"cdk-extensions.k8s_aws.ModifyCondition",
		"keyDoesNotExists",
		[]interface{}{key},
		&returns,
	)

	return returns
}

// Condition that returns true if a specified key exists.
//
// Returns: A ModifyCondition object representing the condition.
func ModifyCondition_KeyExists(key *string) ModifyCondition {
	_init_.Initialize()

	if err := validateModifyCondition_KeyExistsParameters(key); err != nil {
		panic(err)
	}
	var returns ModifyCondition

	_jsii_.StaticInvoke(
		"cdk-extensions.k8s_aws.ModifyCondition",
		"keyExists",
		[]interface{}{key},
		&returns,
	)

	return returns
}

// Condition that returns true if a specified key exists and its value does not match the specified value.
//
// Returns: A ModifyCondition object representing the condition.
func ModifyCondition_KeyValueDoesNotEqual(key *string, value *string) ModifyCondition {
	_init_.Initialize()

	if err := validateModifyCondition_KeyValueDoesNotEqualParameters(key, value); err != nil {
		panic(err)
	}
	var returns ModifyCondition

	_jsii_.StaticInvoke(
		"cdk-extensions.k8s_aws.ModifyCondition",
		"keyValueDoesNotEqual",
		[]interface{}{key, value},
		&returns,
	)

	return returns
}

// Condition that returns true if a specified key exists and its value does not match the specified regular expression.
//
// Returns: A ModifyCondition object representing the condition.
func ModifyCondition_KeyValueDoesNotMatch(key *string, value *string) ModifyCondition {
	_init_.Initialize()

	if err := validateModifyCondition_KeyValueDoesNotMatchParameters(key, value); err != nil {
		panic(err)
	}
	var returns ModifyCondition

	_jsii_.StaticInvoke(
		"cdk-extensions.k8s_aws.ModifyCondition",
		"keyValueDoesNotMatch",
		[]interface{}{key, value},
		&returns,
	)

	return returns
}

// Condition that returns true if a specified key exists and its value matches the specified value.
//
// Returns: A ModifyCondition object representing the condition.
func ModifyCondition_KeyValueEquals(key *string, value *string) ModifyCondition {
	_init_.Initialize()

	if err := validateModifyCondition_KeyValueEqualsParameters(key, value); err != nil {
		panic(err)
	}
	var returns ModifyCondition

	_jsii_.StaticInvoke(
		"cdk-extensions.k8s_aws.ModifyCondition",
		"keyValueEquals",
		[]interface{}{key, value},
		&returns,
	)

	return returns
}

// Condition that returns true if a specified key exists and its value matches the specified regular expression.
//
// Returns: A ModifyCondition object representing the condition.
func ModifyCondition_KeyValueMatches(key *string, value *string) ModifyCondition {
	_init_.Initialize()

	if err := validateModifyCondition_KeyValueMatchesParameters(key, value); err != nil {
		panic(err)
	}
	var returns ModifyCondition

	_jsii_.StaticInvoke(
		"cdk-extensions.k8s_aws.ModifyCondition",
		"keyValueMatches",
		[]interface{}{key, value},
		&returns,
	)

	return returns
}

// Condition that returns true if all keys matching a specified regular expression have values that do not match another regular expression.
//
// Returns: A ModifyCondition object representing the condition.
func ModifyCondition_MatchingKeysDoNotHaveMatchingValues(key *string, value *string) ModifyCondition {
	_init_.Initialize()

	if err := validateModifyCondition_MatchingKeysDoNotHaveMatchingValuesParameters(key, value); err != nil {
		panic(err)
	}
	var returns ModifyCondition

	_jsii_.StaticInvoke(
		"cdk-extensions.k8s_aws.ModifyCondition",
		"matchingKeysDoNotHaveMatchingValues",
		[]interface{}{key, value},
		&returns,
	)

	return returns
}

// Condition that returns true if all keys matching a specified regular expression have values that match another regular expression.
//
// Returns: A ModifyCondition object representing the condition.
func ModifyCondition_MatchingKeysHaveMatchingValues(key *string, value *string) ModifyCondition {
	_init_.Initialize()

	if err := validateModifyCondition_MatchingKeysHaveMatchingValuesParameters(key, value); err != nil {
		panic(err)
	}
	var returns ModifyCondition

	_jsii_.StaticInvoke(
		"cdk-extensions.k8s_aws.ModifyCondition",
		"matchingKeysHaveMatchingValues",
		[]interface{}{key, value},
		&returns,
	)

	return returns
}

// Condition that returns true if no key matches a specified regular expression.
//
// Returns: A ModifyCondition object representing the condition.
func ModifyCondition_NoKeyMatches(regex *string) ModifyCondition {
	_init_.Initialize()

	if err := validateModifyCondition_NoKeyMatchesParameters(regex); err != nil {
		panic(err)
	}
	var returns ModifyCondition

	_jsii_.StaticInvoke(
		"cdk-extensions.k8s_aws.ModifyCondition",
		"noKeyMatches",
		[]interface{}{regex},
		&returns,
	)

	return returns
}

// An escape hatch method that allows fo defining custom conditions to be evaluated by the modify Fluent Bit filter plugin.
//
// Returns: A ModifyCondition object representing the options provided.
func ModifyCondition_Of(condition *string, args *[]*string) ModifyCondition {
	_init_.Initialize()

	if err := validateModifyCondition_OfParameters(condition, args); err != nil {
		panic(err)
	}
	var returns ModifyCondition

	_jsii_.StaticInvoke(
		"cdk-extensions.k8s_aws.ModifyCondition",
		"of",
		[]interface{}{condition, args},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ModifyCondition) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

