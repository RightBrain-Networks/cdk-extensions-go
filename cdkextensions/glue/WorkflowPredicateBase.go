package glue

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/vibe-io/cdk-extensions-go/cdkextensions/jsii"

	"github.com/aws/constructs-go/constructs/v10"
)

// Base class providing common functionality for trigger predicate conditions.
type WorkflowPredicateBase interface {
	// Associates the predicate with a construct that is configuring a trigger for a Glue workflow.
	//
	// Returns: A configuration object that can be used to configure a predicate
	// condition for the Glue trigger.
	BindOptions(_scope constructs.IConstruct) interface{}
}

// The jsii proxy struct for WorkflowPredicateBase
type jsiiProxy_WorkflowPredicateBase struct {
	_ byte // padding
}

// Create a new instance of the WorkflowPredicateBase class.
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

// Create a new instance of the WorkflowPredicateBase class.
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

