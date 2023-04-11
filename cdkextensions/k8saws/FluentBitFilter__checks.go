//go:build !no_runtime_type_checking

package k8saws

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
)

func validateFluentBitFilter_AppendFieldsParameters(match FluentBitMatch, records *[]*AppendedRecord) error {
	if match == nil {
		return fmt.Errorf("parameter match is required, but nil was provided")
	}

	for idx_a94e7b, v := range *records {
		if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter records[%#v]", idx_a94e7b) }); err != nil {
			return err
		}
	}

	return nil
}

func validateFluentBitFilter_BlacklistFieldsParameters(match FluentBitMatch) error {
	if match == nil {
		return fmt.Errorf("parameter match is required, but nil was provided")
	}

	return nil
}

func validateFluentBitFilter_GrepParameters(match FluentBitMatch, pattern *FluentBitGrepRegex) error {
	if match == nil {
		return fmt.Errorf("parameter match is required, but nil was provided")
	}

	if pattern == nil {
		return fmt.Errorf("parameter pattern is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(pattern, func() string { return "parameter pattern" }); err != nil {
		return err
	}

	return nil
}

func validateFluentBitFilter_KubernetesParameters(match FluentBitMatch) error {
	if match == nil {
		return fmt.Errorf("parameter match is required, but nil was provided")
	}

	return nil
}

func validateFluentBitFilter_LiftParameters(match FluentBitMatch, nestedUnder *string) error {
	if match == nil {
		return fmt.Errorf("parameter match is required, but nil was provided")
	}

	if nestedUnder == nil {
		return fmt.Errorf("parameter nestedUnder is required, but nil was provided")
	}

	return nil
}

func validateFluentBitFilter_ModifyParameters(match FluentBitMatch) error {
	if match == nil {
		return fmt.Errorf("parameter match is required, but nil was provided")
	}

	return nil
}

func validateFluentBitFilter_NestParameters(match FluentBitMatch, nestUnder *string) error {
	if match == nil {
		return fmt.Errorf("parameter match is required, but nil was provided")
	}

	if nestUnder == nil {
		return fmt.Errorf("parameter nestUnder is required, but nil was provided")
	}

	return nil
}

func validateFluentBitFilter_ParserParameters(match FluentBitMatch, key *string) error {
	if match == nil {
		return fmt.Errorf("parameter match is required, but nil was provided")
	}

	if key == nil {
		return fmt.Errorf("parameter key is required, but nil was provided")
	}

	return nil
}

func validateFluentBitFilter_RewriteTagParameters(match FluentBitMatch, rules *[]*RewriteTagRule) error {
	if match == nil {
		return fmt.Errorf("parameter match is required, but nil was provided")
	}

	for idx_6c621d, v := range *rules {
		if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter rules[%#v]", idx_6c621d) }); err != nil {
			return err
		}
	}

	return nil
}

func validateFluentBitFilter_ThrottleParameters(match FluentBitMatch, interval awscdk.Duration, rate *float64, window *float64) error {
	if match == nil {
		return fmt.Errorf("parameter match is required, but nil was provided")
	}

	if interval == nil {
		return fmt.Errorf("parameter interval is required, but nil was provided")
	}

	if rate == nil {
		return fmt.Errorf("parameter rate is required, but nil was provided")
	}

	if window == nil {
		return fmt.Errorf("parameter window is required, but nil was provided")
	}

	return nil
}

func validateFluentBitFilter_WhitelistFieldsParameters(match FluentBitMatch) error {
	if match == nil {
		return fmt.Errorf("parameter match is required, but nil was provided")
	}

	return nil
}

