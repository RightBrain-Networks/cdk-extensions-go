package alerting


type JiraTicketPriorityMap struct {
	Critical *string `field:"optional" json:"critical" yaml:"critical"`
	Default *string `field:"optional" json:"default" yaml:"default"`
	High *string `field:"optional" json:"high" yaml:"high"`
	Info *string `field:"optional" json:"info" yaml:"info"`
	Low *string `field:"optional" json:"low" yaml:"low"`
	Medium *string `field:"optional" json:"medium" yaml:"medium"`
}

