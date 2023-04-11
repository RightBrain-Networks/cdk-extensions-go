package ec2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/constructs-go/constructs/v10"
)

type IIpamAllocationConfiguration interface {
	Bind(scope constructs.IConstruct) *ResolvedIpamAllocationConfiguration
}

// The jsii proxy for IIpamAllocationConfiguration
type jsiiProxy_IIpamAllocationConfiguration struct {
	_ byte // padding
}

func (i *jsiiProxy_IIpamAllocationConfiguration) Bind(scope constructs.IConstruct) *ResolvedIpamAllocationConfiguration {
	if err := i.validateBindParameters(scope); err != nil {
		panic(err)
	}
	var returns *ResolvedIpamAllocationConfiguration

	_jsii_.Invoke(
		i,
		"bind",
		[]interface{}{scope},
		&returns,
	)

	return returns
}

