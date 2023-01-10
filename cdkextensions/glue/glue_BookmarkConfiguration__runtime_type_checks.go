//go:build !no_runtime_type_checking
// +build !no_runtime_type_checking

package glue

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func validateBookmarkConfiguration_OfParameters(value *string, range_ *BookmarkRange) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(range_, func() string { return "parameter range_" }); err != nil {
		return err
	}

	return nil
}

func validateBookmarkConfiguration_PauseParameters(range_ *BookmarkRange) error {
	if err := _jsii_.ValidateStruct(range_, func() string { return "parameter range_" }); err != nil {
		return err
	}

	return nil
}

