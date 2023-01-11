package glue


// Configuration options that specify the state a crawler must meet in order to satisfy the conditions of the predicate.
type WorkflowCrawlerPredicateOptions struct {
	// The AWS account ID this resource belongs to.
	Account *string `field:"optional" json:"account" yaml:"account"`
	// ARN to deduce region and account from.
	//
	// The ARN is parsed and the account and region are taken from the ARN.
	// This should be used for imported resources.
	//
	// Cannot be supplied together with either `account` or `region`.
	EnvironmentFromArn *string `field:"optional" json:"environmentFromArn" yaml:"environmentFromArn"`
	// The value passed in by users to the physical name prop of the resource.
	//
	// - `undefined` implies that a physical name will be allocated by
	//    CloudFormation during deployment.
	// - a concrete value implies a specific physical name
	// - `PhysicalName.GENERATE_IF_NEEDED` is a marker that indicates that a physical will only be generated
	//    by the CDK if it is needed for cross-environment references. Otherwise, it will be allocated by CloudFormation.
	PhysicalName *string `field:"optional" json:"physicalName" yaml:"physicalName"`
	// The AWS region this resource belongs to.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// The logical operator which should be applied in determining whether a crawler meets the requested conditions.
	//
	// At the moment, the only supported operator is `EQUALS`.
	// See: [Trigger Predicate.Conditions.LogicalOperator](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-trigger-condition.html#cfn-glue-trigger-condition-logicaloperator)
	//
	LogicalOperator PredicateLogicalOperator `field:"required" json:"logicalOperator" yaml:"logicalOperator"`
	// The state that the crawler must be in in order to meet the criteria to trigger the next stage of the workflow.
	// See: [Trigger Predicate.Conditions.CrawlState](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-trigger-condition.html#cfn-glue-trigger-condition-crawlstate)
	//
	State CrawlerState `field:"optional" json:"state" yaml:"state"`
}

