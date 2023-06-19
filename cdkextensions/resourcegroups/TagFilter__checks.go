//go:build !no_runtime_type_checking

package resourcegroups

import (
	"fmt"

	"github.com/aws/constructs-go/constructs/v10"
)

func (t *jsiiProxy_TagFilter) validateAddResourceTypeParameters(typeId *string) error {
	if typeId == nil {
		return fmt.Errorf("parameter typeId is required, but nil was provided")
	}

	return nil
}

func (t *jsiiProxy_TagFilter) validateAddTagFilterParameters(key *string) error {
	if key == nil {
		return fmt.Errorf("parameter key is required, but nil was provided")
	}

	return nil
}

func (t *jsiiProxy_TagFilter) validateBindParameters(_scope constructs.IConstruct) error {
	if _scope == nil {
		return fmt.Errorf("parameter _scope is required, but nil was provided")
	}

	return nil
}

