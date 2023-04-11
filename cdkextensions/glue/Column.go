package glue

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/vibe-io/cdk-extensions-go/cdkextensions/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsglue"
	"github.com/aws/constructs-go/constructs/v10"
)

type Column interface {
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

// The jsii proxy struct for Column
type jsiiProxy_Column struct {
	_ byte // padding
}

func (j *jsiiProxy_Column) Comment() *string {
	var returns *string
	_jsii_.Get(
		j,
		"comment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Column) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Column) TypeString() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeString",
		&returns,
	)
	return returns
}


func NewColumn_Override(c Column, props *ColumnProps) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk-extensions.glue.Column",
		[]interface{}{props},
		c,
	)
}

func (c *jsiiProxy_Column) Bind(scope constructs.IConstruct) *awsglue.CfnTable_ColumnProperty {
	if err := c.validateBindParameters(scope); err != nil {
		panic(err)
	}
	var returns *awsglue.CfnTable_ColumnProperty

	_jsii_.Invoke(
		c,
		"bind",
		[]interface{}{scope},
		&returns,
	)

	return returns
}

