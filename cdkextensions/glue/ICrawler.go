package glue

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/vibe-io/cdk-extensions-go/cdkextensions/glue/internal"
)

type ICrawler interface {
	constructs.IConstruct
	// The Amazon Resource Name (ARN) of the crawler.
	CrawlerArn() *string
	// The name of the crawler.
	CrawlerName() *string
}

// The jsii proxy for ICrawler
type jsiiProxy_ICrawler struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_ICrawler) CrawlerArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"crawlerArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ICrawler) CrawlerName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"crawlerName",
		&returns,
	)
	return returns
}

