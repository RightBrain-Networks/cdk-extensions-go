package securityhub

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/vibe-io/cdk-extensions-go/cdkextensions/securityhub/internal"
)

type IStandard interface {
	awscdk.IResource
	StandardArn() *string
}

// The jsii proxy for IStandard
type jsiiProxy_IStandard struct {
	internal.Type__awscdkIResource
}

func (j *jsiiProxy_IStandard) StandardArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"standardArn",
		&returns,
	)
	return returns
}

