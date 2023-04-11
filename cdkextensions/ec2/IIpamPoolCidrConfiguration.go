package ec2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/constructs-go/constructs/v10"
)

type IIpamPoolCidrConfiguration interface {
	Bind(scope constructs.IConstruct) *ResolvedIpamPoolCidrConfiguration
	Inline() *bool
}

// The jsii proxy for IIpamPoolCidrConfiguration
type jsiiProxy_IIpamPoolCidrConfiguration struct {
	_ byte // padding
}

func (i *jsiiProxy_IIpamPoolCidrConfiguration) Bind(scope constructs.IConstruct) *ResolvedIpamPoolCidrConfiguration {
	if err := i.validateBindParameters(scope); err != nil {
		panic(err)
	}
	var returns *ResolvedIpamPoolCidrConfiguration

	_jsii_.Invoke(
		i,
		"bind",
		[]interface{}{scope},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_IIpamPoolCidrConfiguration) Inline() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"inline",
		&returns,
	)
	return returns
}

