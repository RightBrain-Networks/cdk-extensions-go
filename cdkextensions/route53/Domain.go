package route53

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/vibe-io/cdk-extensions-go/cdkextensions/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2/awscertificatemanager"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsroute53"
	"github.com/aws/constructs-go/constructs/v10"
)

type Domain interface {
	Certificate() awscertificatemanager.ICertificate
	Fqdn() *string
	IsPublic() *bool
	Subdomain() *string
	Zone() awsroute53.IHostedZone
	ZoneName() *string
	Visit(node constructs.IConstruct)
}

// The jsii proxy struct for Domain
type jsiiProxy_Domain struct {
	_ byte // padding
}

func (j *jsiiProxy_Domain) Certificate() awscertificatemanager.ICertificate {
	var returns awscertificatemanager.ICertificate
	_jsii_.Get(
		j,
		"certificate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Domain) Fqdn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqdn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Domain) IsPublic() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isPublic",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Domain) Subdomain() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subdomain",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Domain) Zone() awsroute53.IHostedZone {
	var returns awsroute53.IHostedZone
	_jsii_.Get(
		j,
		"zone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Domain) ZoneName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"zoneName",
		&returns,
	)
	return returns
}


func NewDomain(zone awsroute53.IHostedZone, isPublic *bool, options *DomainOptions) Domain {
	_init_.Initialize()

	if err := validateNewDomainParameters(zone, isPublic, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_Domain{}

	_jsii_.Create(
		"cdk-extensions.route53.Domain",
		[]interface{}{zone, isPublic, options},
		&j,
	)

	return &j
}

func NewDomain_Override(d Domain, zone awsroute53.IHostedZone, isPublic *bool, options *DomainOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk-extensions.route53.Domain",
		[]interface{}{zone, isPublic, options},
		d,
	)
}

func (d *jsiiProxy_Domain) Visit(node constructs.IConstruct) {
	if err := d.validateVisitParameters(node); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"visit",
		[]interface{}{node},
	)
}

