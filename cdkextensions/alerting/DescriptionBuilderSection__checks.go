//go:build !no_runtime_type_checking

package alerting

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsstepfunctions"
	"github.com/aws/constructs-go/constructs/v10"
)

func (d *jsiiProxy_DescriptionBuilderSection) validateAddIteratorParameters(id *string, props *DescriptionBuilderIteratorProps) error {
	if id == nil {
		return fmt.Errorf("parameter id is required, but nil was provided")
	}

	if props == nil {
		return fmt.Errorf("parameter props is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

func (d *jsiiProxy_DescriptionBuilderSection) validateAddReferenceParameters(id *string, props *AddReferenceProps) error {
	if id == nil {
		return fmt.Errorf("parameter id is required, but nil was provided")
	}

	if props == nil {
		return fmt.Errorf("parameter props is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

func (d *jsiiProxy_DescriptionBuilderSection) validateAddReferenceCheckParameters(ref *string) error {
	if ref == nil {
		return fmt.Errorf("parameter ref is required, but nil was provided")
	}

	return nil
}

func (d *jsiiProxy_DescriptionBuilderSection) validateBuildIdParameters(prefix *string) error {
	if prefix == nil {
		return fmt.Errorf("parameter prefix is required, but nil was provided")
	}

	return nil
}

func (d *jsiiProxy_DescriptionBuilderSection) validateRegisterBuilderParameters(builder IDescriptionBuilderComponent) error {
	if builder == nil {
		return fmt.Errorf("parameter builder is required, but nil was provided")
	}

	return nil
}

func (d *jsiiProxy_DescriptionBuilderSection) validateRegisterChainableParameters(chainable awsstepfunctions.IChainable) error {
	if chainable == nil {
		return fmt.Errorf("parameter chainable is required, but nil was provided")
	}

	return nil
}

func (d *jsiiProxy_DescriptionBuilderSection) validateSetDelimiterParameters(id *string, props *SetDelimiterProps) error {
	if id == nil {
		return fmt.Errorf("parameter id is required, but nil was provided")
	}

	if props == nil {
		return fmt.Errorf("parameter props is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

func (d *jsiiProxy_DescriptionBuilderSection) validateWriteParameters(id *string, props *WriteProps) error {
	if id == nil {
		return fmt.Errorf("parameter id is required, but nil was provided")
	}

	if props == nil {
		return fmt.Errorf("parameter props is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

func validateDescriptionBuilderSection_IsConstructParameters(x interface{}) error {
	if x == nil {
		return fmt.Errorf("parameter x is required, but nil was provided")
	}

	return nil
}

func validateNewDescriptionBuilderSectionParameters(scope constructs.IConstruct, id *string, props *DescriptionBuilderSectionProps) error {
	if scope == nil {
		return fmt.Errorf("parameter scope is required, but nil was provided")
	}

	if id == nil {
		return fmt.Errorf("parameter id is required, but nil was provided")
	}

	if props == nil {
		return fmt.Errorf("parameter props is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

