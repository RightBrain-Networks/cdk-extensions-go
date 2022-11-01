package glue

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/vibe-io/cdk-extensions-go/cdkextensions/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsglue"
	"github.com/aws/constructs-go/constructs/v10"
)

type BasicColumn interface {
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
	Bind(scope constructs.IConstruct) *awsglue.CfnTable_ColumnProperty
}

// The jsii proxy struct for BasicColumn
type jsiiProxy_BasicColumn struct {
	jsiiProxy_Column
}

func (j *jsiiProxy_BasicColumn) Comment() *string {
	var returns *string
	_jsii_.Get(
		j,
		"comment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BasicColumn) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BasicColumn) TypeString() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeString",
		&returns,
	)
	return returns
}


func NewBasicColumn(props *BasicColumnProps) BasicColumn {
	_init_.Initialize()

	if err := validateNewBasicColumnParameters(props); err != nil {
		panic(err)
	}
	j := jsiiProxy_BasicColumn{}

	_jsii_.Create(
		"cdk-extensions.glue.BasicColumn",
		[]interface{}{props},
		&j,
	)

	return &j
}

func NewBasicColumn_Override(b BasicColumn, props *BasicColumnProps) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk-extensions.glue.BasicColumn",
		[]interface{}{props},
		b,
	)
}

func (b *jsiiProxy_BasicColumn) Bind(scope constructs.IConstruct) *awsglue.CfnTable_ColumnProperty {
	if err := b.validateBindParameters(scope); err != nil {
		panic(err)
	}
	var returns *awsglue.CfnTable_ColumnProperty

	_jsii_.Invoke(
		b,
		"bind",
		[]interface{}{scope},
		&returns,
	)

	return returns
}

