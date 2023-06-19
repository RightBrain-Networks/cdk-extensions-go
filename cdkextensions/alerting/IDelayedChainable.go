package alerting

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsstepfunctions"
)

type IDelayedChainable interface {
	Render() awsstepfunctions.IChainable
}

// The jsii proxy for IDelayedChainable
type jsiiProxy_IDelayedChainable struct {
	_ byte // padding
}

func (i *jsiiProxy_IDelayedChainable) Render() awsstepfunctions.IChainable {
	var returns awsstepfunctions.IChainable

	_jsii_.Invoke(
		i,
		"render",
		nil, // no parameters
		&returns,
	)

	return returns
}

