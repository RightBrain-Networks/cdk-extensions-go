package alerting


type OpenSearchEventTypeProps struct {
	DetailType *string `field:"required" json:"detailType" yaml:"detailType"`
	EventName *string `field:"required" json:"eventName" yaml:"eventName"`
}

