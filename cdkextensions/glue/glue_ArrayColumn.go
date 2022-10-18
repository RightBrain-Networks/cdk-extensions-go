package glue

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/vibe-io/cdk-extensions-go/cdkextensions/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsglue"
	"github.com/aws/constructs-go/constructs/v10"
)

type ArrayColumn interface {
	Column
	Comment() *string
	Name() *string
	TypeString() *string
	Bind(scope constructs.IConstruct) *awsglue.CfnTable_ColumnProperty
}

// The jsii proxy struct for ArrayColumn
type jsiiProxy_ArrayColumn struct {
	jsiiProxy_Column
}

func (j *jsiiProxy_ArrayColumn) Comment() *string {
	var returns *string
	_jsii_.Get(
		j,
		"comment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArrayColumn) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArrayColumn) TypeString() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeString",
		&returns,
	)
	return returns
}


func NewArrayColumn(props *ArrayColumnProps) ArrayColumn {
	_init_.Initialize()

	if err := validateNewArrayColumnParameters(props); err != nil {
		panic(err)
	}
	j := jsiiProxy_ArrayColumn{}

	_jsii_.Create(
		"cdk-extensions.glue.ArrayColumn",
		[]interface{}{props},
		&j,
	)

	return &j
}

func NewArrayColumn_Override(a ArrayColumn, props *ArrayColumnProps) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk-extensions.glue.ArrayColumn",
		[]interface{}{props},
		a,
	)
}

func (a *jsiiProxy_ArrayColumn) Bind(scope constructs.IConstruct) *awsglue.CfnTable_ColumnProperty {
	if err := a.validateBindParameters(scope); err != nil {
		panic(err)
	}
	var returns *awsglue.CfnTable_ColumnProperty

	_jsii_.Invoke(
		a,
		"bind",
		[]interface{}{scope},
		&returns,
	)

	return returns
}

