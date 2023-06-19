package alerting

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsstepfunctions"
)

type IDescriptionBuilderComponent interface {
	IDelayedChainable
	AddIterator(id *string, props *DescriptionBuilderIteratorProps) DescriptionBuilderIterator
	AddReference(id *string, props *AddReferenceProps) awsstepfunctions.IChainable
	SetDelimiter(id *string, props *SetDelimiterProps) awsstepfunctions.IChainable
	Write(id *string, props *WriteProps) awsstepfunctions.IChainable
	Classifier() *string
}

// The jsii proxy for IDescriptionBuilderComponent
type jsiiProxy_IDescriptionBuilderComponent struct {
	jsiiProxy_IDelayedChainable
}

func (i *jsiiProxy_IDescriptionBuilderComponent) AddIterator(id *string, props *DescriptionBuilderIteratorProps) DescriptionBuilderIterator {
	if err := i.validateAddIteratorParameters(id, props); err != nil {
		panic(err)
	}
	var returns DescriptionBuilderIterator

	_jsii_.Invoke(
		i,
		"addIterator",
		[]interface{}{id, props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IDescriptionBuilderComponent) AddReference(id *string, props *AddReferenceProps) awsstepfunctions.IChainable {
	if err := i.validateAddReferenceParameters(id, props); err != nil {
		panic(err)
	}
	var returns awsstepfunctions.IChainable

	_jsii_.Invoke(
		i,
		"addReference",
		[]interface{}{id, props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IDescriptionBuilderComponent) SetDelimiter(id *string, props *SetDelimiterProps) awsstepfunctions.IChainable {
	if err := i.validateSetDelimiterParameters(id, props); err != nil {
		panic(err)
	}
	var returns awsstepfunctions.IChainable

	_jsii_.Invoke(
		i,
		"setDelimiter",
		[]interface{}{id, props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IDescriptionBuilderComponent) Write(id *string, props *WriteProps) awsstepfunctions.IChainable {
	if err := i.validateWriteParameters(id, props); err != nil {
		panic(err)
	}
	var returns awsstepfunctions.IChainable

	_jsii_.Invoke(
		i,
		"write",
		[]interface{}{id, props},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_IDescriptionBuilderComponent) Classifier() *string {
	var returns *string
	_jsii_.Get(
		j,
		"classifier",
		&returns,
	)
	return returns
}

