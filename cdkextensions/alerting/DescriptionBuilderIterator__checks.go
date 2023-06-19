//go:build !no_runtime_type_checking

package alerting

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsstepfunctions"
	"github.com/aws/constructs-go/constructs/v10"
)

func (d *jsiiProxy_DescriptionBuilderIterator) validateAddIteratorParameters(id *string, props *DescriptionBuilderIteratorProps) error {
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

func (d *jsiiProxy_DescriptionBuilderIterator) validateAddReferenceParameters(id *string, props *AddReferenceProps) error {
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

func (d *jsiiProxy_DescriptionBuilderIterator) validateBuildIdParameters(prefix *string) error {
	if prefix == nil {
		return fmt.Errorf("parameter prefix is required, but nil was provided")
	}

	return nil
}

func (d *jsiiProxy_DescriptionBuilderIterator) validateRegisterBuilderParameters(builder IDescriptionBuilderComponent) error {
	if builder == nil {
		return fmt.Errorf("parameter builder is required, but nil was provided")
	}

	return nil
}

func (d *jsiiProxy_DescriptionBuilderIterator) validateRegisterChainableParameters(chainable awsstepfunctions.IChainable) error {
	if chainable == nil {
		return fmt.Errorf("parameter chainable is required, but nil was provided")
	}

	return nil
}

func (d *jsiiProxy_DescriptionBuilderIterator) validateSetDelimiterParameters(id *string, props *SetDelimiterProps) error {
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

func (d *jsiiProxy_DescriptionBuilderIterator) validateWriteParameters(id *string, props *WriteProps) error {
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

func validateDescriptionBuilderIterator_IsConstructParameters(x interface{}) error {
	if x == nil {
		return fmt.Errorf("parameter x is required, but nil was provided")
	}

	return nil
}

func validateNewDescriptionBuilderIteratorParameters(scope constructs.IConstruct, id *string, props *DescriptionBuilderIteratorProps) error {
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

