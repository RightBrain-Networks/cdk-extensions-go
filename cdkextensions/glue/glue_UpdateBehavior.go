package glue


type UpdateBehavior string

const (
	UpdateBehavior_UPDATE_IN_DATABASE UpdateBehavior = "UPDATE_IN_DATABASE"
	UpdateBehavior_LOG UpdateBehavior = "LOG"
)

