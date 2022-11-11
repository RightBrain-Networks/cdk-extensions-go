package route53

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awscertificatemanager"
)

type DomainOptions struct {
	Certificate awscertificatemanager.ICertificate `field:"optional" json:"certificate" yaml:"certificate"`
	Subdomain *string `field:"optional" json:"subdomain" yaml:"subdomain"`
}

