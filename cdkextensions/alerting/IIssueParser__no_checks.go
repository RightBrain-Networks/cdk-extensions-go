//go:build no_runtime_type_checking

package alerting

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_IIssueParser) validateBindParameters(node constructs.IConstruct) error {
	return nil
}

