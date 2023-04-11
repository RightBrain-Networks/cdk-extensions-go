package ram

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/constructs-go/constructs/v10"
)

type ISharedPrincipal interface {
	Bind(scope constructs.IConstruct) *string
}

// The jsii proxy for ISharedPrincipal
type jsiiProxy_ISharedPrincipal struct {
	_ byte // padding
}

func (i *jsiiProxy_ISharedPrincipal) Bind(scope constructs.IConstruct) *string {
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

