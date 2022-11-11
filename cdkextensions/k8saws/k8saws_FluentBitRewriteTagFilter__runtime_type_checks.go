//go:build !no_runtime_type_checking
// +build !no_runtime_type_checking

package k8saws

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/constructs-go/constructs/v10"
)

func (f *jsiiProxy_FluentBitRewriteTagFilter) validateAddRuleParameters(rule *RewriteTagRule) error {
	if rule == nil {
		return fmt.Errorf("parameter rule is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(rule, func() string { return "parameter rule" }); err != nil {
		return err
	}

	return nil
}

func (f *jsiiProxy_FluentBitRewriteTagFilter) validateBindParameters(_scope constructs.IConstruct) error {
	if _scope == nil {
		return fmt.Errorf("parameter _scope is required, but nil was provided")
	}

	return nil
}

func (f *jsiiProxy_FluentBitRewriteTagFilter) validateRenderConfigFileParameters(config *map[string]interface{}) error {
	if config == nil {
		return fmt.Errorf("parameter config is required, but nil was provided")
	}

	return nil
}

func validateNewFluentBitRewriteTagFilterParameters(options *FluentBitRewriteTagFilterOptions) error {
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

