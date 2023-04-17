package athena

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/constructs-go/constructs/v10"
)

type IAthenaResultEncryption interface {
	Bind(scope constructs.IConstruct) *AthenaResultEncryptionConfiguration
}

// The jsii proxy for IAthenaResultEncryption
type jsiiProxy_IAthenaResultEncryption struct {
	_ byte // padding
}

func (i *jsiiProxy_IAthenaResultEncryption) Bind(scope constructs.IConstruct) *AthenaResultEncryptionConfiguration {
	if err := i.validateBindParameters(scope); err != nil {
		panic(err)
	}
	var returns *AthenaResultEncryptionConfiguration

	_jsii_.Invoke(
		i,
		"bind",
		[]interface{}{scope},
		&returns,
	)

	return returns
}

