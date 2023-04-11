package aps


// Configuration for the alert manager time interval.
type TimeIntervalEntryProps struct {
	// A list of ranges specifying the days of the month that the time interval should apply for.
	DaysOfTheMonth *[]*DayOfMonthRange `field:"optional" json:"daysOfTheMonth" yaml:"daysOfTheMonth"`
	// A list of ranges specifying the months that the time interval should apply for.
	Months *[]*MonthRange `field:"optional" json:"months" yaml:"months"`
	// A list of ranges specifying the time periods that the time interval should apply for.
	Times *[]*TimeRange `field:"optional" json:"times" yaml:"times"`
	// A string that matches a location in the IANA time zone database.
	//
	// For example, 'Australia/Sydney'. The location provides the time zone for
	// the time interval.
	//
	// You may also use `Local` as a location to use the local time of the
	// machine where Alertmanager is running, or `UTC` for UTC time. If no
	// timezone is provided, the time interval is taken to be in UTC time.
	TimeZone *string `field:"optional" json:"timeZone" yaml:"timeZone"`
	// A list of ranges specifying the weekdays that the time interval should apply for.
	Weekdays *[]*WeekdayRange `field:"optional" json:"weekdays" yaml:"weekdays"`
	// A list of ranges specifying the years that the time interval should apply for.
	Years *[]*YearRange `field:"optional" json:"years" yaml:"years"`
}

