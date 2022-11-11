package route53

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/vibe-io/cdk-extensions-go/cdkextensions/jsii"

	"github.com/aws/constructs-go/constructs/v10"
)

type DomainManager interface {
}

// The jsii proxy struct for DomainManager
type jsiiProxy_DomainManager struct {
	_ byte // padding
}

func NewDomainManager() DomainManager {
	_init_.Initialize()

	j := jsiiProxy_DomainManager{}

	_jsii_.Create(
		"cdk-extensions.route53.DomainManager",
		nil, // no parameters
		&j,
	)

	return &j
}

func NewDomainManager_Override(d DomainManager) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk-extensions.route53.DomainManager",
		nil, // no parameters
		d,
	)
}

func DomainManager_IsDnsResolvable(node constructs.IConstruct) *bool {
	_init_.Initialize()

	if err := validateDomainManager_IsDnsResolvableParameters(node); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"cdk-extensions.route53.DomainManager",
		"isDnsResolvable",
		[]interface{}{node},
		&returns,
	)

	return returns
}

