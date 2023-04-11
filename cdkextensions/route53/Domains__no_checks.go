//go:build no_runtime_type_checking

package route53

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_Domains) validateAddParameters(hostedZone awsroute53.IHostedZone, isPublic *bool, options *DomainOptions) error {
	return nil
}

func validateDomains_OfParameters(scope constructs.IConstruct) error {
	return nil
}

