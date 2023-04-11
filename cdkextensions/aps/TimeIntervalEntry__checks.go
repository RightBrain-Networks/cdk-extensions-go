//go:build !no_runtime_type_checking

package aps

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/constructs-go/constructs/v10"
)

func (t *jsiiProxy_TimeIntervalEntry) validateAddDaysOfTheMonthParameters(start *float64, end *float64) error {
	if start == nil {
		return fmt.Errorf("parameter start is required, but nil was provided")
	}

	if end == nil {
		return fmt.Errorf("parameter end is required, but nil was provided")
	}

	return nil
}

func (t *jsiiProxy_TimeIntervalEntry) validateAddMonthParameters(start *float64) error {
	if start == nil {
		return fmt.Errorf("parameter start is required, but nil was provided")
	}

	return nil
}

func (t *jsiiProxy_TimeIntervalEntry) validateAddTimesParameters(start *string, end *string) error {
	if start == nil {
		return fmt.Errorf("parameter start is required, but nil was provided")
	}

	if end == nil {
		return fmt.Errorf("parameter end is required, but nil was provided")
	}

	return nil
}

func (t *jsiiProxy_TimeIntervalEntry) validateAddWeekdayParameters(start Weekday) error {
	if start == "" {
		return fmt.Errorf("parameter start is required, but nil was provided")
	}

	return nil
}

func (t *jsiiProxy_TimeIntervalEntry) validateAddYearsParameters(start *float64) error {
	if start == nil {
		return fmt.Errorf("parameter start is required, but nil was provided")
	}

	return nil
}

func (t *jsiiProxy_TimeIntervalEntry) validateBindParameters(_scope constructs.IConstruct) error {
	if _scope == nil {
		return fmt.Errorf("parameter _scope is required, but nil was provided")
	}

	return nil
}

func validateNewTimeIntervalEntryParameters(options *TimeIntervalEntryProps) error {
	if options == nil {
		return fmt.Errorf("parameter options is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

