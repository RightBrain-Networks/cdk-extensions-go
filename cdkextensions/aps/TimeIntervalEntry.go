package aps

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/vibe-io/cdk-extensions-go/cdkextensions/jsii"

	"github.com/aws/constructs-go/constructs/v10"
)

// An object specifying a collection of ranges that together make up an interval of time.
//
// Referenced by alert manager to define periods for which certain settings
// should apply.
// See: [Time Interval Official Documentation](https://prometheus.io/docs/alerting/latest/configuration/#time_interval-0)
//
type TimeIntervalEntry interface {
	// Collection of day of the month ranges for which this time interval will apply.
	DaysOfTheMonth() *[]*DayOfMonthRange
	// Collection of month ranges for which this time interval will apply.
	Months() *[]*MonthRange
	// Collection of time ranges for which this time interval will apply.
	Times() *[]*TimeRange
	// A string that matches a location in the IANA time zone database.
	//
	// For example, 'Australia/Sydney'. The location provides the time zone for
	// the time interval.
	//
	// You may also use `Local` as a location to use the local time of the
	// machine where Alertmanager is running, or `UTC` for UTC time. If no
	// timezone is provided, the time interval is taken to be in UTC time.
	TimeZone() *string
	// Collection of weekday ranges for which this time interval will apply.
	Weekdays() *[]*WeekdayRange
	// Collection of year ranges for which this time interval will apply.
	Years() *[]*YearRange
	// Adds a range specifying the numerical days in the month.
	//
	// Days begin at 1. Negative values are also accepted which begin at the end
	// of the month, e.g. -1 during January would represent January 31. For
	// example: `start: 1` and `end: 5` or `start:-3` and `end: -1` would both be
	// valid ranges.
	//
	// Extending past the start or end of the month will cause it to be clamped.
	// E.g. specifying `start: 1` and `end: 31` during February will clamp the
	// actual end date to 28 or 29 depending on leap years.
	//
	// Inclusive on both ends.
	//
	// Returns: The time interval the range was added to.
	AddDaysOfTheMonth(start *float64, end *float64) TimeIntervalEntry
	// A range of calendar months identified by number, where January = 1.
	//
	// Ranges are also accepted by specifying `end` inclusive on both ends.
	//
	// Returns: The time interval the range was added to.
	AddMonth(start *float64, end *float64) TimeIntervalEntry
	// Ranges inclusive of the starting time and exclusive of the end time to make it easy to represent times that start/end on hour boundaries.
	//
	// For example, `start: '17:00'` and `end: '24:00'` will begin at 17:00 and
	// finish immediately before 24:00.
	//
	// Returns: The time interval the range was added to.
	AddTimes(start *string, end *string) TimeIntervalEntry
	// Adds a day of the week, where the week begins on Sunday and ends on Saturday.
	//
	// For convenience, ranges are also accepted by specifying `end` and are
	// inclusive on both ends.
	//
	// Returns: The time interval the range was added to.
	AddWeekday(start Weekday, end Weekday) TimeIntervalEntry
	// Adds a numerical range of years.
	//
	// Ranges to cover multiple years are accepted. For example, `start: 2020` and
	// `end: 2022`.
	//
	// Inclusive on both ends.
	//
	// Returns: The time interval the range was added to.
	AddYears(start *float64, end *float64) TimeIntervalEntry
	// Associates the time interval with a construct that is handling the configuration of alert manager.
	//
	// Returns: An alert manager `time_interval` configuration object.
	Bind(_scope constructs.IConstruct) *map[string]interface{}
}

// The jsii proxy struct for TimeIntervalEntry
type jsiiProxy_TimeIntervalEntry struct {
	_ byte // padding
}

func (j *jsiiProxy_TimeIntervalEntry) DaysOfTheMonth() *[]*DayOfMonthRange {
	var returns *[]*DayOfMonthRange
	_jsii_.Get(
		j,
		"daysOfTheMonth",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TimeIntervalEntry) Months() *[]*MonthRange {
	var returns *[]*MonthRange
	_jsii_.Get(
		j,
		"months",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TimeIntervalEntry) Times() *[]*TimeRange {
	var returns *[]*TimeRange
	_jsii_.Get(
		j,
		"times",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TimeIntervalEntry) TimeZone() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timeZone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TimeIntervalEntry) Weekdays() *[]*WeekdayRange {
	var returns *[]*WeekdayRange
	_jsii_.Get(
		j,
		"weekdays",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TimeIntervalEntry) Years() *[]*YearRange {
	var returns *[]*YearRange
	_jsii_.Get(
		j,
		"years",
		&returns,
	)
	return returns
}


// Creates a new instance of the TimeIntervalEntry class.
func NewTimeIntervalEntry(options *TimeIntervalEntryProps) TimeIntervalEntry {
	_init_.Initialize()

	if err := validateNewTimeIntervalEntryParameters(options); err != nil {
		panic(err)
	}
	j := jsiiProxy_TimeIntervalEntry{}

	_jsii_.Create(
		"cdk-extensions.aps.TimeIntervalEntry",
		[]interface{}{options},
		&j,
	)

	return &j
}

// Creates a new instance of the TimeIntervalEntry class.
func NewTimeIntervalEntry_Override(t TimeIntervalEntry, options *TimeIntervalEntryProps) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk-extensions.aps.TimeIntervalEntry",
		[]interface{}{options},
		t,
	)
}

func (t *jsiiProxy_TimeIntervalEntry) AddDaysOfTheMonth(start *float64, end *float64) TimeIntervalEntry {
	if err := t.validateAddDaysOfTheMonthParameters(start, end); err != nil {
		panic(err)
	}
	var returns TimeIntervalEntry

	_jsii_.Invoke(
		t,
		"addDaysOfTheMonth",
		[]interface{}{start, end},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TimeIntervalEntry) AddMonth(start *float64, end *float64) TimeIntervalEntry {
	if err := t.validateAddMonthParameters(start); err != nil {
		panic(err)
	}
	var returns TimeIntervalEntry

	_jsii_.Invoke(
		t,
		"addMonth",
		[]interface{}{start, end},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TimeIntervalEntry) AddTimes(start *string, end *string) TimeIntervalEntry {
	if err := t.validateAddTimesParameters(start, end); err != nil {
		panic(err)
	}
	var returns TimeIntervalEntry

	_jsii_.Invoke(
		t,
		"addTimes",
		[]interface{}{start, end},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TimeIntervalEntry) AddWeekday(start Weekday, end Weekday) TimeIntervalEntry {
	if err := t.validateAddWeekdayParameters(start); err != nil {
		panic(err)
	}
	var returns TimeIntervalEntry

	_jsii_.Invoke(
		t,
		"addWeekday",
		[]interface{}{start, end},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TimeIntervalEntry) AddYears(start *float64, end *float64) TimeIntervalEntry {
	if err := t.validateAddYearsParameters(start); err != nil {
		panic(err)
	}
	var returns TimeIntervalEntry

	_jsii_.Invoke(
		t,
		"addYears",
		[]interface{}{start, end},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TimeIntervalEntry) Bind(_scope constructs.IConstruct) *map[string]interface{} {
	if err := t.validateBindParameters(_scope); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		t,
		"bind",
		[]interface{}{_scope},
		&returns,
	)

	return returns
}

