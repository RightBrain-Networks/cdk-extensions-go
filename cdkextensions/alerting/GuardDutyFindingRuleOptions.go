package alerting


type GuardDutyFindingRuleOptions struct {
	Overrides *[]IssueHandlerOverride `field:"optional" json:"overrides" yaml:"overrides"`
	Severity *[]GuardDutySeverity `field:"optional" json:"severity" yaml:"severity"`
}

