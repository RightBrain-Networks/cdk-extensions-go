package resourcegroups


type CloudFormationStackProps struct {
	ResourceTypes *[]*string `field:"optional" json:"resourceTypes" yaml:"resourceTypes"`
}

