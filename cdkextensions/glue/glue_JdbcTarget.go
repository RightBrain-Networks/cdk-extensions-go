package glue

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/vibe-io/cdk-extensions-go/cdkextensions/jsii"
)

type JdbcTarget interface {
	ICrawlerTarget
	Connection() Connection
	AddExclusion(exclusion *string)
	AddPath(path *string)
	Bind(_crawler Crawler) *CrawlerTargetCollection
}

// The jsii proxy struct for JdbcTarget
type jsiiProxy_JdbcTarget struct {
	jsiiProxy_ICrawlerTarget
}

func (j *jsiiProxy_JdbcTarget) Connection() Connection {
	var returns Connection
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}


func NewJdbcTarget(connection Connection, options *JdbcTargetOptions) JdbcTarget {
	_init_.Initialize()

	if err := validateNewJdbcTargetParameters(connection, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_JdbcTarget{}

	_jsii_.Create(
		"cdk-extensions.glue.JdbcTarget",
		[]interface{}{connection, options},
		&j,
	)

	return &j
}

func NewJdbcTarget_Override(j JdbcTarget, connection Connection, options *JdbcTargetOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk-extensions.glue.JdbcTarget",
		[]interface{}{connection, options},
		j,
	)
}

func (j *jsiiProxy_JdbcTarget) AddExclusion(exclusion *string) {
	if err := j.validateAddExclusionParameters(exclusion); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"addExclusion",
		[]interface{}{exclusion},
	)
}

func (j *jsiiProxy_JdbcTarget) AddPath(path *string) {
	if err := j.validateAddPathParameters(path); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"addPath",
		[]interface{}{path},
	)
}

func (j *jsiiProxy_JdbcTarget) Bind(_crawler Crawler) *CrawlerTargetCollection {
	if err := j.validateBindParameters(_crawler); err != nil {
		panic(err)
	}
	var returns *CrawlerTargetCollection

	_jsii_.Invoke(
		j,
		"bind",
		[]interface{}{_crawler},
		&returns,
	)

	return returns
}

