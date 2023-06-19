package alerting


type EcrScanFindingEventOptions struct {
	Overrides *[]IssueHandlerOverride `field:"optional" json:"overrides" yaml:"overrides"`
	Severity IEcrImageScanSeverityConfiguration `field:"optional" json:"severity" yaml:"severity"`
}

