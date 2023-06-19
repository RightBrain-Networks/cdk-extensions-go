package alerting


type OpenSearchEventRuleOptions struct {
	Overrides *[]IssueHandlerOverride `field:"optional" json:"overrides" yaml:"overrides"`
	Severity *[]OpenSearchEventSeverity `field:"optional" json:"severity" yaml:"severity"`
	Types *[]OpenSearchEventType `field:"optional" json:"types" yaml:"types"`
}

