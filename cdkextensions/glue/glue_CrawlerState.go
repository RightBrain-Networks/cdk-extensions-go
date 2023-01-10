package glue


type CrawlerState string

const (
	CrawlerState_CANCELLED CrawlerState = "CANCELLED"
	CrawlerState_FAILED CrawlerState = "FAILED"
	CrawlerState_SUCCEEDED CrawlerState = "SUCCEEDED"
)

