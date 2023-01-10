package glue

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/vibe-io/cdk-extensions-go/cdkextensions/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsglue"
	"github.com/aws/constructs-go/constructs/v10"
)

type CrawlerPredicate interface {
	PredicateBase
	ITriggerPredicate
	Crawler() ICrawler
	LogicalOperator() PredicateLogicalOperator
	State() CrawlerState
	Bind(scope constructs.IConstruct) *awsglue.CfnTrigger_ConditionProperty
	BindOptions(_scope constructs.IConstruct) interface{}
}

// The jsii proxy struct for CrawlerPredicate
type jsiiProxy_CrawlerPredicate struct {
	jsiiProxy_PredicateBase
	jsiiProxy_ITriggerPredicate
}

func (j *jsiiProxy_CrawlerPredicate) Crawler() ICrawler {
	var returns ICrawler
	_jsii_.Get(
		j,
		"crawler",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CrawlerPredicate) LogicalOperator() PredicateLogicalOperator {
	var returns PredicateLogicalOperator
	_jsii_.Get(
		j,
		"logicalOperator",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CrawlerPredicate) State() CrawlerState {
	var returns CrawlerState
	_jsii_.Get(
		j,
		"state",
		&returns,
	)
	return returns
}


func NewCrawlerPredicate(crawler ICrawler, options *CrawlerPredicateOptions) CrawlerPredicate {
	_init_.Initialize()

	if err := validateNewCrawlerPredicateParameters(crawler, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CrawlerPredicate{}

	_jsii_.Create(
		"cdk-extensions.glue.CrawlerPredicate",
		[]interface{}{crawler, options},
		&j,
	)

	return &j
}

func NewCrawlerPredicate_Override(c CrawlerPredicate, crawler ICrawler, options *CrawlerPredicateOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk-extensions.glue.CrawlerPredicate",
		[]interface{}{crawler, options},
		c,
	)
}

func (c *jsiiProxy_CrawlerPredicate) Bind(scope constructs.IConstruct) *awsglue.CfnTrigger_ConditionProperty {
	if err := c.validateBindParameters(scope); err != nil {
		panic(err)
	}
	var returns *awsglue.CfnTrigger_ConditionProperty

	_jsii_.Invoke(
		c,
		"bind",
		[]interface{}{scope},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CrawlerPredicate) BindOptions(_scope constructs.IConstruct) interface{} {
	if err := c.validateBindOptionsParameters(_scope); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		c,
		"bindOptions",
		[]interface{}{_scope},
		&returns,
	)

	return returns
}

