package alerting


type SecurityHubFindingEventOptions struct {
	Overrides *[]IssueHandlerOverride `field:"optional" json:"overrides" yaml:"overrides"`
	Severity ISecurityHubSeverityConfiguration `field:"optional" json:"severity" yaml:"severity"`
}

