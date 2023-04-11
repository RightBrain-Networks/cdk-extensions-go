//go:build no_runtime_type_checking

package glue

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_Crawler) validateAddClassifierParameters(classifier *string) error {
	return nil
}

func (c *jsiiProxy_Crawler) validateAddTargetParameters(target ICrawlerTarget) error {
	return nil
}

func (c *jsiiProxy_Crawler) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (c *jsiiProxy_Crawler) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (c *jsiiProxy_Crawler) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func validateCrawler_FromCrawlerArnParameters(scope constructs.IConstruct, id *string, crawlerArn *string) error {
	return nil
}

func validateCrawler_FromCrawlerNameParameters(scope constructs.IConstruct, id *string, crawlerName *string) error {
	return nil
}

func validateCrawler_IsConstructParameters(x interface{}) error {
	return nil
}

func validateCrawler_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCrawler_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewCrawlerParameters(scope constructs.Construct, id *string, props *CrawlerProps) error {
	return nil
}

