package glue

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/vibe-io/cdk-extensions-go/cdkextensions/jsii"

	"github.com/aws/constructs-go/constructs/v10"
)

type PredicateBase interface {
	BindOptions(_scope constructs.IConstruct) interface{}
}

// The jsii proxy struct for PredicateBase
type jsiiProxy_PredicateBase struct {
	_ byte // padding
}

func NewPredicateBase(_options *PredicateOptions) PredicateBase {
	_init_.Initialize()

	if err := validateNewPredicateBaseParameters(_options); err != nil {
		panic(err)
	}
	j := jsiiProxy_PredicateBase{}

	_jsii_.Create(
		"cdk-extensions.glue.PredicateBase",
		[]interface{}{_options},
		&j,
	)

	return &j
}

func NewPredicateBase_Override(p PredicateBase, _options *PredicateOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk-extensions.glue.PredicateBase",
		[]interface{}{_options},
		p,
	)
}

func (p *jsiiProxy_PredicateBase) BindOptions(_scope constructs.IConstruct) interface{} {
	if err := p.validateBindOptionsParameters(_scope); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		p,
		"bindOptions",
		[]interface{}{_scope},
		&returns,
	)

	return returns
}

