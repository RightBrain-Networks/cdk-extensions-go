package glue


type CrawlerConfiguration struct {
	PartitionUpdateBehavior PartitionUpdateBehavior `field:"optional" json:"partitionUpdateBehavior" yaml:"partitionUpdateBehavior"`
	TableGroupingPolicy TableGroupingPolicy `field:"optional" json:"tableGroupingPolicy" yaml:"tableGroupingPolicy"`
	TableLevel *float64 `field:"optional" json:"tableLevel" yaml:"tableLevel"`
	TableUpdateBehavior TableUpdateBehavior `field:"optional" json:"tableUpdateBehavior" yaml:"tableUpdateBehavior"`
	Version ConfigurationVersion `field:"optional" json:"version" yaml:"version"`
}

