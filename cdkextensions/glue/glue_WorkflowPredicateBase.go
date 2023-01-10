package glue

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/vibe-io/cdk-extensions-go/cdkextensions/jsii"

	"github.com/aws/constructs-go/constructs/v10"
)

type WorkflowPredicateBase interface {
	BindOptions(_scope constructs.IConstruct) interface{}
}

// The jsii proxy struct for WorkflowPredicateBase
type jsiiProxy_WorkflowPredicateBase struct {
	_ byte // padding
}

func NewWorkflowPredicateBase(_options *WorkflowPredicateOptions) WorkflowPredicateBase {
	_init_.Initialize()

	if err := validateNewWorkflowPredicateBaseParameters(_options); err != nil {
		panic(err)
	}
	j := jsiiProxy_WorkflowPredicateBase{}

	_jsii_.Create(
		"cdk-extensions.glue.WorkflowPredicateBase",
		[]interface{}{_options},
		&j,
	)

	return &j
}

func NewWorkflowPredicateBase_Override(w WorkflowPredicateBase, _options *WorkflowPredicateOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk-extensions.glue.WorkflowPredicateBase",
		[]interface{}{_options},
		w,
	)
}

func (w *jsiiProxy_WorkflowPredicateBase) BindOptions(_scope constructs.IConstruct) interface{} {
	if err := w.validateBindOptionsParameters(_scope); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		w,
		"bindOptions",
		[]interface{}{_scope},
		&returns,
	)

	return returns
}

