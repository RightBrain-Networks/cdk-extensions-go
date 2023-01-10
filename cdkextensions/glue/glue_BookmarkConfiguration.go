package glue

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/vibe-io/cdk-extensions-go/cdkextensions/jsii"
)

type BookmarkConfiguration interface {
	Range() *BookmarkRange
	Value() *string
}

// The jsii proxy struct for BookmarkConfiguration
type jsiiProxy_BookmarkConfiguration struct {
	_ byte // padding
}

func (j *jsiiProxy_BookmarkConfiguration) Range() *BookmarkRange {
	var returns *BookmarkRange
	_jsii_.Get(
		j,
		"range",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BookmarkConfiguration) Value() *string {
	var returns *string
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


func BookmarkConfiguration_Disable() BookmarkConfiguration {
	_init_.Initialize()

	var returns BookmarkConfiguration

	_jsii_.StaticInvoke(
		"cdk-extensions.glue.BookmarkConfiguration",
		"disable",
		nil, // no parameters
		&returns,
	)

	return returns
}

func BookmarkConfiguration_Enable() BookmarkConfiguration {
	_init_.Initialize()

	var returns BookmarkConfiguration

	_jsii_.StaticInvoke(
		"cdk-extensions.glue.BookmarkConfiguration",
		"enable",
		nil, // no parameters
		&returns,
	)

	return returns
}

func BookmarkConfiguration_Of(value *string, range_ *BookmarkRange) BookmarkConfiguration {
	_init_.Initialize()

	if err := validateBookmarkConfiguration_OfParameters(value, range_); err != nil {
		panic(err)
	}
	var returns BookmarkConfiguration

	_jsii_.StaticInvoke(
		"cdk-extensions.glue.BookmarkConfiguration",
		"of",
		[]interface{}{value, range_},
		&returns,
	)

	return returns
}

func BookmarkConfiguration_Pause(range_ *BookmarkRange) BookmarkConfiguration {
	_init_.Initialize()

	if err := validateBookmarkConfiguration_PauseParameters(range_); err != nil {
		panic(err)
	}
	var returns BookmarkConfiguration

	_jsii_.StaticInvoke(
		"cdk-extensions.glue.BookmarkConfiguration",
		"pause",
		[]interface{}{range_},
		&returns,
	)

	return returns
}

