package alerting


type JiraTicketOverrideOptions struct {
	Assignee *string `field:"optional" json:"assignee" yaml:"assignee"`
	IssuePriority *string `field:"optional" json:"issuePriority" yaml:"issuePriority"`
	IssueType *string `field:"optional" json:"issueType" yaml:"issueType"`
	Project *string `field:"optional" json:"project" yaml:"project"`
}

