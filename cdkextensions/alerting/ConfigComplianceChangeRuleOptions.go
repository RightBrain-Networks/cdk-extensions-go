package alerting


type ConfigComplianceChangeRuleOptions struct {
	IncludeSecurityHub *bool `field:"optional" json:"includeSecurityHub" yaml:"includeSecurityHub"`
	Overrides *[]IssueHandlerOverride `field:"optional" json:"overrides" yaml:"overrides"`
	Severity *[]InspectorSeverity `field:"optional" json:"severity" yaml:"severity"`
}

