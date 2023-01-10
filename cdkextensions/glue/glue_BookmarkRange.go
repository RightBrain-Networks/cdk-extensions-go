package glue


type BookmarkRange struct {
	From *string `field:"required" json:"from" yaml:"from"`
	To *string `field:"required" json:"to" yaml:"to"`
}

