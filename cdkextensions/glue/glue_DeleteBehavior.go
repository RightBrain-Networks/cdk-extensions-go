package glue


type DeleteBehavior string

const (
	DeleteBehavior_DELETE_FROM_DATABASE DeleteBehavior = "DELETE_FROM_DATABASE"
	DeleteBehavior_DEPRECATE_IN_DATABASE DeleteBehavior = "DEPRECATE_IN_DATABASE"
	DeleteBehavior_LOG DeleteBehavior = "LOG"
)

