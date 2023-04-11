package glue

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/vibe-io/cdk-extensions-go/cdkextensions/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsglue"
	"github.com/aws/constructs-go/constructs/v10"
)

type StructColumn interface {
	Column
	// A free-form text comment.
	// See: [AWS::Glue::Table Column](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-table-column.html#cfn-glue-table-column-comment)
	//
	Comment() *string
	// The name of the Column.
	// See: [AWS::Glue::Table Column](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-table-column.html#cfn-glue-table-column-name)
	//
	Name() *string
	TypeString() *string
	AddColumn(column Column)
	Bind(scope constructs.IConstruct) *awsglue.CfnTable_ColumnProperty
}

// The jsii proxy struct for StructColumn
type jsiiProxy_StructColumn struct {
	jsiiProxy_Column
}

func (j *jsiiProxy_StructColumn) Comment() *string {
	var returns *string
	_jsii_.Get(
		j,
		"comment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StructColumn) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StructColumn) TypeString() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeString",
		&returns,
	)
	return returns
}


func NewStructColumn(props *StructColumnProps) StructColumn {
	_init_.Initialize()

	if err := validateNewStructColumnParameters(props); err != nil {
		panic(err)
	}
	j := jsiiProxy_StructColumn{}

	_jsii_.Create(
		"cdk-extensions.glue.StructColumn",
		[]interface{}{props},
		&j,
	)

	return &j
}

func NewStructColumn_Override(s StructColumn, props *StructColumnProps) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk-extensions.glue.StructColumn",
		[]interface{}{props},
		s,
	)
}

func (s *jsiiProxy_StructColumn) AddColumn(column Column) {
	if err := s.validateAddColumnParameters(column); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"addColumn",
		[]interface{}{column},
	)
}

func (s *jsiiProxy_StructColumn) Bind(scope constructs.IConstruct) *awsglue.CfnTable_ColumnProperty {
	if err := s.validateBindParameters(scope); err != nil {
		panic(err)
	}
	var returns *awsglue.CfnTable_ColumnProperty

	_jsii_.Invoke(
		s,
		"bind",
		[]interface{}{scope},
		&returns,
	)

	return returns
}

