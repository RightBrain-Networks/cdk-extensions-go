//go:build no_runtime_type_checking

package resourcegroups

// Building without runtime type checking enabled, so all the below just return nil

func (t *jsiiProxy_TagFilter) validateAddResourceTypeParameters(typeId *string) error {
	return nil
}

func (t *jsiiProxy_TagFilter) validateAddTagFilterParameters(key *string) error {
	return nil
}

func (t *jsiiProxy_TagFilter) validateBindParameters(_scope constructs.IConstruct) error {
	return nil
}

