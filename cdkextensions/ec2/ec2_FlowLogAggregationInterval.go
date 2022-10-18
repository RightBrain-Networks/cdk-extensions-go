package ec2


type FlowLogAggregationInterval string

const (
	// Flow logs will be written at least every 60 seconds.
	FlowLogAggregationInterval_ONE_MINUTE FlowLogAggregationInterval = "ONE_MINUTE"
	// Flow logs will be written at least every ten minutes.
	FlowLogAggregationInterval_TEN_MINUTES FlowLogAggregationInterval = "TEN_MINUTES"
)

