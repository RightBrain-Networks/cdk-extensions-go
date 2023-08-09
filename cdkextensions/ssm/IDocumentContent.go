package ssm

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/constructs-go/constructs/v10"
)

type IDocumentContent interface {
	Bind(scope constructs.IConstruct) *DocumentContentResult
}

// The jsii proxy for IDocumentContent
type jsiiProxy_IDocumentContent struct {
	_ byte // padding
}

func (i *jsiiProxy_IDocumentContent) Bind(scope constructs.IConstruct) *DocumentContentResult {
	if err := i.validateBindParameters(scope); err != nil {
		panic(err)
	}
	var returns *DocumentContentResult

	_jsii_.Invoke(
		i,
		"bind",
		[]interface{}{scope},
		&returns,
	)

	return returns
}

