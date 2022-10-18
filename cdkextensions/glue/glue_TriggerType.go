package glue


type TriggerType string

const (
	TriggerType_CONDITIONAL TriggerType = "CONDITIONAL"
	TriggerType_EVENT TriggerType = "EVENT"
	TriggerType_ON_DEMAND TriggerType = "ON_DEMAND"
	TriggerType_SCHEDULED TriggerType = "SCHEDULED"
)

