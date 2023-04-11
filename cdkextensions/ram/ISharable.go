package ram

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/constructs-go/constructs/v10"
)

// Represents an AWS resource that can be shared via AWS Resource Access Manager (RAM).
type ISharable interface {
	// Configures resource sharing for the associated resource.
	Share(scope constructs.IConstruct) *string
}

// The jsii proxy for ISharable
type jsiiProxy_ISharable struct {
	_ byte // padding
}

func (i *jsiiProxy_ISharable) Share(scope constructs.IConstruct) *string {
	if err := i.validateShareParameters(scope); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		i,
		"share",
		[]interface{}{scope},
		&returns,
	)

	return returns
}

