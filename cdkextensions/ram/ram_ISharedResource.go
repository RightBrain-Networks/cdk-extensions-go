package ram

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/constructs-go/constructs/v10"
)

type ISharedResource interface {
	Bind(scope constructs.IConstruct) *string
}

// The jsii proxy for ISharedResource
type jsiiProxy_ISharedResource struct {
	_ byte // padding
}

func (i *jsiiProxy_ISharedResource) Bind(scope constructs.IConstruct) *string {
	if err := i.validateBindParameters(scope); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		i,
		"bind",
		[]interface{}{scope},
		&returns,
	)

	return returns
}

