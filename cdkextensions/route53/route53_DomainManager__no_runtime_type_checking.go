//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package route53

// Building without runtime type checking enabled, so all the below just return nil

func validateDomainManager_IsDnsResolvableParameters(node constructs.IConstruct) error {
	return nil
}

