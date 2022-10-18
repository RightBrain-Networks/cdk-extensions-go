//go:build !no_runtime_type_checking
// +build !no_runtime_type_checking

package glue

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func (j *jsiiProxy_JdbcTarget) validateAddExclusionParameters(exclusion *string) error {
	if exclusion == nil {
		return fmt.Errorf("parameter exclusion is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_JdbcTarget) validateAddPathParameters(path *string) error {
	if path == nil {
		return fmt.Errorf("parameter path is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_JdbcTarget) validateBindParameters(_crawler Crawler) error {
	if _crawler == nil {
		return fmt.Errorf("parameter _crawler is required, but nil was provided")
	}

	return nil
}

func validateNewJdbcTargetParameters(connection Connection, options *JdbcTargetOptions) error {
	if connection == nil {
		return fmt.Errorf("parameter connection is required, but nil was provided")
	}

	if options == nil {
		return fmt.Errorf("parameter options is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

