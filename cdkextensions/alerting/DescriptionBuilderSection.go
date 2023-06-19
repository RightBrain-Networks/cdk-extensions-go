package alerting

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/vibe-io/cdk-extensions-go/cdkextensions/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsstepfunctions"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/vibe-io/cdk-extensions-go/cdkextensions/alerting/internal"
)

type DescriptionBuilderSection interface {
	constructs.Construct
	IDelayedChainable
	IDescriptionBuilderComponent
	Classifier() *string
	// The tree node.
	Node() constructs.Node
	Refs() *[]*string
	Title() *string
	AddIterator(id *string, props *DescriptionBuilderIteratorProps) DescriptionBuilderIterator
	AddReference(id *string, props *AddReferenceProps) awsstepfunctions.IChainable
	AddReferenceCheck(ref *string)
	BuildId(prefix *string, id *string) *string
	RegisterBuilder(builder IDescriptionBuilderComponent) IDescriptionBuilderComponent
	RegisterChainable(chainable awsstepfunctions.IChainable) awsstepfunctions.IChainable
	Render() awsstepfunctions.IChainable
	SetDelimiter(id *string, props *SetDelimiterProps) awsstepfunctions.IChainable
	// Returns a string representation of this construct.
	ToString() *string
	Write(id *string, props *WriteProps) awsstepfunctions.IChainable
}

// The jsii proxy struct for DescriptionBuilderSection
type jsiiProxy_DescriptionBuilderSection struct {
	internal.Type__constructsConstruct
	jsiiProxy_IDelayedChainable
	jsiiProxy_IDescriptionBuilderComponent
}

func (j *jsiiProxy_DescriptionBuilderSection) Classifier() *string {
	var returns *string
	_jsii_.Get(
		j,
		"classifier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DescriptionBuilderSection) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DescriptionBuilderSection) Refs() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"refs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DescriptionBuilderSection) Title() *string {
	var returns *string
	_jsii_.Get(
		j,
		"title",
		&returns,
	)
	return returns
}


func NewDescriptionBuilderSection(scope constructs.IConstruct, id *string, props *DescriptionBuilderSectionProps) DescriptionBuilderSection {
	_init_.Initialize()

	if err := validateNewDescriptionBuilderSectionParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_DescriptionBuilderSection{}

	_jsii_.Create(
		"cdk-extensions.alerting.DescriptionBuilderSection",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewDescriptionBuilderSection_Override(d DescriptionBuilderSection, scope constructs.IConstruct, id *string, props *DescriptionBuilderSectionProps) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk-extensions.alerting.DescriptionBuilderSection",
		[]interface{}{scope, id, props},
		d,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead.
func DescriptionBuilderSection_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDescriptionBuilderSection_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"cdk-extensions.alerting.DescriptionBuilderSection",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DescriptionBuilderSection) AddIterator(id *string, props *DescriptionBuilderIteratorProps) DescriptionBuilderIterator {
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

func (d *jsiiProxy_DescriptionBuilderSection) AddReference(id *string, props *AddReferenceProps) awsstepfunctions.IChainable {
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

func (d *jsiiProxy_DescriptionBuilderSection) AddReferenceCheck(ref *string) {
	if err := d.validateAddReferenceCheckParameters(ref); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addReferenceCheck",
		[]interface{}{ref},
	)
}

func (d *jsiiProxy_DescriptionBuilderSection) BuildId(prefix *string, id *string) *string {
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

func (d *jsiiProxy_DescriptionBuilderSection) RegisterBuilder(builder IDescriptionBuilderComponent) IDescriptionBuilderComponent {
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

func (d *jsiiProxy_DescriptionBuilderSection) RegisterChainable(chainable awsstepfunctions.IChainable) awsstepfunctions.IChainable {
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

func (d *jsiiProxy_DescriptionBuilderSection) Render() awsstepfunctions.IChainable {
	var returns awsstepfunctions.IChainable

	_jsii_.Invoke(
		d,
		"render",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DescriptionBuilderSection) SetDelimiter(id *string, props *SetDelimiterProps) awsstepfunctions.IChainable {
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

func (d *jsiiProxy_DescriptionBuilderSection) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DescriptionBuilderSection) Write(id *string, props *WriteProps) awsstepfunctions.IChainable {
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

