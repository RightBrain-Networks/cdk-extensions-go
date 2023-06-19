package alerting


type InspectorFindingEventOptions struct {
	Overrides *[]IssueHandlerOverride `field:"optional" json:"overrides" yaml:"overrides"`
	Severity *[]InspectorSeverity `field:"optional" json:"severity" yaml:"severity"`
}

