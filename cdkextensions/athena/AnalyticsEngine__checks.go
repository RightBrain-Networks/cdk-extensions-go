//go:build !no_runtime_type_checking

package athena

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func validateAnalyticsEngine_ApacheSparkParameters(options ApacheSparkEngineOptions) error {
	if options == nil {
		return fmt.Errorf("parameter options is required, but nil was provided")
	}

	return nil
}

func validateAnalyticsEngine_AthenaSqlParameters(options *AthenaSqlEngineOptions) error {
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

