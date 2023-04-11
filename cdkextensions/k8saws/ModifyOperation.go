package k8saws

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/vibe-io/cdk-extensions-go/cdkextensions/jsii"
)

type ModifyOperation interface {
	// Collection of arguments that apply to the operation.
	Args() *[]*string
	// The name of the operation being performed.
	Operation() *string
	// Gets a string representation of the arguments of this operation for use in a Fluent Bit plugin field.
	//
	// Returns: A fluent bit value string.
	ToString() *string
}

// The jsii proxy struct for ModifyOperation
type jsiiProxy_ModifyOperation struct {
	_ byte // padding
}

func (j *jsiiProxy_ModifyOperation) Args() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"args",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ModifyOperation) Operation() *string {
	var returns *string
	_jsii_.Get(
		j,
		"operation",
		&returns,
	)
	return returns
}


// Sets a field in the output to a specific value.
//
// If a field with the same name already exists it will be kept as is.
//
// Returns: A ModifyOperation object representing the add operation.
func ModifyOperation_Add(key *string, value *string) ModifyOperation {
	_init_.Initialize()

	if err := validateModifyOperation_AddParameters(key, value); err != nil {
		panic(err)
	}
	var returns ModifyOperation

	_jsii_.StaticInvoke(
		"cdk-extensions.k8s_aws.ModifyOperation",
		"add",
		[]interface{}{key, value},
		&returns,
	)

	return returns
}

// Copies a field from the input to a field with a new name if the field exists and a field with the new name does not exist.
//
// If a field with the new name already exists it is overwritten.
//
// Returns: A ModifyOperation object representing the copy operation.
func ModifyOperation_Copy(originalKey *string, newKey *string) ModifyOperation {
	_init_.Initialize()

	if err := validateModifyOperation_CopyParameters(originalKey, newKey); err != nil {
		panic(err)
	}
	var returns ModifyOperation

	_jsii_.StaticInvoke(
		"cdk-extensions.k8s_aws.ModifyOperation",
		"copy",
		[]interface{}{originalKey, newKey},
		&returns,
	)

	return returns
}

// Copies a field from the input to a field with a new name if the field exists and a field with the new name does not exist.
//
// Returns: A ModifyOperation object representing the copy operation.
func ModifyOperation_HardCopy(originalKey *string, newKey *string) ModifyOperation {
	_init_.Initialize()

	if err := validateModifyOperation_HardCopyParameters(originalKey, newKey); err != nil {
		panic(err)
	}
	var returns ModifyOperation

	_jsii_.StaticInvoke(
		"cdk-extensions.k8s_aws.ModifyOperation",
		"hardCopy",
		[]interface{}{originalKey, newKey},
		&returns,
	)

	return returns
}

// Renames a field from the input if the field exists.
//
// If a field with the desired name already exists it is overwritten.
//
// Returns: A ModifyOperation object representing the rename operation.
func ModifyOperation_HardRename(originalKey *string, renamedKey *string) ModifyOperation {
	_init_.Initialize()

	if err := validateModifyOperation_HardRenameParameters(originalKey, renamedKey); err != nil {
		panic(err)
	}
	var returns ModifyOperation

	_jsii_.StaticInvoke(
		"cdk-extensions.k8s_aws.ModifyOperation",
		"hardRename",
		[]interface{}{originalKey, renamedKey},
		&returns,
	)

	return returns
}

// Moves fiels matching the given wildcard key to the end of the message.
//
// Returns: A ModifyOperation object representing the move operation.
func ModifyOperation_MoveToEnd(key *string) ModifyOperation {
	_init_.Initialize()

	if err := validateModifyOperation_MoveToEndParameters(key); err != nil {
		panic(err)
	}
	var returns ModifyOperation

	_jsii_.StaticInvoke(
		"cdk-extensions.k8s_aws.ModifyOperation",
		"moveToEnd",
		[]interface{}{key},
		&returns,
	)

	return returns
}

// Moves fiels matching the given wildcard key to the start of the message.
//
// Returns: A ModifyOperation object representing the move operation.
func ModifyOperation_MoveToStart(key *string) ModifyOperation {
	_init_.Initialize()

	if err := validateModifyOperation_MoveToStartParameters(key); err != nil {
		panic(err)
	}
	var returns ModifyOperation

	_jsii_.StaticInvoke(
		"cdk-extensions.k8s_aws.ModifyOperation",
		"moveToStart",
		[]interface{}{key},
		&returns,
	)

	return returns
}

// An escape hatch method that allows fo defining custom operations to be performed by the modify Fluent Bit filter plugin.
//
// Returns: A ModifyOperation object representing the options provided.
func ModifyOperation_Of(operation *string, args *[]*string) ModifyOperation {
	_init_.Initialize()

	if err := validateModifyOperation_OfParameters(operation, args); err != nil {
		panic(err)
	}
	var returns ModifyOperation

	_jsii_.StaticInvoke(
		"cdk-extensions.k8s_aws.ModifyOperation",
		"of",
		[]interface{}{operation, args},
		&returns,
	)

	return returns
}

// Removes a field in the output with a specific key.
//
// Returns: A ModifyOperation object representing the remove operation.
func ModifyOperation_Remove(key *string) ModifyOperation {
	_init_.Initialize()

	if err := validateModifyOperation_RemoveParameters(key); err != nil {
		panic(err)
	}
	var returns ModifyOperation

	_jsii_.StaticInvoke(
		"cdk-extensions.k8s_aws.ModifyOperation",
		"remove",
		[]interface{}{key},
		&returns,
	)

	return returns
}

// Removes all fields in the output matching the regular expression.
//
// Returns: A ModifyOperation object representing the remove operation.
func ModifyOperation_RemoveRegex(regex *string) ModifyOperation {
	_init_.Initialize()

	if err := validateModifyOperation_RemoveRegexParameters(regex); err != nil {
		panic(err)
	}
	var returns ModifyOperation

	_jsii_.StaticInvoke(
		"cdk-extensions.k8s_aws.ModifyOperation",
		"removeRegex",
		[]interface{}{regex},
		&returns,
	)

	return returns
}

// Removes all fields in the output matching the wildcard key.
//
// Returns: A ModifyOperation object representing the remove operation.
func ModifyOperation_RemoveWildcard(key *string) ModifyOperation {
	_init_.Initialize()

	if err := validateModifyOperation_RemoveWildcardParameters(key); err != nil {
		panic(err)
	}
	var returns ModifyOperation

	_jsii_.StaticInvoke(
		"cdk-extensions.k8s_aws.ModifyOperation",
		"removeWildcard",
		[]interface{}{key},
		&returns,
	)

	return returns
}

// Renames a field from the input if the field exists and a field with the new name does not exist.
//
// Returns: A ModifyOperation object representing the rename operation.
func ModifyOperation_Rename(originalKey *string, renamedKey *string) ModifyOperation {
	_init_.Initialize()

	if err := validateModifyOperation_RenameParameters(originalKey, renamedKey); err != nil {
		panic(err)
	}
	var returns ModifyOperation

	_jsii_.StaticInvoke(
		"cdk-extensions.k8s_aws.ModifyOperation",
		"rename",
		[]interface{}{originalKey, renamedKey},
		&returns,
	)

	return returns
}

// Sets a field in the output to a specific value.
//
// If a field with the same name already exists it will be overridden with
// the specified value.
//
// Returns: A ModifyOperation object representing the set operation.
func ModifyOperation_Set(key *string, value *string) ModifyOperation {
	_init_.Initialize()

	if err := validateModifyOperation_SetParameters(key, value); err != nil {
		panic(err)
	}
	var returns ModifyOperation

	_jsii_.StaticInvoke(
		"cdk-extensions.k8s_aws.ModifyOperation",
		"set",
		[]interface{}{key, value},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ModifyOperation) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

