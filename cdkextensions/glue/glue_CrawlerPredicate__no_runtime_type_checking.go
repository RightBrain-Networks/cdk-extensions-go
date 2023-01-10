//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package glue

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CrawlerPredicate) validateBindParameters(scope constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CrawlerPredicate) validateBindOptionsParameters(_scope constructs.IConstruct) error {
	return nil
}

func validateNewCrawlerPredicateParameters(crawler ICrawler, options *CrawlerPredicateOptions) error {
	return nil
}

