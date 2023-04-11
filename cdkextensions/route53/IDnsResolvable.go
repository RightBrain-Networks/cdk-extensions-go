package route53

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/vibe-io/cdk-extensions-go/cdkextensions/route53/internal"
)

type IDnsResolvable interface {
	constructs.IConstruct
	RegisterDomain(domain Domain)
	DomainDiscovery() DomainDiscovery
}

// The jsii proxy for IDnsResolvable
type jsiiProxy_IDnsResolvable struct {
	internal.Type__constructsIConstruct
}

func (i *jsiiProxy_IDnsResolvable) RegisterDomain(domain Domain) {
	if err := i.validateRegisterDomainParameters(domain); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"registerDomain",
		[]interface{}{domain},
	)
}

func (j *jsiiProxy_IDnsResolvable) DomainDiscovery() DomainDiscovery {
	var returns DomainDiscovery
	_jsii_.Get(
		j,
		"domainDiscovery",
		&returns,
	)
	return returns
}

