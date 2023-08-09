package ssm

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/vibe-io/cdk-extensions-go/cdkextensions/ssm/internal"
)

type IDocument interface {
	awscdk.IResource
	DocumentArn() *string
	DocumentName() *string
}

// The jsii proxy for IDocument
type jsiiProxy_IDocument struct {
	internal.Type__awscdkIResource
}

func (j *jsiiProxy_IDocument) DocumentArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"documentArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IDocument) DocumentName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"documentName",
		&returns,
	)
	return returns
}

