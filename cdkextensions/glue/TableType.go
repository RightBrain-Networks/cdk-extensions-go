package glue


type TableType string

const (
	TableType_EXTERNAL_TABLE TableType = "EXTERNAL_TABLE"
	TableType_VIRTUAL_VIEW TableType = "VIRTUAL_VIEW"
)

