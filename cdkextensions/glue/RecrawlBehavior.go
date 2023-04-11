package glue


type RecrawlBehavior string

const (
	RecrawlBehavior_EVENT_MODE RecrawlBehavior = "EVENT_MODE"
	RecrawlBehavior_EVERYTHING RecrawlBehavior = "EVERYTHING"
	RecrawlBehavior_NEW_FOLDERS_ONLY RecrawlBehavior = "NEW_FOLDERS_ONLY"
)

