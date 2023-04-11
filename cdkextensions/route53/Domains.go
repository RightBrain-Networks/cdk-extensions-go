package route53

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/vibe-io/cdk-extensions-go/cdkextensions/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsroute53"
	"github.com/aws/constructs-go/constructs/v10"
)

type Domains interface {
	Add(hostedZone awsroute53.IHostedZone, isPublic *bool, options *DomainOptions)
}

// The jsii proxy struct for Domains
type jsiiProxy_Domains struct {
	_ byte // padding
}

func Domains_Of(scope constructs.IConstruct) Domains {
	_init_.Initialize()

	if err := validateDomains_OfParameters(scope); err != nil {
		panic(err)
	}
	var returns Domains

	_jsii_.StaticInvoke(
		"cdk-extensions.route53.Domains",
		"of",
		[]interface{}{scope},
		&returns,
	)

	return returns
}

func Domains_ROOT() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"cdk-extensions.route53.Domains",
		"ROOT",
		&returns,
	)
	return returns
}

func (d *jsiiProxy_Domains) Add(hostedZone awsroute53.IHostedZone, isPublic *bool, options *DomainOptions) {
	if err := d.validateAddParameters(hostedZone, isPublic, options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"add",
		[]interface{}{hostedZone, isPublic, options},
	)
}

