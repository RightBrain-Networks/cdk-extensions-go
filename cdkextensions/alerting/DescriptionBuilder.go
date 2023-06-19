package alerting

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/vibe-io/cdk-extensions-go/cdkextensions/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsstepfunctions"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/vibe-io/cdk-extensions-go/cdkextensions/alerting/internal"
)

type DescriptionBuilder interface {
	constructs.Construct
	IDelayedChainable
	IDescriptionBuilderComponent
	Classifier() *string
	InitialDescription() *string
	InitializeNode() awsstepfunctions.IChainable
	// The tree node.
	Node() constructs.Node
	AddIterator(id *string, props *DescriptionBuilderIteratorProps) DescriptionBuilderIterator
	AddReference(id *string, props *AddReferenceProps) awsstepfunctions.IChainable
	AddSection(id *string, props *DescriptionBuilderSectionProps) DescriptionBuilderSection
	BuildId(prefix *string, id *string) *string
	Initialize() awsstepfunctions.IChainable
	RegisterBuilder(builder IDescriptionBuilderComponent) IDescriptionBuilderComponent
	RegisterChainable(chainable awsstepfunctions.IChainable) awsstepfunctions.IChainable
	Render() awsstepfunctions.IChainable
	SetDelimiter(id *string, props *SetDelimiterProps) awsstepfunctions.IChainable
	// Returns a string representation of this construct.
	ToString() *string
	Write(id *string, props *WriteProps) awsstepfunctions.IChainable
}

// The jsii proxy struct for DescriptionBuilder
type jsiiProxy_DescriptionBuilder struct {
	internal.Type__constructsConstruct
	jsiiProxy_IDelayedChainable
	jsiiProxy_IDescriptionBuilderComponent
}

func (j *jsiiProxy_DescriptionBuilder) Classifier() *string {
	var returns *string
	_jsii_.Get(
		j,
		"classifier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DescriptionBuilder) InitialDescription() *string {
	var returns *string
	_jsii_.Get(
		j,
		"initialDescription",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DescriptionBuilder) InitializeNode() awsstepfunctions.IChainable {
	var returns awsstepfunctions.IChainable
	_jsii_.Get(
		j,
		"initializeNode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DescriptionBuilder) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


func NewDescriptionBuilder(scope constructs.IConstruct, props *DescriptionBuilderProps) DescriptionBuilder {
	_init_.Initialize()

	if err := validateNewDescriptionBuilderParameters(scope, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_DescriptionBuilder{}

	_jsii_.Create(
		"cdk-extensions.alerting.DescriptionBuilder",
		[]interface{}{scope, props},
		&j,
	)

	return &j
}

func NewDescriptionBuilder_Override(d DescriptionBuilder, scope constructs.IConstruct, props *DescriptionBuilderProps) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk-extensions.alerting.DescriptionBuilder",
		[]interface{}{scope, props},
		d,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead.
func DescriptionBuilder_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDescriptionBuilder_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"cdk-extensions.alerting.DescriptionBuilder",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DescriptionBuilder) AddIterator(id *string, props *DescriptionBuilderIteratorProps) DescriptionBuilderIterator {
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

func (d *jsiiProxy_DescriptionBuilder) AddReference(id *string, props *AddReferenceProps) awsstepfunctions.IChainable {
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

func (d *jsiiProxy_DescriptionBuilder) AddSection(id *string, props *DescriptionBuilderSectionProps) DescriptionBuilderSection {
	if err := d.validateAddSectionParameters(id, props); err != nil {
		panic(err)
	}
	var returns DescriptionBuilderSection

	_jsii_.Invoke(
		d,
		"addSection",
		[]interface{}{id, props},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DescriptionBuilder) BuildId(prefix *string, id *string) *string {
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

func (d *jsiiProxy_DescriptionBuilder) Initialize() awsstepfunctions.IChainable {
	var returns awsstepfunctions.IChainable

	_jsii_.Invoke(
		d,
		"initialize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DescriptionBuilder) RegisterBuilder(builder IDescriptionBuilderComponent) IDescriptionBuilderComponent {
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

func (d *jsiiProxy_DescriptionBuilder) RegisterChainable(chainable awsstepfunctions.IChainable) awsstepfunctions.IChainable {
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

func (d *jsiiProxy_DescriptionBuilder) Render() awsstepfunctions.IChainable {
	var returns awsstepfunctions.IChainable

	_jsii_.Invoke(
		d,
		"render",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DescriptionBuilder) SetDelimiter(id *string, props *SetDelimiterProps) awsstepfunctions.IChainable {
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

func (d *jsiiProxy_DescriptionBuilder) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DescriptionBuilder) Write(id *string, props *WriteProps) awsstepfunctions.IChainable {
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

