//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package aps

// Building without runtime type checking enabled, so all the below just return nil

func (t *jsiiProxy_TimeIntervalEntry) validateAddDaysOfTheMonthParameters(start *float64, end *float64) error {
	return nil
}

func (t *jsiiProxy_TimeIntervalEntry) validateAddMonthParameters(start *float64) error {
	return nil
}

func (t *jsiiProxy_TimeIntervalEntry) validateAddTimesParameters(start *string, end *string) error {
	return nil
}

func (t *jsiiProxy_TimeIntervalEntry) validateAddWeekdayParameters(start Weekday) error {
	return nil
}

func (t *jsiiProxy_TimeIntervalEntry) validateAddYearsParameters(start *float64) error {
	return nil
}

func (t *jsiiProxy_TimeIntervalEntry) validateBindParameters(_scope constructs.IConstruct) error {
	return nil
}

func validateNewTimeIntervalEntryParameters(options *TimeIntervalEntryProps) error {
	return nil
}

