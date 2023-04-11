package glue

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

type ICrawlerTarget interface {
	Bind(crawler Crawler) *CrawlerTargetCollection
}

// The jsii proxy for ICrawlerTarget
type jsiiProxy_ICrawlerTarget struct {
	_ byte // padding
}

func (i *jsiiProxy_ICrawlerTarget) Bind(crawler Crawler) *CrawlerTargetCollection {
	if err := i.validateBindParameters(crawler); err != nil {
		panic(err)
	}
	var returns *CrawlerTargetCollection

	_jsii_.Invoke(
		i,
		"bind",
		[]interface{}{crawler},
		&returns,
	)

	return returns
}

