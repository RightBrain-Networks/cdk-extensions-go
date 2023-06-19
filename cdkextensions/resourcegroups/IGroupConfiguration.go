package resourcegroups

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/constructs-go/constructs/v10"
)

type IGroupConfiguration interface {
	Bind(scope constructs.IConstruct) *BoundGroupConfiguration
}

// The jsii proxy for IGroupConfiguration
type jsiiProxy_IGroupConfiguration struct {
	_ byte // padding
}

func (i *jsiiProxy_IGroupConfiguration) Bind(scope constructs.IConstruct) *BoundGroupConfiguration {
	if err := i.validateBindParameters(scope); err != nil {
		panic(err)
	}
	var returns *BoundGroupConfiguration

	_jsii_.Invoke(
		i,
		"bind",
		[]interface{}{scope},
		&returns,
	)

	return returns
}

