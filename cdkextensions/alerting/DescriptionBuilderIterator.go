package alerting

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/vibe-io/cdk-extensions-go/cdkextensions/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsstepfunctions"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/vibe-io/cdk-extensions-go/cdkextensions/alerting/internal"
)

type DescriptionBuilderIterator interface {
	constructs.Construct
	IDelayedChainable
	IDescriptionBuilderComponent
	ArrayRef() *string
	Classifier() *string
	FieldDelimiter() AppendDelimiter
	// The tree node.
	Node() constructs.Node
	RecordDelimiter() AppendDelimiter
	ResultPath() *string
	SectionDelimiter() AppendDelimiter
	Title() *string
	AddIterator(id *string, props *DescriptionBuilderIteratorProps) DescriptionBuilderIterator
	AddReference(id *string, props *AddReferenceProps) awsstepfunctions.IChainable
	BuildId(prefix *string, id *string) *string
	RegisterBuilder(builder IDescriptionBuilderComponent) IDescriptionBuilderComponent
	RegisterChainable(chainable awsstepfunctions.IChainable) awsstepfunctions.IChainable
	Render() awsstepfunctions.IChainable
	SetDelimiter(id *string, props *SetDelimiterProps) awsstepfunctions.IChainable
	// Returns a string representation of this construct.
	ToString() *string
	Write(id *string, props *WriteProps) awsstepfunctions.IChainable
}

// The jsii proxy struct for DescriptionBuilderIterator
type jsiiProxy_DescriptionBuilderIterator struct {
	internal.Type__constructsConstruct
	jsiiProxy_IDelayedChainable
	jsiiProxy_IDescriptionBuilderComponent
}

func (j *jsiiProxy_DescriptionBuilderIterator) ArrayRef() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arrayRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DescriptionBuilderIterator) Classifier() *string {
	var returns *string
	_jsii_.Get(
		j,
		"classifier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DescriptionBuilderIterator) FieldDelimiter() AppendDelimiter {
	var returns AppendDelimiter
	_jsii_.Get(
		j,
		"fieldDelimiter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DescriptionBuilderIterator) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DescriptionBuilderIterator) RecordDelimiter() AppendDelimiter {
	var returns AppendDelimiter
	_jsii_.Get(
		j,
		"recordDelimiter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DescriptionBuilderIterator) ResultPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resultPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DescriptionBuilderIterator) SectionDelimiter() AppendDelimiter {
	var returns AppendDelimiter
	_jsii_.Get(
		j,
		"sectionDelimiter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DescriptionBuilderIterator) Title() *string {
	var returns *string
	_jsii_.Get(
		j,
		"title",
		&returns,
	)
	return returns
}


func NewDescriptionBuilderIterator(scope constructs.IConstruct, id *string, props *DescriptionBuilderIteratorProps) DescriptionBuilderIterator {
	_init_.Initialize()

	if err := validateNewDescriptionBuilderIteratorParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_DescriptionBuilderIterator{}

	_jsii_.Create(
		"cdk-extensions.alerting.DescriptionBuilderIterator",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewDescriptionBuilderIterator_Override(d DescriptionBuilderIterator, scope constructs.IConstruct, id *string, props *DescriptionBuilderIteratorProps) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk-extensions.alerting.DescriptionBuilderIterator",
		[]interface{}{scope, id, props},
		d,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead.
func DescriptionBuilderIterator_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDescriptionBuilderIterator_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"cdk-extensions.alerting.DescriptionBuilderIterator",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DescriptionBuilderIterator) AddIterator(id *string, props *DescriptionBuilderIteratorProps) DescriptionBuilderIterator {
	if err := d.validateAddIteratorParameters(id, props); err != nil {
		panic(err)
	}
	var returns DescriptionBuilderIterator

	_jsii_.Invoke(
		d,
		"addIterator",
		[]interface{}{id, props},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DescriptionBuilderIterator) AddReference(id *string, props *AddReferenceProps) awsstepfunctions.IChainable {
	if err := d.validateAddReferenceParameters(id, props); err != nil {
		panic(err)
	}
	var returns awsstepfunctions.IChainable

	_jsii_.Invoke(
		d,
		"addReference",
		[]interface{}{id, props},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DescriptionBuilderIterator) BuildId(prefix *string, id *string) *string {
	if err := d.validateBuildIdParameters(prefix); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		d,
		"buildId",
		[]interface{}{prefix, id},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DescriptionBuilderIterator) RegisterBuilder(builder IDescriptionBuilderComponent) IDescriptionBuilderComponent {
	if err := d.validateRegisterBuilderParameters(builder); err != nil {
		panic(err)
	}
	var returns IDescriptionBuilderComponent

	_jsii_.Invoke(
		d,
		"registerBuilder",
		[]interface{}{builder},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DescriptionBuilderIterator) RegisterChainable(chainable awsstepfunctions.IChainable) awsstepfunctions.IChainable {
	if err := d.validateRegisterChainableParameters(chainable); err != nil {
		panic(err)
	}
	var returns awsstepfunctions.IChainable

	_jsii_.Invoke(
		d,
		"registerChainable",
		[]interface{}{chainable},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DescriptionBuilderIterator) Render() awsstepfunctions.IChainable {
	var returns awsstepfunctions.IChainable

	_jsii_.Invoke(
		d,
		"render",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DescriptionBuilderIterator) SetDelimiter(id *string, props *SetDelimiterProps) awsstepfunctions.IChainable {
	if err := d.validateSetDelimiterParameters(id, props); err != nil {
		panic(err)
	}
	var returns awsstepfunctions.IChainable

	_jsii_.Invoke(
		d,
		"setDelimiter",
		[]interface{}{id, props},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DescriptionBuilderIterator) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DescriptionBuilderIterator) Write(id *string, props *WriteProps) awsstepfunctions.IChainable {
	if err := d.validateWriteParameters(id, props); err != nil {
		panic(err)
	}
	var returns awsstepfunctions.IChainable

	_jsii_.Invoke(
		d,
		"write",
		[]interface{}{id, props},
		&returns,
	)

	return returns
}

